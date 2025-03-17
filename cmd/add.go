/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"strconv"
	"strings"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/y3933y3933/idfc-tracker/internal/database"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
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
		points := getAddPoints()
		reason := getReason()

		userPoint, err := dbQueries.GetPointByUserID(cmd.Context(), activeUser.ID)
		if err != nil {
			log.Fatalf("GetPointByUserID fail: %v", userPoint)
		}

		err = dbQueries.InsertHistory(ctx, database.InsertHistoryParams{
			UserID: activeUser.ID,
			Point:  int64(points),
			Reason: reason,
		})
		if err != nil {
			log.Fatalf("Insert History fail: %v\n", err)
		}

		err = dbQueries.UpdateTotalByUserID(ctx, database.UpdateTotalByUserIDParams{
			Total:  userPoint.Total + int64(points),
			UserID: activeUser.ID,
		})

		if err != nil {
			log.Fatalf("update Total fail: %v", err)
		}

		pterm.Success.Printf("📢 Attention! %s just added %d points. Why? Because: %s\n", activeUser.Name, points, reason)

	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getAddPoints() uint {
	var points uint
	for {
		input, _ := pterm.DefaultInteractiveTextInput.Show("Enter the number of points to add (must be a positive integer)")
		// 轉換輸入為整數
		num, err := strconv.Atoi(input)

		// 確保是正整數
		if err == nil && num > 0 {
			points = uint(num)
			break
		}

		// 錯誤訊息
		pterm.Error.Println("Invalid input. Please enter a positive integer.")

	}

	return points
}

func getReason() string {
	var reason string
	options := []string{
		"加班加到爆",
		"同事太雷",
		"老闆/主管太雞歪",
		"Other",
	}

	reason, _ = pterm.DefaultInteractiveSelect.WithOptions(options).Show("Select a reason for adding points")

	if reason == "Other" {
		for {
			input, _ := pterm.DefaultInteractiveTextInput.Show("Enter your own reason (1-50 characters)")
			input = strings.TrimSpace(input)

			if len(input) >= 1 && len(input) <= 50 {
				reason = input
				break
			}

			// 錯誤訊息
			pterm.Error.Println("Invalid input. Please enter a reason between 1 and 50 characters.")

		}
	}

	return reason
}
