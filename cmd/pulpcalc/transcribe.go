package pulpcalc

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
)

var (
	file string
)

type TranscribeResponse interface{}

var transcribeCommand = &cobra.Command{
	Use:   "transcribe",
	Short: "Transcribe a video or audio file, and score it",
	Long:  `Transcribe a video or audio file, and score it`,
	Run: func(cmd *cobra.Command, args []string) {

		res, err := cfg.OpenAIClient.CreateTranscription(
			context.Background(),
			openai.AudioRequest{
				Model:    openai.Whisper1,
				FilePath: file,
			},
		)

		if err != nil {
			fmt.Println(err)

			return
		}

		fmt.Println(res.Text)
	},
}

func init() {
	transcribeCommand.Flags().StringVarP(&file, "file", "f", "", "The path of the file to transcribe")

	rootCmd.AddCommand(transcribeCommand)
}
