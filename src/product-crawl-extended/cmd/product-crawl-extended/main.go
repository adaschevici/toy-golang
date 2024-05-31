package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"product-crawl-extended/internal/first_automation"
)

func main() {
	var rootCmd = &cobra.Command{Use: "spider"}
	var cmd = &cobra.Command{
		Use:   "crawl",
		Short: "crawl basic set automation1",
		Long:  "This is the first crawl command and does very basic set automation.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("args", args)
		},
	}
	var basic = &cobra.Command{
		Use:   "basic",
		Short: "crawl basic set automation",
		Long:  "This is the first crawl command and does basic set automation.",
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("cmd running basic automation")
			log.Println("args", args)
			first_automation.Crawl()
		},
	}
	// var cmd3 = &cobra.Command{
	// 	Use:   "website",
	// 	Short: "crawl basic set automation3",
	// 	Long:  "This is the first crawl command and does very basic set a.",
	// 	Run: func(cmd *cobra.Command, args []string) {
	// 		fmt.Println("cmd runs")
	// 		fmt.Println("args", args)
	// 		stage_one.Crawl()
	// 	},
	// }
	// rootCmd.AddCommand(cmd)
	cmd.AddCommand(basic)
	rootCmd.Execute()
}
