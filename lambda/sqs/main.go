package main

import (
	"context"
	"log"
	"errors"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

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
