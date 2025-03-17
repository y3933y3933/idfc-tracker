/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"sort"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/y3933y3933/idfc-tracker/internal/database"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
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
		}

		userPoint, err := dbQueries.GetPointByUserID(cmd.Context(), activeUser.ID)
		if err != nil {
			pterm.Error.Printf("Database error: %v\n", err)
			return
		}

		percent := (userPoint.Total) / (userPoint.Goal)

		pterm.DefaultBasicText.Printf("Hi %s!\n", activeUser.Name)
		pterm.DefaultBasicText.Printf("你的目標點數是: %d\n", userPoint.Goal)
		pterm.DefaultBasicText.Printf("你累積的離職點數是: %d\n", userPoint.Total)
		pterm.DefaultBasicText.Printf("----------------------------------\n")
		pterm.DefaultBasicText.Printf("你目前的狀態為：%s", getClosestStatus(int(percent)))

	},
}

func init() {
	rootCmd.AddCommand(statusCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// statusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// statusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getStatusMap() map[int]string {
	return map[int]string{
		0:   pterm.Gray("我願意為公司付出心臟"),
		20:  pterm.Blue("好像哪裡怪怪的"),
		40:  pterm.Cyan("心一跳，對工作的愛就開始煎熬"),
		60:  pterm.Yellow("履歷該更新了"),
		80:  pterm.LightRed("快爆炸了！"),
		99:  pterm.Red("忍無可忍"),
		100: pterm.BgRed.Sprint("今天就提離職"),
		101: pterm.BgYellow.Sprint("早該走了吧怎麼還在"),
	}

}

func getClosestStatus(percentage int) string {
	statusMap := getStatusMap()
	keys := make([]int, 0, len(statusMap))
	for k := range statusMap {
		keys = append(keys, k)
	}
	sort.Ints(keys) // 確保 keys 是遞增排序

	closest := keys[0]
	for _, key := range keys {
		if percentage >= key {
			closest = key
		} else {
			break
		}
	}

	return statusMap[closest]
}
