package third_step

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"log"
	"time"
)

type Product struct {
	name, price string
}

func Crawl() {
	var products []Product
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

	if err := chromedp.Run(ctx,
		// visit the target page
		chromedp.Navigate("https://scrapingclub.com/exercise/list_infinite_scroll/"),
		chromedp.Evaluate(script, nil),
		chromedp.Sleep(5*time.Second),
		chromedp.Nodes(`.post`, &productNodes, chromedp.ByQueryAll),
	); err != nil {
		log.Fatal("Error while trying to grab product items.", err)
	}
	var name, price string
	for _, node := range productNodes {
		if err := chromedp.Run(ctx,
			chromedp.Text(`h4`, &name, chromedp.ByQuery, chromedp.FromNode(node)),
			chromedp.Text(`h5`, &price, chromedp.ByQuery, chromedp.FromNode(node)),
		); err != nil {
			log.Fatal("Error while trying to grab product items.", err)
		}
		products = append(products, Product{name: name, price: price})
	}
	fmt.Println(products)

}
