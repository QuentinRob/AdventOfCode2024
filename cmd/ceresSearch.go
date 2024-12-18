/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/QuentinRob/AdventOfCode2024/internal/ceresSearch"

	"github.com/spf13/cobra"
)

// ceresSearchCmd represents the ceresSearch command
var ceresSearchCmd = &cobra.Command{
	Use:     "ceresSearch",
	Aliases: []string{"4"},
	Short:   "How many times does XMAS appear?",
	Long: `How many times does XMAS appear?

This word search allows words to be horizontal, vertical, diagonal, written backwards, or even overlapping other words.
It's a little unusual, though, as you don't merely need to find one instance of XMAS - you need to find all of them.`,
	Run: func(cmd *cobra.Command, args []string) {
		inputPath, _ := cmd.Flags().GetString("input")

		ceres := ceresSearch.NewCeresSearch(inputPath, "XMAS", "MAS")

		fmt.Println("XMAS occurrences found: \t", ceres.CountSearchOccurrences())
		fmt.Println("X-MAS occurrences found: \t", ceres.CountXSearchOccurences())
	},
}

func init() {
	rootCmd.AddCommand(ceresSearchCmd)

	ceresSearchCmd.Flags().StringP("input", "i", "./inputs/day4.txt", "Input file path")

}
