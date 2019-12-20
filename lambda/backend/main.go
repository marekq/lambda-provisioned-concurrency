package main

import (
	"context"
	"errors"
	"log"

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

func httphandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Body size = %d. \n", len(request.Body))
	log.Println("Headers:")

	for key, value := range request.Headers {
		log.Printf("  %s: %s\n", key, value)
	}

	return events.APIGatewayProxyResponse{Body: "GET", StatusCode: 200}, nil
}

func main1() {
	//mode := os.Getenv("HTTPSQS")
	mode := "HTTP" 

	if mode == "SQS" {
		lambda.Start(sqshandler)

	} else {
		lambda.Start(httphandler)
	}
}

func main() {

	// TODO - only supporting http for now
	lambda.Start(httphandler)
}