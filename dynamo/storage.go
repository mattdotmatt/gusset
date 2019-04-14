package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Dynamodb wrapper
type storage struct {
	*dynamodb.DynamoDB
}

func createSession() *session.Session {
	config := &aws.Config{
		Region:   aws.String("eu-west-1"),
	}

	return session.Must(session.NewSession(config))
}