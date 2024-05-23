package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp"
	// "github.com/chromedp/chromedp/kb"
)

func main() {

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		// chromedp.Flag("headless", false),
		// chromedp.DisableGPU,
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

	// err := chromedp.Run(taskCtx,
	// 	chromedp.Navigate(`https://en.wikipedia.org/wiki/Main_Page`),
	// 	chromedp.WaitVisible(`.cdx-text-input__input`, chromedp.ByQuery),
	// 	chromedp.SendKeys(`.cdx-text-input__input`, "golang", chromedp.ByQuery),
	//
	// 	chromedp.SendKeys(`.cdx-text-input__input`, kb.Enter, chromedp.ByQuery),
	// )
	// time.Sleep(2 * time.Second)
	var screenshot []byte
	err := chromedp.Run(taskCtx,
		chromedp.Navigate(`https://bot.sannysoft.com/`),
		chromedp.ActionFunc(func(ctx context.Context) error {
			// Inject JavaScript to spoof WebGL properties
			js := `
		() => {
		    const getParameter = WebGLRenderingContext.prototype.getParameter;
		    WebGLRenderingContext.prototype.getParameter = function(parameter) {
			if (parameter === 37445) {
			    return 'Intel Inc.';
			}
			if (parameter === 37446) {
			    return 'Intel Iris OpenGL Engine';
			}
			return getParameter(parameter);
		    };
		}
		// Spoof other common bot-detection properties
                Object.defineProperty(navigator, 'webdriver', {
                    get: () => false,
                });

                Object.defineProperty(window, 'chrome', {
                    get: () => true,
                });

                window.navigator.permissions.query = (parameters) => (
                    parameters.name === 'notifications' ?
                    Promise.resolve({ state: Notification.permission }) :
                    Promise.resolve({ state: 'denied' })
                );

                // Spoof languages
                Object.defineProperty(navigator, 'languages', {
                    get: () => ['en-US', 'en'],
                });

                // Spoof plugins
                Object.defineProperty(navigator, 'plugins', {
                    get: () => [1, 2, 3, 4, 5],
                });
	    `
			_, _, err := runtime.Evaluate(js).Do(ctx)
			if err != nil {
				return err
			}
			return nil
		}),
		// chromedp.WaitVisible(`body`, chromedp.ByQuery), // Wait for the body to be visible
		chromedp.Sleep(2*time.Second), // Optional: Wait for any animations or dynamic content to load
		chromedp.CaptureScreenshot(&screenshot),
	)
	time.Sleep(2 * time.Second)
	// ensure that the browser process is started
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile("screenshot.png", screenshot, 0644)
	if err != nil {
		fmt.Println("Error saving screenshot:", err)
		return
	}

}
