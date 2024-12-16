/*
Package cmd
Copyright © 2024 Quentin ROBCIS
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// historianHysteriaCmd represents the historianHysteria command
var historianHysteriaCmd = &cobra.Command{
	Use:     "historianHysteria",
	Aliases: []string{"1"},
	Short:   "What is the total distance between your lists ?",
	Long: `What is the total distance between your lists ?

Compute the total distance between two lists.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(historianHysteriaCmd)
}
