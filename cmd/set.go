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

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		goal, _ := cmd.Flags().GetUint("goal")

		ctx := cmd.Context()
		dbQueries := ctx.Value("dbQueries").(*database.Queries)

		activeUser, err := database.GetActiveUser(ctx, dbQueries)

		if err != nil {
			pterm.Warning.Println("No user found. Run init to create a user and set your goal.")
		}

		err = dbQueries.UpdateGoalByUserID(ctx, database.UpdateGoalByUserIDParams{
			Goal:   int64(goal),
			UserID: activeUser.ID,
		})

		if err != nil {
			log.Fatalf("Database error: %v\n", err)
		}

		pterm.Success.Printf("User: %s\nGoal Set: %d points\n", activeUser.Name, goal)
	},
}

func init() {
	rootCmd.AddCommand(setCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	setCmd.Flags().UintP("goal", "g", 0, "Set the goal points.")
	setCmd.MarkFlagRequired("goal")

}
