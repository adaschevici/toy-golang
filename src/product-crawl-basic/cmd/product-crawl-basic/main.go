package main

import (
	"context"
	// "encoding/json"
	"fmt"
	// "log"
	// "strings"
	"time"

	"github.com/chromedp/cdproto/cdp"
	// "github.com/chromedp/cdproto/dom"
	// "github.com/chromedp/cdproto/page"
	// "github.com/chromedp/cdproto/target"
	"github.com/chromedp/chromedp"
	// "github.com/chromedp/chromedp/runner"
)

func main() {
	var initialOptions = append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("disable-gpu", false),
		chromedp.Flag("headless", false),
	)
	// create context
	startCtx, _ := chromedp.NewExecAllocator(context.Background(), initialOptions...)
	ctx, cancel := chromedp.NewContext(startCtx)
	defer cancel()
	if err := chromedp.Run(ctx,
		chromedp.Navigate(`http://localhost:8000/root.html`),
	); err != nil {
		fmt.Println(err)
	}

	var iframes []*cdp.Node
	if err := chromedp.Run(ctx, chromedp.Nodes(`iframe`, &iframes, chromedp.ByQuery)); err != nil {
		fmt.Println(err)
	}

	if err := chromedp.Run(ctx, chromedp.Nodes(`iframe`, &iframes, chromedp.ByQuery, chromedp.FromNode(iframes[0]))); err != nil {
		fmt.Println(err)
	}
	var text string
	if err := chromedp.Run(ctx,
		chromedp.Text("#cucamanga2", &text, chromedp.ByQuery, chromedp.FromNode(iframes[0])),
		chromedp.Sleep(5*time.Second),
	); err != nil {
		fmt.Println(err)
	}
	fmt.Println(text)

}
