package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"product-crawl-extended/internal/stage_one"
)

func main() {
	var rootCmd = &cobra.Command{Use: "spider"}
	var cmd1 = &cobra.Command{
		Use:   "crawl1",
		Short: "crawl basic set automation1",
		Long:  "This is the first crawl command and does very basic set automation.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("cmd run1")
			fmt.Println("args", args)
			stage_one.Crawl()
		},
	}
	var cmd2 = &cobra.Command{
		Use:   "crawl2",
		Short: "crawl basic set automation2",
		Long:  "This is the first crawl command and doesbasic set automation.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("cmd run2")
			fmt.Println("args", args)
			stage_one.Crawl()
		},
	}
	var cmd3 = &cobra.Command{
		Use:   "website",
		Short: "crawl basic set automation3",
		Long:  "This is the first crawl command and does very basic set a.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("cmd runs")
			fmt.Println("args", args)
			stage_one.Crawl()
		},
	}
	rootCmd.AddCommand(cmd1, cmd2)
	cmd1.AddCommand(cmd3)
	rootCmd.Execute()
}
