package main

import (
	"context"
	"errors"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-xray-sdk-go/xray"
	"github.com/aws/aws-xray-sdk-go/xraylog"
)

func init() {
	xray.SetLogger(xraylog.NullLogger)
}

func sqshandler(ctx context.Context, sqsEvent events.SQSEvent) error {
	xray.Configure(xray.Config{LogLevel: "trace"})
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
