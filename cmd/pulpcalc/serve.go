package pulpcalc

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/baribari2/pulpcalc/api"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:     "serve",
	Short:   "Start the server",
	Long:    `Start the server`,
	Aliases: []string{"sv"},
	Run: func(cmd *cobra.Command, args []string) {
		l := charts.NewLine()
		l.SetGlobalOptions(
			charts.WithTitleOpts(opts.Title{
				Title:    "Thread score",
				Subtitle: "The score of a thread over time",
			}),
		)

		rt := api.SetupRouter(cfg, l)

		go rt.Run(":8080")

		fmt.Println("Simulation api server started on port 8080")

		// Block until a SIGING signal is received
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
