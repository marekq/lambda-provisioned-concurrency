package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-xray-sdk-go/xray"
)

func init() {
	xray.Configure(xray.Config{
		DaemonAddr:     "127.0.0.1:2000",
		LogLevel:       "info",
		ServiceVersion: "1.2.3",
	})
}

func httphandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Body size = %d. \n", len(request.Body))
	log.Println("Headers:")

	for key, value := range request.Headers {
		log.Printf("  %s: %s\n", key, value)
	}

	return events.APIGatewayProxyResponse{
		Body: "<html><body>hello</body></html>",
		Headers: map[string]string{
			"Content-Type": "text/html",
		},
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(httphandler)
}
