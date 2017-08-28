package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
)

const cacheFilename = ".monova.cache"
const checkpointPrefix = "Version:"
const checkpointSuffix = " generated by monova"

var (
	errNoCommits = errors.New("No commits in repository")
)

// Repository is an object that represents git repository
type Repository struct {
	Path         string
	commitCursor int
	Cache        *Cache
	Config       *Config
}

// GetNextCommit returns a Commit object to which points commitCursor
func (r *Repository) GetNextCommit() (*Commit, error) {
	cmd := exec.Command("git", "log", "--format='%H%s'", "--max-count=1", fmt.Sprintf("--skip=%d", r.commitCursor))
	cmd.Dir = r.Path
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	if len(output) == 0 {
		return nil, errNoCommits
	}
	commit, err := CreateCommit(string(output))
	if err != nil {
		return nil, err
	}
	r.commitCursor = r.commitCursor + 1
	return commit, nil
}

// CreateCheckpoint creates empty commit which fixes the version
func (r *Repository) CreateCheckpoint(args []string) error {
	switch len(args) {
	case 0:
		subject, _ := r.generateCheckpointSubject("")
		err := r.createCheckpointCommit(subject)
		if err != nil {
			return err
		}
		fmt.Printf("Commit with subject %s created\n", subject)
	default:
		subject, err := r.generateCheckpointSubject(args[0])
		if err != nil {
			return err
		}
		err = r.createCheckpointCommit(subject)
		if err != nil {
			return err
		}
		fmt.Printf("Commit with subject %s created\n", subject)

	}
	return nil
}

func (r *Repository) createCheckpointCommit(subject string) error {
	cmd := exec.Command("git", "commit", "--allow-empty", "-m", subject)
	cmd.Dir = r.Path
	err := cmd.Run()
	return err
}

func (r *Repository) generateCheckpointSubject(checkpoint string) (string, error) {
	var template = checkpointPrefix + "%s" + checkpointSuffix
	switch checkpoint {
	case "":
		return fmt.Sprintf(template, r.GetVersion()), nil
	default:
		version, err := r.cleanCheckpoint(checkpoint)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf(template, version), nil
	}
}

func (r *Repository) cleanCheckpoint(checkpoint string) (string, error) {
	numbers, err := SplitVersion(checkpoint)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d.%d.%d", numbers[0], numbers[1], numbers[2]), nil
}

// UpdateVersion updates and returns package version
func (r *Repository) UpdateVersion() (string, error) {
	var commits []*Commit
commitLoop:
	for {
		commit, err := r.GetNextCommit()
		switch err {
		case errNoCommits:
			break commitLoop
		}
		if err != nil {
			return "", err
		}
		if commit.ID != r.Cache.CommitID {
			commits = append(commits, commit)
		} else {
			break
		}
	}
	for _, value := range ReverseCommits(commits) {
		err := r.Cache.Version.Update(value.Subject, r.Config)
		if err != nil {
			fmt.Println(err)
		}
		err = SaveHistory(value, r.GetVersion())
		if err != nil {
			fmt.Println(err)
		}
		if *debugFlag {
			fmt.Printf("%s --> %s\n", value.Subject, r.GetVersion())
		}
	}
	if len(commits) > 0 {
		r.Cache.CommitID = commits[len(commits)-1].ID
	}
	r.Cache.Save()
	return r.GetVersion(), nil
}

// GetVersion returns current version
func (r *Repository) GetVersion() string {
	version := fmt.Sprintf("%d.%d.%d", r.Cache.Version.Major, r.Cache.Version.Minor, r.Cache.Version.Patch)
	return version
}

// CreateRepository returns Repository object for given path
func CreateRepository(repoPath string) (*Repository, error) {
	stat, err := os.Stat(repoPath)
	if err != nil {
		return nil, err
	}
	if !stat.IsDir() {
		return nil, fmt.Errorf("%s is not a directory", repoPath)
	}
	if !isGitRepository(repoPath) {
		return nil, fmt.Errorf("%s is not a git repository", repoPath)
	}
	cachePath := path.Join(repoPath, cacheFilename)
	cache, err := CreateCache(cachePath)
	if err != nil {
		return nil, err
	}
	repo := Repository{
		Path:   repoPath,
		Cache:  cache,
		Config: CreateConfig(),
	}
	return &repo, nil
}

func isGitRepository(path string) bool {
	cmd := exec.Command("git", "status")
	cmd.Dir = path
	err := cmd.Run()
	if err != nil {
		return false
	}
	return true
}

// ReverseCommits reverses order of the commits
func ReverseCommits(commits []*Commit) []*Commit {
	for i, j := 0, len(commits)-1; i < j; i, j = i+1, j-1 {
		commits[i], commits[j] = commits[j], commits[i]
	}
	return commits
}

// SplitVersion returns array of version [major, minor, patch]
func SplitVersion(checkpoint string) ([]int, error) {
	var numbers []int
	splits := strings.Split(checkpoint, ".")
	if len(splits) != 3 {
		return numbers, fmt.Errorf("Not valid checkpoint: %s", checkpoint)
	}
	for _, value := range splits {
		number, err := strconv.Atoi(value)
		if err != nil {
			return numbers, fmt.Errorf("Not valid checkpoint: %s", checkpoint)
		}
		numbers = append(numbers, number)
	}
	return numbers, nil
}
