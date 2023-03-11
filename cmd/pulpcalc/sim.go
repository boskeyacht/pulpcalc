package pulpcalc

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/baribari2/pulp-calculator/grpc"
	"github.com/baribari2/pulp-calculator/simulator"
	"github.com/baribari2/pulp-calculator/simulator/sets"
	"github.com/fatih/color"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/spf13/cobra"
)

var (
	tick   int64
	length int64
	freq   int64
	users  int64
)

// TODO: change SimulateThread params in api
var simCmd = &cobra.Command{
	Use:     "sim",
	Short:   "Simulate a thread",
	Long:    `Simulate a thread over time, given a starting tick, end tick, and frequency`,
	Aliases: []string{"s"},
	Run: func(cmd *cobra.Command, args []string) {
		tick := time.Duration(time.Duration(tick) * time.Second)
		leng := time.Duration(time.Duration(length) * time.Second)
		fmt.Printf("\x1b[32m%s\x1b[0m", " Simulating thread...\n\n")

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
		tree, ctable, ttable, err := simulator.SimulateThread(cfg, l, users, tick, leng, freq)
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

var (
	setFile string
)

var setsCmd = &cobra.Command{
	Use:     "sets",
	Short:   "Simulate a debate based on unique sets of users",
	Long:    "Simulate a debate based on unique sets of users",
	Aliases: []string{"s"},
	Run: func(cmd *cobra.Command, args []string) {
		debates := make([]*simulator.Debate, 0)
		s, err := sets.NewSimulationSetsFromFile(setFile)
		if err != nil {
			fmt.Println(err)

			return
		}

		if len(s) == 1 {
			fmt.Printf("\x1b[32m%s\x1b[0m", fmt.Sprintf("Simulating 1 %s debate...\n\n", s[0].GetSimulationType().String()))
		} else {
			fmt.Printf("\x1b[32m%s\x1b[0m", fmt.Sprintf("Simulating %d debates...\n\n", len(s)))
		}

		wg := &sync.WaitGroup{}
		errChan := make(chan error)
		debateChan := make(chan *simulator.Debate, len(s))

		for _, set := range s {
			wg.Add(1)

			go set.RunSimulation(wg, cfg, debateChan, errChan)
		}

		wg.Wait()
		close(debateChan)
		close(errChan)

		for debate := range debateChan {
			fmt.Println(debate)

			debates = append(debates, debate)
		}

		for err := range errChan {
			fmt.Println(err)
		}
	},
}

func init() {
	simCmd.Flags().Int64VarP(&tick, "tick", "t", 0, "The starting tick")
	simCmd.Flags().Int64VarP(&length, "len", "l", 0, "The total runtime (in seconds)")
	simCmd.Flags().Int64VarP(&freq, "freq", "f", 0, "The frequency of the simulation")
	simCmd.Flags().Int64VarP(&users, "users", "u", 0, "The number of users to simulate")

	setsCmd.Flags().StringVarP(&setFile, "file", "f", "", "The file to read the simulation set from")

	simCmd.AddCommand(setsCmd)

	rootCmd.AddCommand(simCmd)
}
