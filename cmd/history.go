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

// historyCmd represents the history command
var historyCmd = &cobra.Command{
	Use:   "history",
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

		history, err := dbQueries.GetHistoryByUserID(ctx, activeUser.ID)
		if err != nil {
			log.Fatalf("get history err:%v\n", err)
		}
		if len(history) == 0 {
			pterm.Info.Println("No history records found.")
			return
		}

		pterm.Info.Println("Fetching history...")
		pterm.Print("\n")

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
}
