package ninth_step

import (
	"context"
	"fmt"
	"log"
	// "time"

	"github.com/chromedp/cdproto/dom"
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
	var html string
	if err := chromedp.Run(ctx,
		chromedp.Navigate("https://www.g2.com/products/jira/reviews/"),
		// chromedp.Sleep(60*time.Second),
		// we have an iframe, so we need to switch to it before we can interact with it
		chromedp.WaitVisible(`label.cb-lb > input`),
		chromedp.Click(`label.cb-lb > input`),
		chromedp.ActionFunc(func(ctx context.Context) error {
			rootNode, err := dom.GetDocument().Do(ctx)
			if err != nil {
				return err
			}
			html, err = dom.GetOuterHTML().WithNodeID(rootNode.NodeID).Do(ctx)
			return err
		}),
	); err != nil {
		log.Fatal("Error while trying to grab current ip.", err)
		return
	}

	fmt.Println("HTML: ", html)
}
