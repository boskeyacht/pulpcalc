package pulpcalc

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/baribari2/pulp-calculator/grpc"
	"github.com/baribari2/pulp-calculator/tree"
	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/spf13/cobra"
)

var (
	tick  int64
	len   int64
	freq  int64
	users int64
)

// TODO: change SimulateThread params in api
var simCmd = &cobra.Command{
	Use:     "sim",
	Short:   "Simulate a thread",
	Long:    `Simulate a thread over time, given a starting tick, end tick, and frequency`,
	Aliases: []string{"s"},
	Run: func(cmd *cobra.Command, args []string) {
		tick := time.Duration(time.Duration(tick) * time.Second)
		len := time.Duration(time.Duration(len) * time.Second)
		spinner := spinner.New(spinner.CharSets[35], 100*time.Millisecond, spinner.WithColor("green"))
		spinner.Suffix = " Simulating thread...\n\n"

		// spinner.Start()

		l := charts.NewLine()
		l.SetGlobalOptions(
			charts.WithTitleOpts(opts.Title{
				Title:    "Thread score",
				Subtitle: "The score of a thread over time",
			}),
		)

		hf := color.New(color.FgGreen, color.Underline).SprintfFunc()
		cf := color.New(color.FgYellow).SprintfFunc()

		// Comment table & Time table
		tree, ctable, ttable, err := tree.SimulateThread(cfg, l, users, tick, len, freq)
		if err != nil {
			fmt.Println(err)
		}

		ctable.WithHeaderFormatter(hf).WithFirstColumnFormatter(cf)
		ttable.WithHeaderFormatter(hf).WithFirstColumnFormatter(cf)

		fmt.Printf("Comment table: \n")
		ctable.Print()

		fmt.Println()

		fmt.Printf("Time table: \n")
		ttable.Print()

		// spinner.Stop()

		http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
			l.Render(w)
		})

		fmt.Println("\n📣 Listening on port 8080 (navigate to http://localhost:8080 to view the chart)")
		go func() {
			http.ListenAndServe(":8080", nil)
		}()

		side := grpc.NewGrpcCalcServer(tree)
		go side.Start()

		fmt.Println("\n📣 gRPC server started on port 8081")

		// Block until a SIGING signal is received
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
	},
}

func init() {
	simCmd.Flags().Int64VarP(&tick, "tick", "t", 0, "The starting tick")
	simCmd.Flags().Int64VarP(&len, "len", "l", 0, "The total runtime (in seconds)")
	simCmd.Flags().Int64VarP(&freq, "freq", "f", 0, "The frequency of the simulation")
	simCmd.Flags().Int64VarP(&users, "users", "u", 0, "The number of users to simulate")

	rootCmd.AddCommand(simCmd)
}
