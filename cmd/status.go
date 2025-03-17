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
	Short: "Show the current points and goal status",
	Long: `This command displays the current total points of the active user along with their goal points.
It also includes a progress bar showing how close the user is to reaching their goal.
Example usage:

  $ idfc-tracker status

This will show the user's total points, their goal, and a visual representation of the progress toward their goal.
For example, if the user has 80 points and the goal is 100, the progress bar will show 80% completion.`,
	Run: func(cmd *cobra.Command, args []string) {

		ctx := cmd.Context()
		dbQueries := ctx.Value("dbQueries").(*database.Queries)

		activeUser, err := database.GetActiveUser(ctx, dbQueries)

		if err != nil {
			pterm.Warning.Println("No user found. Run init to create a user and set your goal.")
			return
		}

		userPoint, err := dbQueries.GetPointByUserID(cmd.Context(), activeUser.ID)
		if err != nil {
			pterm.Error.Printf("Database error: %v\n", err)
			return
		}

		percent := float64(userPoint.Total) / float64(userPoint.Goal) * 100

		pterm.DefaultBasicText.Printf("Hi %s!\n", activeUser.Name)
		showProgress(int(userPoint.Total), int(userPoint.Goal))
		pterm.DefaultBasicText.Printf("你目前的狀態為：%s\n", getClosestStatus(int(percent)))

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

func showProgress(current, total int) {
	p, _ := pterm.DefaultProgressbar.WithTotal(total).WithTitle("目前進度").WithShowElapsedTime(false).Start()

	for i := 0; i < current; i++ {
		p.Increment()
	}

	p.Stop()
}
