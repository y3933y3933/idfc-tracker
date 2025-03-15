/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"database/sql"
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/y3933y3933/idfc-tracker/internal/database"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		dbQueries := ctx.Value("dbQueries").(*database.Queries)

		var name string
		for {
			name, _ = pterm.DefaultInteractiveTextInput.Show("What is your name? (in English)")
			name = strings.TrimSpace(name)

			if name == "" {
				pterm.Warning.Println("Name cannot be empty. Please enter your name.")
				continue
			}

			_, err := dbQueries.GetUserByName(cmd.Context(), name)
			if err == nil {
				pterm.Error.Println("Name already exists. Please enter a different one.")
				continue
			}

			if errors.Is(err, sql.ErrNoRows) {
				err = dbQueries.CreateUser(cmd.Context(), name)
				if err != nil {
					pterm.Error.Printf("Create user fail: %v\n", err)
					os.Exit(1)
				}
				break
			}
			checkDbError(err)

		}

		goalOptions := []string{"10", "20", "30", "40", "50"}
		selectedGoal, _ := pterm.DefaultInteractiveSelect.WithOptions(goalOptions).WithDefaultOption("10").Show("Please select your goal")

		user, err := dbQueries.GetUserByName(cmd.Context(), name)
		checkDbError(err)

		goalInt64, err := strconv.ParseInt(selectedGoal, 10, 64)

		if err != nil {
			pterm.Error.Printf("string convert to int error: %v\n", err)
			os.Exit(1)
		}

		err = dbQueries.CreatePoint(cmd.Context(), database.CreatePointParams{
			UserID: user.ID,
			Goal:   goalInt64,
		})
		checkDbError(err)
		userIDStr := strconv.FormatInt(user.ID, 10)
		err = dbQueries.SetActiveUserID(cmd.Context(), userIDStr)
		checkDbError(err)

		pterm.Println("🎉 Great job,", pterm.LightYellow(name)+"! Your goal is set to", pterm.LightYellow(selectedGoal), "points. Let's get started! 🚀")

	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func checkDbError(err error) {
	if err != nil {
		pterm.Error.Printf("Database error: %v\n", err)
		os.Exit(1)
	}
}
