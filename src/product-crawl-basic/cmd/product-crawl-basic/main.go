package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func enableLifeCycleEvents() chromedp.ActionFunc {
	return func(ctx context.Context) error {
		err := page.Enable().Do(ctx)
		if err != nil {
			return err
		}
		err = page.SetLifecycleEventsEnabled(true).Do(ctx)
		if err != nil {
			return err
		}
		return nil
	}
}

func navigateAndWaitForLoad(url string, eventName string) chromedp.ActionFunc {
	return func(ctx context.Context) error {
		_, _, _, err := page.Navigate(url).Do(ctx)
		if err != nil {
			return err
		}
		return waitFor(ctx, eventName)
	}
}

func waitFor(ctx context.Context, eventName string) error {
	ch := make(chan struct{})
	cctx, cancel := context.WithCancel(ctx)
	defer cancel()

	chromedp.ListenTarget(cctx, func(ev interface{}) {
		switch e := ev.(type) {
		case *page.EventLifecycleEvent:
			if e.Name == eventName {
				cancel()
				close(ch)
			}
		}
	})

	select {
	case <-ch:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func main() {
	var initialOptions = append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("disable-gpu", false),
		chromedp.Flag("headless", false),
		chromedp.Flag("no-sandbox", true),
	)
	// create context
	startCtx, _ := chromedp.NewExecAllocator(context.Background(), initialOptions...)
	ctx, cancel := chromedp.NewContext(startCtx)
	defer cancel()

	var html string
	var iframeContent string
	err := chromedp.Run(ctx,
		// visit the target page
		chromedp.Tasks{
			navigateAndWaitForLoad("http://localhost:8000/root.html", "networkIdle"),
		},
		chromedp.WaitReady("iframe", chromedp.ByQuery),
		// Switch to the iframe context
		chromedp.Tasks{
			chromedp.ActionFunc(func(ctx context.Context) error {
				// Get the iframe node ID
				var iframeNodeID chromedp.Node
				if err := chromedp.NodeIDs("iframe", &iframeNodeID).Do(ctx); err != nil {
					return err
				}

				// Switch to the iframe
				if err := chromedp.Frame(iframeNodeID).Do(ctx); err != nil {
					return err
				}

				// Wait until the iframe body is fully loaded
				if err := chromedp.WaitReady("body").Do(ctx); err != nil {
					return err
				}

				// Get the iframe content
				if err := chromedp.Evaluate(`document.documentElement.outerHTML`, &iframeContent).Do(ctx); err != nil {
					return err
				}

				return nil
			}),
		},
		// get the outer HTML of the page
		// chromedp.ActionFunc(func(ctx context.Context) error {
		// 	node, err := dom.GetDocument().Do(ctx)
		// 	if err != nil {
		// 		return err
		// 	}
		// 	html, err = dom.GetOuterHTML().WithNodeID(node.NodeID).Do(ctx)
		// 	return err
		// }),
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(html)
}
