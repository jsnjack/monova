/*
Copyright © 2025 YAUHEN SHULITSKI
*/
package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// checkpointCmd represents the checkpoint command
var checkpointCmd = &cobra.Command{
	Use:   "checkpoint version",
	Short: "manually set the version",
	RunE: func(cmd *cobra.Command, args []string) error {

		if len(args) != 1 {
			return fmt.Errorf("provide the version to set")
		}
		versionArg := args[0]

		DebugLogger.Printf("Setting version to %s\n", versionArg)

		// Validate the version
		version := &PackageVersion{}
		var err error
		version.Major, version.Minor, version.Patch, err = parseStringVersion(versionArg)
		if err != nil {
			return fmt.Errorf("invalid version: %w", err)
		}

		// Create a commit message
		gitCmd := exec.Command("git", "commit", "--allow-empty", "-m", version.ToCheckpointMessage())
		err = gitCmd.Run()
		if err != nil {
			return fmt.Errorf("failed to create git commit: %w", err)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(checkpointCmd)
}
