/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/radekska/wc/pkg"
	"github.com/spf13/cobra"
)

// linesCmd represents the lines command
var linesCmd = &cobra.Command{
	Use:   "lines",
	Short: "count by lines",
	Run:   count(pkg.CountLines),
}

func init() {
	rootCmd.AddCommand(linesCmd)
}
