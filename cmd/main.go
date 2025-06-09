package main

import (
	"github.com/besedi/together-ai-home/internal/srv"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "key-server",
	Short: "Key Server",
	Run: func(cmd *cobra.Command, args []string) {
		size, _ := cmd.PersistentFlags().GetString("max-size")
		port, _ := cmd.PersistentFlags().GetString("src-port")
		srv.Serve(size, port)
	},
}

func main() {
	var maxSize string
	var srvPort string
	rootCmd.PersistentFlags().StringVarP(&maxSize, "max-size", "s", "", "maximum key size (default 1024)")
	rootCmd.PersistentFlags().StringVarP(&srvPort, "srv-port", "p", "", "server listening port (default 1123)")
	err := rootCmd.MarkPersistentFlagRequired("max-size")
	if err != nil {
		panic(err)
	}
	err = rootCmd.MarkPersistentFlagRequired("srv-port")
	if err != nil {
		panic(err)
	}
}
