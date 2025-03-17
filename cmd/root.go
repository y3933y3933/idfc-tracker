/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
	"github.com/y3933y3933/idfc-tracker/internal/database"
)

const dbName = "app.db"

type contextKey string

const dbQueriesKey contextKey = "dbQueries"

var db *sql.DB

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "idfc-tracker",
	Short: "A CLI tool to track your progress towards a goal through points",
	Long: `IDFC Tracker is a command-line tool that helps you track your progress towards a personal goal by adding points over time. 
You can create a user, set a goal, and then incrementally add points to track your progress.
Each point addition is logged with a reason, and you can view your history, set a new goal, and clear all your data at any time.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		var err error
		db, err = sql.Open("sqlite3", dbName)
		if err != nil {
			return err
		}
		if err := db.Ping(); err != nil {
			return err
		}

		dbQueries := database.New(db)

		cmd.SetContext(context.WithValue(cmd.Context(), dbQueriesKey, dbQueries))
		return nil
	},

	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		if db != nil {
			err := db.Close()
			if err != nil {
				log.Fatalf("db close fail: %v", err)
			}
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.idfc-tracker.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
