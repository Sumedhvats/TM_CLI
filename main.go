/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"os"
	"path/filepath"

	"github.com/Sumedhvats/task/cmd"
	"github.com/Sumedhvats/task/db"
)

func main() {
	home,_:= os.UserHomeDir()
	dbPath:=filepath.Join(home,"tasks.db")
	err:=db.Init(dbPath)
	if err!=nil{
		panic(err)
	}
	cmd.Execute()
}
