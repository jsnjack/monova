/*
Copyright © 2025 YAUHEN SHULITSKI
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// historyCmd represents the history command
var historyCmd = &cobra.Command{
	Use:   "history",
	Short: "print the history of versions",
	RunE: func(cmd *cobra.Command, args []string) error {
		history, err := NewHistory(CurrentDir)
		if err != nil {
			return err
		}

		history.Reset()

		history.GetVersion(true)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(historyCmd)
}
