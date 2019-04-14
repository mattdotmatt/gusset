package main

import (
	"log"
	"time"

	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Handle the arrival of an asset
func simpleHandler(ctx context.Context, s3Event events.S3Event) error {
	log.Print("simpleHandler fired")
	log.Print(time.Now().String())
	log.Print("simpleHandler finished")

	return nil
}

func main(){
	lambda.Start(simpleHandler)
}