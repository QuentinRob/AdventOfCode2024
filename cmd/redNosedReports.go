/*
Copyright © 2024 Quentin ROBCIS
*/
package cmd

import (
	"fmt"
	"github.com/QuentinRob/AdventOfCode2024/internal/redNosedReports"

	"github.com/spf13/cobra"
)

var redNosedReportsCmd = &cobra.Command{
	Use:     "redNosedReports",
	Aliases: []string{"2"},
	Short:   "How many reports are safe?",
	Long: `How many reports are safe?

The engineers are trying to figure out which reports are safe. The Red-Nosed reactor safety systems can only tolerate levels that are either gradually increasing or gradually decreasing. So, a report only counts as safe if both of the following are true:

    The levels are either all increasing or all decreasing.
    Any two adjacent levels differ by at least one and at most three.

In the example above, the reports can be found safe or unsafe by checking those rules:

    7 6 4 2 1: Safe because the levels are all decreasing by 1 or 2.
    1 2 7 8 9: Unsafe because 2 7 is an increase of 5.
    9 7 6 2 1: Unsafe because 6 2 is a decrease of 4.
    1 3 2 4 5: Unsafe because 1 3 is increasing but 3 2 is decreasing.
    8 6 4 4 1: Unsafe because 4 4 is neither an increase or a decrease.
    1 3 6 7 9: Safe because the levels are all increasing by 1, 2, or 3.
`,
	Run: func(cmd *cobra.Command, args []string) {
		inputPath, _ := cmd.Flags().GetString("input")

		reports := redNosedReports.NewRedNosedReports(inputPath)

		fmt.Println("Total safe reports: \t", reports.GetTotalSafeReports())
		fmt.Println("Total safe problem dampener reports: \t", reports.GetTotalSafeProblemDampenerReports())
	},
}

func init() {
	rootCmd.AddCommand(redNosedReportsCmd)

	redNosedReportsCmd.Flags().StringP("input", "i", "./inputs/day2.txt", "Input file path")

}
