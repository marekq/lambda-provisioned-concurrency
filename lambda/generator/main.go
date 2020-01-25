package main

import (
	"context"
	"log"
	"math/rand"
	"net/http"

	"os"
	"strconv"
	"sync"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

const (
	msgPerThread = 100
)

func handler(ctx context.Context) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// set counter for total sent messages
	allCount := 0

	// get the amount of messages per go routine, default 10
	msgCount, _ := strconv.Atoi(os.Getenv("MessageAmount"))

	// get whether HTTP or SQS messages should be sent by the generator
	mode := os.Getenv("HTTPSQS")

	// if SQS messages should be sent
	if mode == "SQS" {

		// get the SQS queue URL
		q := os.Getenv("SQSurl")

		// create a session with sqs
		svc1 := sqs.New(sess)

		// spawn a new go routine for every http request

		for a := 0; a < msgCount; a++ {

			// create a wait group to wait for go subroutines
			var wg sync.WaitGroup

			// spawn go routines depending on total count
			for b := 0; b < msgPerThread; b++ {

				// add one count to the workgroup
				wg.Add(1)
				allCount += 1

				// run the send message command in parallel
				go func() {

					defer wg.Done()
					ri := strconv.Itoa(rand.Intn(9999999) + rand.Intn(9999999)*rand.Intn(9999999))

					// send the message to the sqs queue
					_, err := svc1.SendMessage(&sqs.SendMessageInput{MessageBody: aws.String(ri), QueueUrl: aws.String(q)})

					if err != nil {
						log.Println(err)
					}
				}()

			}
			// wait for all routines to finish
			wg.Wait()

			//print the total count of messages
			log.Println("* sent " + strconv.Itoa(allCount) + " SQS messages.")
		}
	}

	// if HTTP requests should be sent
	if mode == "HTTP" {

		// get the API GW URL
		s3uri := os.Getenv("HTTPurl")

		// spawn a new go routine for every http request
		for a := 0; a < msgCount; a++ {

			// create a wait group to wait for go subroutines
			var wg sync.WaitGroup

			// spawn go routines depending on total count
			for b := 0; b < msgPerThread; b++ {

				// add one count to the workgroup
				wg.Add(1)
				allCount += 1

				// run the send message command in parallel
				go func() {
					defer wg.Done()
					ri := strconv.Itoa(rand.Intn(9999999) + rand.Intn(9999999)*rand.Intn(9999999))

					// send the http message to the api gateway
					_, err := http.Get(s3uri + ri)

					if err != nil {
						log.Println(err)
					}
				}()

			}
			// wait for all routines to finish
			wg.Wait()

			//print the total count of messages
			log.Println("* sent " + strconv.Itoa(allCount) + " HTTP messages to " + s3uri)
		}
	}
}

func main() {
	lambda.Start(handler)
}
