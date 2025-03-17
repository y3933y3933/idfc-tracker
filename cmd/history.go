/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"strconv"
	"time"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/y3933y3933/idfc-tracker/internal/database"
)

var (
	start string
	end   string
)

// historyCmd represents the history command
var historyCmd = &cobra.Command{
	Use:   "history",
	Short: "View the user's point history",
	Long: `This command displays the history of all points added by the active user.
The data will be shown in a table format, listing the date, points added, and reason for each entry.
You can also filter the results by specifying a date range using the --start and --end flags.
Example usage:

  $ idfc-tracker history
  $ idfc-tracker history --start "2025-01-01" --end "2025-01-31"

This will show the user's history between the specified start and end dates, making it easier to review recent activity.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		dbQueries := ctx.Value("dbQueries").(*database.Queries)

		activeUser, err := database.GetActiveUser(ctx, dbQueries)

		if err != nil {
			pterm.Warning.Println("No user found. Run init to create a user and set your goal.")
			return
		}

		startDate, endDate := parseDateRange(start, end)

		pterm.Info.Println("Fetching history...")
		pterm.Print("\n")

		history, err := dbQueries.GetHistoryByUserIDAndDateRange(ctx, database.GetHistoryByUserIDAndDateRangeParams{
			UserID:        activeUser.ID,
			FromCreatedAt: startDate,
			ToCreatedAt:   endDate,
		})
		if err != nil {
			pterm.Error.Println("Failed to fetch history.")
			log.Fatal(err)
		}
		if err != nil {
			log.Fatalf("get history err:%v\n", err)
		}
		if len(history) == 0 {
			pterm.Info.Println("No history records found.")
			return
		}

		// 設定表格標頭
		tableData := [][]string{
			{"ID", "Points", "Reason", "Date"},
		}

		// 逐筆加入資料
		for _, record := range history {
			dateStr := record.CreatedAt.Format(time.DateOnly) // 轉換為易讀格式
			tableData = append(tableData, []string{
				strconv.Itoa(int(record.ID)),
				strconv.Itoa(int(record.Point)),
				record.Reason,
				dateStr,
			})
		}

		// 使用 PTerm 印出表格
		pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(tableData).Render()

		// 顯示成功訊息
		pterm.Success.Println("History displayed successfully.")

	},
}

func init() {
	rootCmd.AddCommand(historyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// historyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// historyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	historyCmd.Flags().StringVar(&start, "start", "", "Start date (format: YYYY-MM-DD)")
	historyCmd.Flags().StringVar(&end, "end", "", "End date (format: YYYY-MM-DD)")

}

func parseDateRange(start string, end string) (time.Time, time.Time) {
	const dateFormat = time.DateOnly

	startTime := time.Now().AddDate(0, 0, -365)
	if start != "" {
		parsedFrom, err := time.Parse(dateFormat, start)
		if err == nil {
			startTime = parsedFrom
		} else {
			pterm.Warning.Println("Invalid --from date format. Using default (365 days ago).")
		}
	}

	endTime := time.Now()
	if end != "" {
		parsedTo, err := time.Parse(dateFormat, end)
		if err == nil {
			endTime = parsedTo
		} else {
			pterm.Warning.Println("Invalid --to date format. Using default (today).")
		}
	}

	return startTime, endTime
}
