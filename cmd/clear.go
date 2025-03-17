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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
