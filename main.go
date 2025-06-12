package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/chromedp/chromedp"
)

func Handler(_ context.Context, _ any) error {
	opts := []chromedp.ExecAllocatorOption{
		chromedp.NoSandbox,
		chromedp.Flag("disable-setuid-sandbox", true),
		chromedp.Flag("disable-dev-shm-usage", true),
		chromedp.Flag("single-process", true),
		chromedp.Flag("no-zygote", true),
		chromedp.Headless,
	}

	cctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(cctx, chromedp.WithDebugf(log.Printf))
	defer cancel()

	url := "https://google.com"
	err := chromedp.Run(ctx, chromedp.Navigate(url))
	if err != nil {
		log.Println("error navigating to website")
		return fmt.Errorf("error navigating to website '%s': %v", url, err)
	}

	return nil
}

func main() {
	lambda.Start(Handler)
}
