package pulpcalc

import (
	"github.com/baribari2/pulp-calculator/common/types"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:  "pulpcalc",
		Long: "pulpcalc is a tool for calculating the score of a thread over time",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	cfg *types.Config
)

func Execute(c *types.Config) {
	cfg = c

	rootCmd.Execute()
}
