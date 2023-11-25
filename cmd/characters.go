/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/radekska/wc/pkg"
	"github.com/spf13/cobra"
)

// charactersCmd represents the characters command
var charactersCmd = &cobra.Command{
	Use:   "characters",
	Short: "count by characters",
	Run:   count(pkg.CountCharacters),
}

func init() {
	rootCmd.AddCommand(charactersCmd)
}
