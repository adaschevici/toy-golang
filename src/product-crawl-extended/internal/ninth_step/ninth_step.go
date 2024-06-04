package ninth_step

import (
	"context"
	"fmt"
	"log"
	// "time"

	// "github.com/chromedp/cdproto/dom"
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
	var injectedScript = `(function() {
		let iframe = document.querySelector('div.challenge-stage iframe');
		console.log('iframe: ', iframe);
		return iframe;
	})();`
	var html string
	if err := chromedp.Run(ctx,
		chromedp.Navigate("https://www.g2.com/products/jira/reviews/"),
		chromedp.Evaluate(injectedScript, &html),
		// chromedp.Sleep(60*time.Second),
		// we have an iframe, so we need to switch to it before we can interact with it
		// chromedp.WaitVisible(`label.cb-lb > input`),
		// chromedp.Click(`label.cb-lb > input`),
		// chromedp.ActionFunc(func(ctx context.Context) error {
		// 	rootNode, err := dom.GetDocument().Do(ctx)
		// 	if err != nil {
		// 		return err
		// 	}
		// 	html, err = dom.GetOuterHTML().WithNodeID(rootNode.NodeID).Do(ctx)
		// 	return err
		// }),
	); err != nil {
		log.Fatal("Error while trying to grab current ip.", err)
		return
	}

	fmt.Println("HTML: ", html)
}

// snippet to allow Chrome to shutdown properly
// ctx, cancel := context.WithTimeout(context.Background(), 300 * time.Second)
// defer cancel()
//
// // create chrome instance
// var c *chromedp.CDP
// c, err := chromedp.New(ctx, chromeOption, chromedp.WithLog(log.Printf))
//
// // make sure Chrome is closed when exit
// defer func() {
// 	log.Printf("Shutdown chrome")
// 	// shutdown chrome
// 	err = c.Shutdown(ctx)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	// wait for chrome to finish
// 	// Wait() will hang on Windows and Linux with Chrome headless mode
// 	// we'll need to exit the program when this happens
// 	ch := make(chan error)
// 	go func() {
// 		c.Wait()
// 		ch <- nil
// 	}()
//
// 	select {
// 	case err = <-ch:
// 		log.Println("chrome closed")
// 	case <-time.After(10 * time.Second):
// 		log.Println("chrome didn't shutdown within 10s")
// 	}
// }()
//
// err = taskFunc(ctx, c)
// if err != nil {
// 	// don't use Fatal because we need to deter functions to run
// 	log.Printf("%v", err)
// 	return
// }
//
