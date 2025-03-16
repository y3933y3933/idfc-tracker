/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
	"github.com/y3933y3933/idfc-tracker/internal/database"
)

const dbName = "app.db"

var db *sql.DB

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "idfc-tracker",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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

		cmd.SetContext(context.WithValue(cmd.Context(), "dbQueries", dbQueries))
		return nil
	},

	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		if db != nil {
			db.Close()
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

func getActiveUser(ctx context.Context, dbQueries *database.Queries) (database.GetUserByIDRow, error) {
	activeUserIDStr, err := dbQueries.GetActiveUserID(ctx)
	if err != nil {
		return database.GetUserByIDRow{}, fmt.Errorf("failed to get active user ID: %w", err)

	}

	activeUserID, err := strconv.ParseInt(activeUserIDStr, 10, 64)
	if err != nil {
		return database.GetUserByIDRow{}, fmt.Errorf("failed to convert active user ID to integer: %w", err)
	}

	user, err := dbQueries.GetUserByID(ctx, activeUserID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return database.GetUserByIDRow{}, fmt.Errorf("no user found for active user ID: %w", err)
		}
		return database.GetUserByIDRow{}, fmt.Errorf("failed to retrieve user from DB: %w", err)
	}

	return user, nil
}
