package fourth_step

import (
	"context"
	// "fmt"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"log"
	"os"
	// "time"
)

type Product struct {
	name, price string
}

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

	var script = `
		// scroll down the page 8 times
		let totalHeight = 0;
		let distance = 500;

		// scroll down and then wait for 0.5s
		const scrollInterval = setInterval(() => {
		  const scrollHeight = document.body.scrollHeight;

		  window.scrollBy(0, distance);
		  totalHeight += distance;
		
		  if (totalHeight >= scrollHeight) {
		   clearInterval(scrollInterval)
		  }
		}, 500)
	       `
	var productNodes []*cdp.Node
	var screenshotBuffer []byte

	if err := chromedp.Run(ctx,
		// visit the target page
		chromedp.Navigate("https://scrapingclub.com/exercise/list_infinite_scroll/"),
		chromedp.Evaluate(script, nil),
		chromedp.WaitVisible(".post:nth-child(60)"),
		chromedp.FullScreenshot(&screenshotBuffer, 100),
		chromedp.Screenshot(`.post:nth-child(59)`, &screenshotBuffer, chromedp.NodeVisible),
		chromedp.Nodes(`.post`, &productNodes, chromedp.ByQueryAll),
	); err != nil {
		log.Fatal("Error while trying to grab product items.", err)
	}
	if err := os.WriteFile("screenshot.png", screenshotBuffer, 0644); err != nil {
		log.Fatal("Error while trying to write the screenshot to a file.", err)
	}
}
