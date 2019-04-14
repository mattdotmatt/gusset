package main

import (
	"log"

	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/google/uuid"
)

// Handle the arrival of an asset
func assetHandler(ctx context.Context, s3Event events.S3Event) error {
	log.Print("assetHandler fired")

	defer log.Print("assetHandler finished")

	var s = storage{dynamodb.New(createSession())}

	for _, record := range s3Event.Records {
		asset := Asset{
			AssetId: uuid.Must(uuid.NewRandom()).String(),
			FileName: record.S3.Object.Key,
		}
		err := s.SaveAsset(asset)

		if err != nil {
			return err
		}
	}

	return nil
}

func main(){
	lambda.Start(assetHandler)
}