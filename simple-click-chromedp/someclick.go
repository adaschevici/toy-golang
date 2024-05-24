package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/kenshaw/rasterm"
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

	// run task list
	verbose := flag.Bool("v", false, "verbose")
	timeout := flag.Duration("timeout", 2*time.Minute, "timeout")
	scale := flag.Float64("scale", 1.5, "scale")
	padding := flag.Int("padding", 0, "padding")
	headless := flag.Bool("headless", false, "headless")
	out := flag.String("out", "", "out")
	flag.Parse()
	if err := run(context.Background(), *headless, *verbose, *timeout, *scale, *padding, *out); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, headless bool, verbose bool, timeout time.Duration, scale float64, padding int, out string) error {
	var opts []chromedp.ContextOption
	var initialOptions = append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("disable-gpu", false),
		chromedp.Flag("headless", false),
		chromedp.Flag("no-sandbox", true),
	)
	// create context
	startCtx, _ := chromedp.NewExecAllocator(ctx, initialOptions...)
	if verbose {
		opts = append(opts, chromedp.WithDebugf(log.Printf))
	}
	ctx, cancel := chromedp.NewContext(startCtx, opts...)
	defer cancel()

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, timeout)
	defer cancel()

	start := time.Now()

	// capture screenshot
	var buf []byte
	if err := chromedp.Run(ctx,
		chromedp.Navigate(`https://fast.com`),
		chromedp.WaitVisible(`#speed-value.succeeded`),
		chromedp.Click(`#show-more-details-link`),
		chromedp.WaitVisible(`#upload-value.succeeded`),
		chromedp.ScreenshotScale(`.speed-controls-container`, scale, &buf),
	); err != nil {
		return err
	}
	end := time.Now()

	// decode png
	img, err := png.Decode(bytes.NewReader(buf))
	if err != nil {
		return err
	}

	// pad image
	if padding != 0 {
		bounds := img.Bounds()
		w, h := bounds.Dx(), bounds.Dy()
		dst := image.NewRGBA(image.Rect(0, 0, w+2*padding, h+2*padding))
		for x := 0; x < w+2*padding; x++ {
			for y := 0; y < h+2*padding; y++ {
				dst.Set(x, y, color.White)
			}
		}
		draw.Draw(dst, dst.Bounds(), img, image.Pt(-padding, -padding), draw.Src)
		img = dst
	}

	// write to disk
	if out != "" {
		if err := os.WriteFile(out, buf, 0o644); err != nil {
			return err
		}
	}

	// output
	if err := rasterm.Encode(os.Stdout, img); err != nil {
		return err
	}

	// metrics
	_, err = fmt.Fprintf(os.Stdout, "time: %v\n", end.Sub(start))
	return err
}
