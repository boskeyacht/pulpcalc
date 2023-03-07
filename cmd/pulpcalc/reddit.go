package pulpcalc

import (
	"fmt"

	"github.com/baribari2/pulp-calculator/reddit"
	"github.com/spf13/cobra"
)

var (
	query string
)

var redditCmd = &cobra.Command{
	Use:   "reddit",
	Short: "Calculate the score of a reddit thread",
	Long:  "Calculate the score of a reddit thread over time given a search query",
	Run: func(cmd *cobra.Command, args []string) {
		// _, err := reddit.AppOnlyRequest(cfg)
		// if err != nil {
		// 	fmt.Println(err.Error())

		// 	return
		// }

		_, err := reddit.SearchCMVSubreddit(cfg, query)
		if err != nil {
			fmt.Println(err.Error())

			return
		}

	},
}

func init() {
	redditCmd.Flags().StringVarP(&query, "query", "q", "", "The query to search for")

	rootCmd.AddCommand(redditCmd)
}
