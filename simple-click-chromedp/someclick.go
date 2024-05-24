package main

import (
	"context"
	"log"

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
	var jsDef = `
		Object.defineProperty(window, 'languages', {
			get: function() {
				return true;	
			}
		});

		`

	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://www.google.com/`),
		chromedp.Evaluate(jsDef, nil),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("window object keys: %v", res)
}
