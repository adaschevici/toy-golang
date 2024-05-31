package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"product-crawl-extended/internal/stage_one"
)

func main() {
	var rootCmd = &cobra.Command{Use: "spider"}
	var cmd = &cobra.Command{
		Use:   "crawl",
		Short: "crawl basic set automation",
		Long:  "This is the first crawl command and does very basic set automation.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("cmd run")
			fmt.Println("args", args)
			stage_one.Init()
		},
	}
	rootCmd.AddCommand(cmd)
	rootCmd.Execute()
}
