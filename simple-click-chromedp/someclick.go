package main

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/kb"
)

func main() {

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.DisableGPU,
		chromedp.Flag("enable-automation", false),
		chromedp.Flag("disable-gpu", false),
		chromedp.Flag("enable-webgl", true),
		chromedp.Flag("disable-software-rasterizer", false),
		chromedp.Flag("disable-blink-features", "AutomationControlled"),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	// chromedp.Evaluate
	defer cancel()

	// also set up a custom logger
	taskCtx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancel()

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
