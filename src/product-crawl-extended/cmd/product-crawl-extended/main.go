package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"product-crawl-extended/internal/stage_one"
)

func main() {
	var rootCmd = &cobra.Command{Use: "spider"}
	var cmd = &cobra.Command{
		Use:   "cmd",
		Short: "cmd short",
		Long:  "cmd long",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("cmd run")
			stage_one.Init()
		},
	}
	rootCmd.AddCommand(cmd)
	rootCmd.Execute()
}
