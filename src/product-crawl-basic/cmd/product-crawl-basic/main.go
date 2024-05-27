package main

import (
	"context"
	"encoding/json"
	"fmt"
	// "log"
	"strings"
	"time"

	//"github.com/chromedp/cdproto/dom"
	// "github.com/chromedp/cdproto/cdp"
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
	fmt.Printf("%#v", targets)
	fmt.Printf("%#v", uriPart)
	jsonStr, _ := json.MarshalIndent(targets, "", "  ")
	fmt.Println(string(jsonStr))
	for _, t := range targets {
		if t.Type == "iframe" && strings.Contains(t.URL, uriPart) {
			fmt.Println(t.Title, "|", t.Type, "|", t.URL, "|", t.TargetID)
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
	)
	// create context
	startCtx, _ := chromedp.NewExecAllocator(context.Background(), initialOptions...)
	ctx, cancel := chromedp.NewContext(startCtx)
	defer cancel()

	var header string
	chromedp.Run(ctx,
		chromedp.Navigate("http://localhost:8000/root.html"),
		chromedp.Sleep(2*time.Second),
		chromedp.Text(`h1#cucamanga`, &header, chromedp.ByQuery),
		// chromedp.ActionFunc(func(ctx context.Context) error {
		// 	fmt.Println(header)
		// 	return nil
		// }),
	)
	fmt.Println(header)
	// var iframes []*cdp.Node
	// if err := chromedp.Run(ctx, chromedp.Nodes(`iframe`, &iframes, chromedp.ByQuery)); err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%#v", iframes)
	// if err := chromedp.Run(ctx,
	// 	chromedp.Text(`h1#cucamanga`, &header, chromedp.ByID, chromedp.FromNode(iframes[0])),
	// 	chromedp.ActionFunc(func(ctx context.Context) error {
	// 		fmt.Println(header)
	// 		return nil
	// 	}),
	// ); err != nil {
	// 	log.Fatal(err)
	// }
	//
	// var html string
	// var iframeNode []*cdp.Node
	// err := chromedp.Run(ctx,
	// 	// visit the target page
	// 	chromedp.Tasks{
	// 		navigateAndWaitForLoad("http://localhost:8000/root.html", "networkIdle"),
	// 	},
	// 	chromedp.WaitReady("iframe", chromedp.ByQuery),
	// 	chromedp.Nodes("iframe", &iframeNode, chromedp.ByQuery),
	// 	// get the outer HTML of the page
	// 	// chromedp.ActionFunc(func(ctx context.Context) error {
	// 	// 	node, err := dom.GetDocument().Do(ctx)
	// 	// 	if err != nil {
	// 	// 		return err
	// 	// 	}
	// 	// 	html, err = dom.GetOuterHTML().WithNodeID(node.NodeID).Do(ctx)
	// 	// 	return err
	// 	// }),
	// )
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(html)
}
