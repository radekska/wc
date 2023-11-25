/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/radekska/wc/pkg"
	"github.com/spf13/cobra"
)

// wordsCmd represents the words command
var wordsCmd = &cobra.Command{
	Use:   "words",
	Short: "count by words",
	Run:   count(pkg.CountWords),
}

func init() {
	rootCmd.AddCommand(wordsCmd)
}
