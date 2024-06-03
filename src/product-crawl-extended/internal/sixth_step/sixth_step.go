package sixth_step

import (
	"context"
	"fmt"
	"log"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
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
	done := make(chan bool)
	var requestID network.RequestID
	var urlstr string
	if err := chromedp.Run(ctx,
		chromedp.Navigate("http://localhost:5000/download"),
		chromedp.Evaluate(`document.querySelector('button#download > a').getAttribute('href')`, &urlstr),
	); err != nil {
		log.Fatal("Error while trying to grab product items.", err)
	}
	log.Printf("URL: %v", urlstr)
	chromedp.ListenTarget(ctx, func(ev interface{}) {
		switch ev := ev.(type) {
		case *network.EventRequestWillBeSent:
			log.Printf("EventRequestWillBeSent: %v: %v", ev.RequestID, ev.Request.URL)
			if ev.Request.URL == fmt.Sprintf("http://localhost:5000%s", urlstr) {
				log.Printf("Initializing requestID: %v: %v", ev.RequestID, ev.Request.URL)
				requestID = ev.RequestID
			}
		case *network.EventLoadingFinished:
			log.Printf("EventLoadingFinished: %v", ev.RequestID)
			if ev.RequestID == requestID {
				log.Printf("EventLoadingFinished: %v matched %v", requestID, ev.RequestID)
				done <- true
			}
		}
	})
	if err := chromedp.Run(ctx,
		chromedp.Click(`button#download`),
	); err != nil {
		log.Fatal("Error while trying to grab product items.", err)
	}
	<-done
	fmt.Println("This is the sixth step.")
}
