package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	// "time"

	//"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/cdproto/target"
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
func getIframeContext(ctx context.Context, uriPart string) context.Context {
	targets, _ := chromedp.Targets(ctx)
	var tgt *target.Info
	for _, t := range targets {
		fmt.Println(t.Title, "|", t.Type, "|", t.URL, "|", t.TargetID)
		if t.Type == "iframe" && strings.Contains(t.URL, uriPart) {
			tgt = t
		}
	}
	if tgt != nil {
		ictx, _ := chromedp.NewContext(ctx, chromedp.WithTargetID(tgt.TargetID))
		return ictx
	}
	return nil
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

	ictx := getIframeContext(ctx, "/substring/in/iframe/uri/")
	selector := "a.LinkInIframe"
	script := fmt.Sprintf("document.querySelector(\"%s\").href;", selector)
	var b []byte
	_ = chromedp.Run(
		ictx, // <-- instead of ctx
		chromedp.WaitVisible(selector, chromedp.ByQuery),
		chromedp.Evaluate(script, &b),
	)
	fmt.Println("href in iframe:", string(b))

	var html string
	var iframeNode []*cdp.Node
	err := chromedp.Run(ctx,
		// visit the target page
		chromedp.Tasks{
			navigateAndWaitForLoad("http://localhost:8000/root.html", "networkIdle"),
		},
		chromedp.WaitReady("iframe", chromedp.ByQuery),
		chromedp.Nodes("iframe", &iframeNode, chromedp.ByQuery),
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
