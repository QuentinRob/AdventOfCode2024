/*
Package cmd
Copyright © 2024 Quentin ROBCIS
*/
package cmd

import (
	"fmt"
	"github.com/QuentinRob/AdventOfCode2024/internal/historianHysteria"
	"github.com/spf13/cobra"
	"log"
)

// historianHysteriaCmd represents the historianHysteria command
var historianHysteriaCmd = &cobra.Command{
	Use:     "historianHysteria",
	Aliases: []string{"1"},
	Short:   "What is the total distance between your lists ?",
	Long: `What is the total distance between your lists ?

Compute the total distance between two lists.`,

	Run: func(cmd *cobra.Command, args []string) {
		inputPath, _ := cmd.Flags().GetString("input")
		hysteria := historianHysteria.NewHistorianHysteria(inputPath)
		totalDistance, err := hysteria.GetTotalDistance()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Total distance: \t", totalDistance)

		similarityScore, err := hysteria.GetSimilarityScore()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Similarity score: \t", similarityScore)
	},
}

func init() {
	rootCmd.AddCommand(historianHysteriaCmd)

	historianHysteriaCmd.Flags().StringP("input", "i", "./inputs/day1.txt", "Input file path")
}
