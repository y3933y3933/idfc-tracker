/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/y3933y3933/idfc-tracker/cmd"
	"github.com/y3933y3933/idfc-tracker/internal/database"
)

func main() {
	db, err := sql.Open("sqlite3", "./app.db")
	checkErr(err)

	dbQueries := database.New(db)
	history, err := dbQueries.ListHistory(context.Background())
	checkErr(err)
	fmt.Println(history)
	cmd.Execute()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
