package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/Sumedhvats/task/db"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a command to your task list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
	_,err:=db.CreateTask(task)
if err!=nil {
		fmt.Println("Something went wrong",err.Error())
		os.Exit(1)
	}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
