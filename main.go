package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/chromedp/chromedp"
)

func Handler(_ context.Context, _ any) error {
	opts := []chromedp.ExecAllocatorOption{
		// chromedp.ExecPath(), // TODO: ...
		chromedp.NoSandbox,
		chromedp.Flag("disable-setuid-sandbox", true),
		chromedp.Flag("disable-dev-shm-usage", true),
		chromedp.Flag("single-process", true),
		chromedp.Flag("no-zygote", true),
		chromedp.Headless,
	}

	cctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(cctx /* chromedp.WithDebugf(log.Printf) */)
	defer cancel()

	url := "https://x.com/tweet/status/1934128117171601448"
	err := chromedp.Run(ctx, chromedp.Navigate(url))
	if err != nil {
		log.Printf("error navigating to website: %v", err)
		return fmt.Errorf("error navigating to website '%s': %v", url, err)
	}

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("error making http request: %v", err)
		return err
	}
	defer resp.Body.Close()
	fmt.Printf("\nStatusCode: %d\n", resp.StatusCode)

	return nil
}

func main() {
	if _, exists := os.LookupEnv("AWS_LAMBDA_RUNTIME_API"); exists {
		lambda.Start(Handler)
	} else {
		err := Handler(context.Background(), nil)

		time.Sleep(time.Hour * 20) // TODO: remove

		if err != nil {
			log.Fatal(err)
		}
	}
}
