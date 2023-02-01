package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/baribari2/pulp-calculator/tree"
	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func main() {
	cfg := initConfig()

	// if DictServer == "" {
	// 	DictServer = "dict://dict.dict.org/"
	// } else {
	// 	s, err := dict.Dial("tcp", DictServer)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	cfg.DictServer = s
	// }

	// err := reddit.AppOnlyRequest(cfg)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// tweets, err := twitter.SearchTweets(cfg.TwitterBearerToken, "elon")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// token, err := twitter.AppOnlyRequest(cfg.TwitterApiKey, cfg.TwitterApiSecret)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// cfg.TwitterBearerToken = token

	// tweet, err := twitter.GetTweetById(cfg.TwitterBearerToken, "1597450870819811328")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Printf("%v", tweet)
	spinner := spinner.New(spinner.CharSets[35], 100*time.Millisecond, spinner.WithColor("green"))
	spinner.Suffix = " Simulating thread...\n\n"

	spinner.Start()

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
	ctable, ttable := tree.SimulateThread(cfg, l, time.Duration(1*time.Second), time.Duration(10*time.Second), 1)

	ctable.WithHeaderFormatter(hf).WithFirstColumnFormatter(cf)
	ttable.WithHeaderFormatter(hf).WithFirstColumnFormatter(cf)

	fmt.Printf("Comment table: \n")
	ctable.Print()

	fmt.Println()

	fmt.Printf("Time table: \n")
	ttable.Print()

	spinner.Stop()

	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		l.Render(w)
	})

	fmt.Println("\n📣 Listening on port 8080 (navigate to http://localhost:8080 to view the chart)")
	http.ListenAndServe(":8080", nil)

}
