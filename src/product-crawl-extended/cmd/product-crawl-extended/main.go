package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"product-crawl-extended/internal/first_automation"
	"product-crawl-extended/internal/second_step"
	//	"product-crawl-extended/internal/third_automation"
)

func setup() {
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.JSONFormatter{})
}

func setupLogger() *log.Logger {
	logger := log.New()
	logger.SetFormatter(&log.TextFormatter{
		ForceColors: true, // Enable colors in the output
	})
	return logger
}

func main() {
	var rootCmd = &cobra.Command{Use: "spider"}
	logger := setupLogger()
	var cmd = &cobra.Command{
		Use:   "crawl",
		Short: "crawl basic set automation1",
		Long:  "This is the first crawl command and does very basic set automation.",
		Run: func(cmd *cobra.Command, args []string) {
			logger.Debug("args", args)
		},
	}
	var basic = &cobra.Command{
		Use:   "basic",
		Short: "crawl basic set automation",
		Long:  "This is the first crawl command and does basic set automation.",
		Run: func(cmd *cobra.Command, args []string) {
			logger.Info("cmd running basic automation")
			logger.Debug("args", args)
			first_automation.Crawl()
		},
	}
	var extent_one = &cobra.Command{
		Use:   "extend_one",
		Short: "crawl second stage automation",
		Long:  "This is the second crawl command and does slightly more automation.",
		Run: func(cmd *cobra.Command, args []string) {
			logger.Info("cmd running extended one automation")
			logger.Debug("args", args)
			second_step.Crawl()
		},
	}
	var extent_two = &cobra.Command{
		Use:   "extend_events",
		Short: "crawl second stage automation with events",
		Long:  "This is the third crawl command and does slightly more automation with events included.",
		Run: func(cmd *cobra.Command, args []string) {
			logger.Info("cmd running extended one automation")
			logger.Debug("args", args)
			first_automation.Crawl()
		},
	}
	rootCmd.AddCommand(cmd)
	cmd.AddCommand(basic, extent_one, extent_two)
	rootCmd.Execute()
}
