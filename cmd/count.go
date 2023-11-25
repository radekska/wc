package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func count(countFunc func(string) (error, int)) func(*cobra.Command, []string) {
	return func(cmd *cobra.Command, args []string) {
		file, err := cmd.Flags().GetString("file")
		if err != nil {
			cmd.PrintErr(err)
			os.Exit(1)
		}
		if file == "" {
			cmd.Help()
			return
		}
		err, words := countFunc(file)
		if err != nil {
			cmd.PrintErr(err)
			return
		}
		cmd.Println(words)
	}
}
