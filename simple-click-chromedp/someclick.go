package main

import (
	"context"
	"log"
	"os"

	"github.com/chromedp/chromedp"
)

func main() {
	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// run task list
	var res []string
	// var jsSpoofer = `
	//
	// 	// const getParameter = WebGLRenderingContext.prototype.getParameter;
	// 	// WebGLRenderingContext.prototype.getParameter = function(parameter) {
	// 	//     if (parameter === 37445) {
	// 	// 	return 'Intel Inc.';
	// 	//     }
	// 	//     if (parameter === 37446) {
	// 	// 	return 'Intel Iris OpenGL Engine';
	// 	//     }
	// 	//     return getParameter(parameter);
	// 	// };
	// 	// // Spoof other common bot-detection properties
	//                // Object.defineProperty(navigator, 'webdriver', {
	//                //     get: () => false,
	//                // });
	//
	//                // Object.defineProperty(window, 'chrome', {
	//                //     get: () => true,
	//                // });
	//                // //
	//                // window.navigator.permissions.query = (parameters) => (
	//                //     parameters.name === 'notifications' ?
	//                //     Promise.resolve({ state: Notification.permission }) :
	//                //     Promise.resolve({ state: 'denied' })
	//                // );
	//                //
	//                // // Spoof languages
	//                Object.defineProperty(navigator, 'languages', {
	//                    get: () => ['en-UK', 'en'],
	//                });
	//                //
	//                // // Spoof plugins
	//                // Object.defineProperty(navigator, 'plugins', {
	//                //     get: () => [1, 2, 3, 4, 5],
	//                // });
	//     `
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
		chromedp.Sleep(5),
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
