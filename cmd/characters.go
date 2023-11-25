/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/radekska/wc/pkg"
	"github.com/spf13/cobra"
)

// charactersCmd represents the characters command
var charactersCmd = &cobra.Command{
	Use:   "characters",
	Short: "count by characters",
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
		err, words := pkg.CountCharacters(file)
		if err != nil {
			cmd.PrintErr(err)
			return
		}
		cmd.Println(words)
	},
}

func init() {
	rootCmd.AddCommand(charactersCmd)
}
