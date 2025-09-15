package cmd

import (
	"fmt"
	"os"

	"github.com/Sumedhvats/task/db"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the tasks pending",
Run: func(cmd *cobra.Command, args []string) {
	tasks,err:=db.AllTask()
	if err!=nil {
		fmt.Println("Something went wrong",err.Error())
		os.Exit(1)
	}
	if len(tasks)==0 {
		fmt.Println("You have no tasks")
		return
	}
	for i,v:= range tasks{
fmt.Printf("%d) %s\n",i+1,v.Value)
	}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
