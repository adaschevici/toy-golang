package main

import (
	"context"
	"flag"
	"log"
	"os"
	"time"

	"github.com/chromedp/chromedp"
)

// var jsSpoofer = `
//
//		// const getParameter = WebGLRenderingContext.prototype.getParameter;
//		// WebGLRenderingContext.prototype.getParameter = function(parameter) {
//		//     if (parameter === 37445) {
//		// 	return 'Intel Inc.';
//		//     }
//		//     if (parameter === 37446) {
//		// 	return 'Intel Iris OpenGL Engine';
//		//     }
//		//     return getParameter(parameter);
//		// };
//		// // Spoof other common bot-detection properties
//	               // Object.defineProperty(navigator, 'webdriver', {
//	               //     get: () => false,
//	               // });
//
//	               // Object.defineProperty(window, 'chrome', {
//	               //     get: () => true,
//	               // });
//	               // //
//	               // window.navigator.permissions.query = (parameters) => (
//	               //     parameters.name === 'notifications' ?
//	               //     Promise.resolve({ state: Notification.permission }) :
//	               //     Promise.resolve({ state: 'denied' })
//	               // );
//	               //
//	               // // Spoof languages
//	               Object.defineProperty(navigator, 'languages', {
//	                   get: () => ['en-UK', 'en'],
//	               });
//	               //
//	               // // Spoof plugins
//	               // Object.defineProperty(navigator, 'plugins', {
//	               //     get: () => [1, 2, 3, 4, 5],
//	               // });
//	    `
func main() {
	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// run task list
	var res []string
	verbose := flag.Bool("v", false, "verbose")
	timeout := flag.Duration("timeout", 2*time.Minute, "timeout")
	scale := flag.Float64("scale", 1.5, "scale")
	padding := flag.Int("padding", 0, "padding")
	out := flag.String("out", "", "out")
	flag.Parse()
	var screenshot []byte
	var jsDef = `
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

		`

	err := chromedp.Run(ctx,
		// chromedp.Navigate(`https://bot.sannysoft.com/`),
		chromedp.Navigate(`https://nowsecure.nl/`),
		chromedp.Evaluate(jsDef, nil),
		chromedp.Evaluate(`Object.keys(window)`, &res),
		chromedp.CaptureScreenshot(&screenshot),
	)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = os.WriteFile("screenshot.png", screenshot, 0644)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Printf("window object keys: %v", res)
}

func run(ctx context.Context, verbose bool, timeout time.Duration, scale float64, padding int, out string) error {
	var opts []chromedp.ContextOption
	// create context
	if verbose {
		opts = append(opts, chromedp.WithDebugf(log.Printf))
	}
	ctx, cancel := chromedp.NewContext(ctx, opts...)
	defer cancel()

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, timeout)
	defer cancel()

	start := time.Now()

	// capture screenshot
	var buf []byte
	var err error
	if err = chromedp.Run(ctx,
		chromedp.Navigate(`https://fast.com`),
		chromedp.WaitVisible(`#speed-value.succeeded`),
		chromedp.Click(`#show-more-details-link`),
		chromedp.WaitVisible(`#upload-value.succeeded`),
		chromedp.ScreenshotScale(`.speed-controls-container`, scale, &buf),
	); err != nil {
		return err
	}
	return err
}
