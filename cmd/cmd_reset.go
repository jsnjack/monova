/*
Copyright © 2025 YAUHEN SHULITSKI
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// resetCmd represents the reset command
var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "remove history file",
	RunE: func(cmd *cobra.Command, args []string) error {
		history, err := NewHistory(CurrentDir)
		if err != nil {
			return err
		}

		return history.Reset()
	},
}

func init() {
	rootCmd.AddCommand(resetCmd)
}
