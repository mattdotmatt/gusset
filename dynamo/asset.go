package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// Asset to save
type Asset struct {
	AssetId   string
	FileName  string
}

func (s storage) SaveAsset(asset Asset) error {

	av, err := dynamodbattribute.MarshalMap(asset)
	if err != nil {
		log.Print(err.Error())
		return err
	}

	tableName := "AdminHackDayAssetTable"

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = s.PutItem(input)

	if err != nil {
		log.Print(err.Error())
		return err
	}

	return nil
}