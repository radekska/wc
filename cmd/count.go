package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func count(countFunc func([]string, string) error) func(*cobra.Command, []string) {
	return func(cmd *cobra.Command, args []string) {
		files, err := cmd.Flags().GetStringArray("file")
		if err != nil {
			cmd.PrintErr(err)
			os.Exit(1)
		}
		output, err := cmd.Flags().GetString("output")
		if err != nil {
			cmd.PrintErr(err)
			os.Exit(1)
		}
		fmt.Println(files, output)
		if output == "" {
			cmd.Help()
			return
		}

		if len(files) == 0 {
			cmd.Help()
			return
		}
		err = countFunc(files, output)
		if err != nil {
			cmd.PrintErr(err)
			return
		}
	}
}
