package fifth_step

import (
	"context"
	// "github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"log"
	"time"
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
	if err := chromedp.Run(ctx,
		chromedp.Navigate("http://localhost:5000/form"),
		chromedp.WaitVisible(`#name`),
		chromedp.SendKeys(`#name`, "John Doe"),
		chromedp.SendKeys(`input[name='City']`, "john.d.@wtf.com"),
		chromedp.SendKeys(`input[name='Country']`, "United States"),
		chromedp.Submit(`form`),
		chromedp.Sleep(5*time.Second),
	); err != nil {

		log.Fatal("Error while trying to grab product items.", err)
	}
}
