/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/y3933y3933/idfc-tracker/internal/database"
)

var ofJson bool
var ofCsv bool

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export the history data to a file",
	Long: `This command allows the user to export their history data in either JSON or CSV format.
The user can choose the format using the --json or --csv flag, and the file will be saved in the current directory.
Example usage:

  $ idfc-tracker export --json
  $ idfc-tracker export --csv

This will export all the history records in the selected format, providing a quick way to back up or analyze the data.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		dbQueries := ctx.Value("dbQueries").(*database.Queries)

		activeUser, err := database.GetActiveUser(ctx, dbQueries)

		if err != nil {
			pterm.Warning.Println("No user found. Run init to create a user and set your goal.")
			return
		}
		history, err := dbQueries.GetHistoryByUserID(ctx, activeUser.ID)
		if err != nil {
			log.Fatalf("get history err:%v\n", err)
		}
		if len(history) == 0 {
			pterm.Info.Println("No history records found.")
			return
		}

		format := getFormat()
		fileName := fmt.Sprintf("history_%s.%s", time.Now().Format("20060102_150405"), format)

		if format == "json" {
			err = exportJSON(fileName, history)
		} else {
			err = exportCsv(fileName, history)
		}
		if err != nil {
			pterm.Error.Println("Export file Fail")
			return
		}

	},
}

func init() {
	rootCmd.AddCommand(exportCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// exportCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// exportCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// exportCmd.Flags().StringP("format", "f")
	exportCmd.Flags().BoolVar(&ofJson, "json", false, "Output in JSON")
	exportCmd.Flags().BoolVar(&ofCsv, "csv", false, "Output in CSV")
	exportCmd.MarkFlagsOneRequired("json", "csv")
	exportCmd.MarkFlagsMutuallyExclusive("json", "csv")
}

func getFormat() string {
	var format string
	if ofCsv {
		format = "csv"
	} else {
		format = "json"
	}
	return format

}

func exportCsv(fileName string, history []database.History) error {
	f, err := os.Create(filepath.Clean(fileName))
	if err != nil {
		return err
	}
	defer f.Close()
	writer := csv.NewWriter(f)

	header := []string{"ID", "Points", "Reason", "Date"}
	if err := writer.Write(header); err != nil {
		return err
	}
	for _, record := range history {
		dateStr := record.CreatedAt.Format(time.DateOnly)
		row := []string{
			strconv.Itoa(int(record.ID)),
			strconv.Itoa(int(record.Point)),
			record.Reason,
			dateStr,
		}
		if err := writer.Write(row); err != nil {
			return err
		}
	}
	writer.Flush()
	return nil
}

func exportJSON(fileName string, content interface{}) error {

	f, err := os.Create(filepath.Clean(fileName))
	if err != nil {
		return err
	}
	defer f.Close()

	jsonData, err := json.MarshalIndent(content, "", "  ")
	if err != nil {
		return err
	}

	_, err = f.Write(jsonData)
	return err
}
