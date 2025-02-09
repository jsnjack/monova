/*
Copyright © 2025 YAUHEN SHULITSKI

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var Version = "dev"

var FlagVersion bool
var FlagDebug bool

// CurrentDir is the current working directory
var CurrentDir string

// Logger is the logger instance, which is used to print messages
var Logger *log.Logger

// DebugLogger is the logger instance, which is used to print debug messages,
// if the debug mode is enabled. Otherwise, it is a discard logger. Debug messages
// are printed to stderr.
var DebugLogger *log.Logger

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "monova",
	Short: "calculate new version and print it",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		var err error

		// Check current directory
		CurrentDir, err = os.Getwd()
		if err != nil {
			return err
		}

		// Initialize logger
		Logger = log.New(os.Stdout, "", 0)

		if FlagDebug {
			DebugLogger = log.New(os.Stderr, "", log.Lmicroseconds|log.Lshortfile)
		} else {
			DebugLogger = log.New(io.Discard, "", 0)
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.SilenceUsage = true

		if FlagVersion {
			fmt.Println(Version)
			return nil
		}

		history, err := NewHistory(CurrentDir)
		if err != nil {
			return err
		}

		version, err := history.GetVersion(false)
		if err != nil {
			return err
		}
		fmt.Println(version.String())

		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&FlagVersion, "version", "v", false, "print version and exit")
	rootCmd.PersistentFlags().BoolVarP(&FlagDebug, "debug", "d", false, "enable debug mode")
}
