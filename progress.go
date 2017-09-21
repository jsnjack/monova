package main

import "fmt"

const marker = "#"
const filler = " "

// ProgressBar represents progressbar
type ProgressBar struct {
	count int
}

// ClearLine clears current line in terminal
func (p *ProgressBar) ClearLine() {
	fmt.Print("\033[2K\033[0G")
}

// Next advances progress bar
func (p *ProgressBar) Next() string {
	p.count = p.count + 1
	return fmt.Sprintf("Analyzing commits: @%d", p.count)
}
