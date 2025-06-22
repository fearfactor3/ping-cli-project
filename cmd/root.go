package cmd

import (
	"fmt"
	"os"
	"sync"

	"github.com/fearfactor3/ping-cli-project/pkg/ping"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ping-cli [URL]",
	Short: "Ping services/endpoints and report their status",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var wg sync.WaitGroup
		results := make(chan string, len(args))
		
		for _, url := range args {
			wg.Add(1)
			go func(endpoint string) {
				defer wg.Done()
				results <- ping.CheckHTTP(endpoint)
			}(url)
		}

		wg.Wait()
		close(results)

		for result := range results {
			fmt.Println(result)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
