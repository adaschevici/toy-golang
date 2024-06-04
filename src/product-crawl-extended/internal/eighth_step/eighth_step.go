package eighth_step

import (
	"context"
	"fmt"
	"log"

	// "github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

func Crawl() {
	var initialOptions = append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36"),
		chromedp.Flag("disable-gpu", false),
		chromedp.Flag("headless", false),
		chromedp.Flag("v", "1"), // Verbose logging
		chromedp.Flag("vmodule", "*=3"),
	)
	startCtx, _ := chromedp.NewExecAllocator(context.Background(), initialOptions...)
	// initialize a controllable Chrome instance
	ctx, cancel := chromedp.NewContext(startCtx)
	// to release the browser resources when
	// it is no longer needed
	defer cancel()
	// var urlstr string
	var foundIp string
	if err := chromedp.Run(ctx,
		chromedp.Navigate("https://whatismyipaddress.com/"),
		chromedp.Text(`span#ipv4 > a`, &foundIp, chromedp.ByQueryAll),
	); err != nil {
		log.Fatal("Error while trying to grab current ip.", err)
	}
	fmt.Println("IP Address: ", foundIp)
}
