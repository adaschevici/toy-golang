package seventh_step

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"log"
)

func Crawl() {
	var initialOptions = append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("disable-gpu", false),
		chromedp.Flag("headless", false),
	)
	startCtx, _ := chromedp.NewExecAllocator(context.Background(), initialOptions...)
	// initialize a controllable Chrome instance
	ctx, cancel := chromedp.NewContext(startCtx)
	// to release the browser resources when
	// it is no longer needed
	defer cancel()
	var urlstr = "/form"
	if err := chromedp.Run(ctx,
		// chromedp.Click(`button#download`),
		chromedp.Navigate(fmt.Sprintf("http://localhost:5000%s", urlstr)),
	); err != nil {
		log.Fatal("Error while trying to grab product items.", err)
	}
	fmt.Println("This is the seventh step.")
}
