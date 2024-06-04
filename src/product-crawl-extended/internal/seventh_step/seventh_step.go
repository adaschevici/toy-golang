package seventh_step

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/chromedp/cdproto/browser"
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
	done := make(chan string, 1)
	chromedp.ListenTarget(ctx, func(ev interface{}) {
		if evt, ok := ev.(*browser.EventDownloadProgress); ok {
			completed := "(unknown)"
			if evt.TotalBytes != 0 {
				completed = fmt.Sprintf("%0.2f%%", evt.ReceivedBytes/evt.TotalBytes*100.0)
			}
			log.Printf("state: %s, completed: %s\n", evt.State.String(), completed)
			if evt.State == browser.DownloadProgressStateCompleted {
				done <- evt.GUID
				close(done)
			}
		}
	})

	workingDirPath, err := os.Getwd()
	if err != nil {
		log.Fatal("Error while trying to get cwd.", err)
	}
	if err := chromedp.Run(ctx,
		chromedp.Navigate("http://localhost:5000/download"),
		browser.
			SetDownloadBehavior(browser.SetDownloadBehaviorBehaviorAllowAndName).
			WithDownloadPath(workingDirPath).
			WithEventsEnabled(true),
		chromedp.Click(`button#download`),
	); err != nil && !strings.Contains(err.Error(), "net::ERR_ABORTED") {
		log.Fatal("Error while trying to download a file.", err)
	}
	<-done
}
