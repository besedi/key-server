package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/besedi/key-server/internal/srv"
	"github.com/spf13/cobra"
)

const maxLen = 100_000

var rootCmd = &cobra.Command{
	Use:   "key-server",
	Short: "Key Server",
	Run: func(cmd *cobra.Command, args []string) {
		size, _ := cmd.PersistentFlags().GetInt("max-size")
		port, _ := cmd.PersistentFlags().GetInt("srv-port")
		if size > maxLen {
			fmt.Println("Max size should be not greater then " + strconv.Itoa(maxLen))
			os.Exit(1)
		}
		if port <= 1023 || port >= 65536 {
			fmt.Println("Please specify port between 1024-65535")
			os.Exit(1)
		}
		srv.Serve(size, port)
	},
}

func main() {
	var maxSize int
	var srvPort int
	rootCmd.PersistentFlags().IntVarP(&maxSize, "max-size", "s", 1024, "maximum key size (default 1024)")
	rootCmd.PersistentFlags().IntVarP(&srvPort, "srv-port", "p", 1123, "server listening port (default 1123)")

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}

}
