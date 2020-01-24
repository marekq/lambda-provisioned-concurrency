package main

import (
	"context"
	"errors"
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

func sqshandler(ctx context.Context, sqsEvent events.SQSEvent) error {
	log.Printf("Starting handler")
	if len(sqsEvent.Records) == 0 {
		return errors.New("No SQS message passed to function")
	}

	for _, msg := range sqsEvent.Records {
		log.Printf("Got SQS message %q with body %q\n", msg.MessageId, msg.Body)
	}

	return nil
}

func main() {
	lambda.Start(sqshandler)
}
