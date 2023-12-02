package cmd

import (
	"os"

	"github.com/radekska/wc/pkg"
	"github.com/spf13/cobra"
)

func count(countFunc func([]string) (error, []pkg.Counted)) func(*cobra.Command, []string) {
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
		err, counted := countFunc(files)
		if err != nil {
			cmd.PrintErr(err)
			return
		}
		for _, c := range counted {
			cmd.Printf("%s: %d\n", c.File, c.Count)
		}
	}
}
