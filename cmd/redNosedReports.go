/*
Copyright © 2024 Quentin ROBCIS
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var redNosedReportsCmd = &cobra.Command{
	Use:   "redNosedReports",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		inputPath, _ := cmd.Flags().GetString("input")
	},
}

func init() {
	rootCmd.AddCommand(redNosedReportsCmd)

	redNosedReportsCmd.Flags().StringP("input", "i", "./inputs/day2.txt", "Input file path")

}
