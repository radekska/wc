package cmd

import (
	"os"

	"github.com/radekska/wc/pkg"
	"github.com/spf13/cobra"
)

func count(countFunc func([]string) *pkg.CountedFiles) func(*cobra.Command, []string) {
	return func(cmd *cobra.Command, args []string) {
		files, err := cmd.Flags().GetStringArray("file")
		if err != nil {
			cmd.PrintErr(err)
			os.Exit(1)
		}
		if len(files) == 0 {
			cmd.Help()
			return
		}
		counted := countFunc(files)
		for _, c := range counted.Files {
			if c.Err != nil {
				cmd.PrintErrln(c.Err)
				os.Exit(1)
			}
			cmd.Printf("%s: %d\n", c.File, c.Count)
		}
	}
}
