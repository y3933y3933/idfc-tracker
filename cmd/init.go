/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"strings"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
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
		// ask name
		var name string
		for {
			name, _ = pterm.DefaultInteractiveTextInput.Show("What is your name? (in English)")
			name = strings.TrimSpace(name)

			if name != "" {
				break
			}
			pterm.Warning.Println("Name cannot be empty. Please enter your name.")
		}

		// ask goal
		goalOptions := []string{"10", "20", "30", "40", "50"}
		selectedGoal, _ := pterm.DefaultInteractiveSelect.WithOptions(goalOptions).WithDefaultOption("10").Show("Please select your goal")
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
