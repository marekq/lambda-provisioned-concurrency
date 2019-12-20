package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func httphandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Body size = %d. \n", len(request.Body))
	log.Println("Headers:")

	for key, value := range request.Headers {
		log.Printf("  %s: %s\n", key, value)
	}

	return events.APIGatewayProxyResponse{Body: "GET", StatusCode: 200}, nil
}

func main() {
	lambda.Start(httphandler)
}
