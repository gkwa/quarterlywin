package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var redisCmd = &cobra.Command{
	Use:   "redis",
	Short: "Print redis connection URL",
	Run: func(cmd *cobra.Command, args []string) {
		url := os.Getenv("REDIS_URL")
		if url == "" {
			fmt.Println("REDIS_URL environment variable is not set")
			return
		}
		fmt.Printf("Our redis url is: %s\n", url)
		fmt.Printf("the first character of our variable is: %c\n", url[0])
		fmt.Printf("the last character is: %c\n", url[len(url)-1])
		fmt.Printf("the value length is: %d\n", len(url))
	},
}

func init() {
	rootCmd.AddCommand(redisCmd)
}
