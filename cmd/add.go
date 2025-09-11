package cmd

import (
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a command to your task list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		println(task)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
