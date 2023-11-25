/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/radekska/wc/pkg"
	"github.com/spf13/cobra"
)

// wordsCmd represents the words command
var wordsCmd = &cobra.Command{
	Use:   "words",
	Short: "count by words",
	Run: func(cmd *cobra.Command, args []string) {
		file, err := cmd.Flags().GetString("file")
		if err != nil {
			cmd.PrintErr(err)
			os.Exit(1)
		}
		if file == "" {
			cmd.Help()
			return
		}
		err, words := pkg.CountWords(file)
		if err != nil {
			cmd.PrintErr(err)
			return
		}
		cmd.Println(words)
	},
}

func init() {
	rootCmd.AddCommand(wordsCmd)
}
