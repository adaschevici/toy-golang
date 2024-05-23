package main

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/kb"
	// "github.com/go-rod/stealth"
)

func main() {

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.DisableGPU,
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	// chromedp.Evaluate
	defer cancel()

	// also set up a custom logger
	taskCtx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancel()
	// chromedp.Evaluate(stealth.JS, nil)

	err := chromedp.Run(taskCtx,
		chromedp.Navigate(`https://en.wikipedia.org/wiki/Main_Page`),
		chromedp.WaitVisible(`.cdx-text-input__input`, chromedp.ByQuery),
		chromedp.SendKeys(`.cdx-text-input__input`, "golang", chromedp.ByQuery),

		chromedp.SendKeys(`.cdx-text-input__input`, kb.Enter, chromedp.ByQuery),
	)
	time.Sleep(2 * time.Second)
	err = chromedp.Run(taskCtx,
		chromedp.Navigate(`https://bot.sannysoft.com/`),
	)
	time.Sleep(20 * time.Second)
	// ensure that the browser process is started
	if err != nil {
		log.Fatal(err)
	}

}
