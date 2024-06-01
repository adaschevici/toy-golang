package main

import (
	"github.com/spf13/cobra"
	log "log/slog"
	// "os"
	"product-crawl-extended/internal/first_automation"
)

func main() {
	var rootCmd = &cobra.Command{Use: "spider"}
	var cmd = &cobra.Command{
		Use:   "crawl",
		Short: "crawl basic set automation1",
		Long:  "This is the first crawl command and does very basic set automation.",
		Run: func(cmd *cobra.Command, args []string) {
			log.Debug("args", args)
		},
	}
	var basic = &cobra.Command{
		Use:   "basic",
		Short: "crawl basic set automation",
		Long:  "This is the first crawl command and does basic set automation.",
		Run: func(cmd *cobra.Command, args []string) {
			log.Info("cmd running basic automation")
			log.Debug("args", args)
			first_automation.Crawl()
		},
	}
	rootCmd.AddCommand(cmd)
	cmd.AddCommand(basic)
	rootCmd.Execute()
}
