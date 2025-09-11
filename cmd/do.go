package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Complete a particular command",
	Run: func(cmd *cobra.Command, args []string) {
		tasks := []int{}
		for _, arg := range args {
			curr, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse the argument: ", arg)
			} else {

				tasks = append(tasks, curr)
			}
		}
		fmt.Println(tasks)
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
