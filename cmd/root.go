package cmd

import (

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "task",
	Short: "Task is a cli task manager",
}

func Execute()error {
	err := rootCmd.Execute()
return err	
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
