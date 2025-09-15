package cmd

import (
	"fmt"
	"strconv"

	"github.com/Sumedhvats/task/db"
	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Complete a particular command",
	Run: func(cmd *cobra.Command, args []string) {
		ids := []int{}
		for _, arg := range args {
			curr, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse the argument: ", arg)
			} else {

				ids = append(ids, curr)
			}
		}
		tasks,err:=db.AllTask()
		if err!=nil {
			fmt.Println("Something went wrong!\n",err.Error())
			return
		}
		for _,id:=range ids {
			if id<=0||id>len(tasks) {
				fmt.Println("Invalid Task number: ",id)
				continue
			}
			task:=tasks[id-1]
			err:=db.DeleteTask(task.Key)
			if err!=nil {
				fmt.Printf("Failed to Mark %d as complete. Error: %s\n",id,err)
			}else{
				
				fmt.Printf("Completed task %d \n",id)
			}
			
		}
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
