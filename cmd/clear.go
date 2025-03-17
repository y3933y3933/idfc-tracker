/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/y3933y3933/idfc-tracker/internal/database"
)

// clearCmd represents the clear command
var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear the current user's points and history",
	Long: `This command allows the active user to clear all their accumulated points and delete all entries in the history.
It will reset the 'points' table and remove all data from the 'history' table related to the active user.
Example usage:

  $ idfc-tracker clear
  All points and history for the active user will be cleared.
  
This is useful if you want to start fresh or reset your progress.`,

	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		dbQueries := ctx.Value("dbQueries").(*database.Queries)

		activeUser, err := database.GetActiveUser(ctx, dbQueries)

		if err != nil {
			pterm.Warning.Println("No user found. Run init to create a user and set your goal.")
			return
		}

		confirm, _ := pterm.DefaultInteractiveConfirm.Show("⚠️ Are you sure you want to clear all history and reset total points to 0? This action cannot be undone.")
		if !confirm {
			pterm.Warning.Println("Operation cancelled. No changes were made.")
			return
		}

		pterm.Info.Println("Clearing history and resetting total points...")

		err = dbQueries.ResetUserPoints(ctx, activeUser.ID)
		err = dbQueries.ClearUserHistory(ctx, activeUser.ID)
		if err != nil {
			log.Fatalf("Database error:%v\n", err)
		}

		pterm.Success.Println("✅ All history has been cleared, and total points have been reset to 0.")

	},
}

func init() {
	rootCmd.AddCommand(clearCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clearCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clearCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
