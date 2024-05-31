package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "root"}
	var cmd = &cobra.Command{
		Use:   "cmd",
		Short: "cmd short",
		Long:  "cmd long",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("cmd run")
		},
	}
	rootCmd.AddCommand(cmd)
	rootCmd.Execute()
}
