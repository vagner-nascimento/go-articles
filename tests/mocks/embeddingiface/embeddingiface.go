package embeddingiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

func GetItem(svc dynamodbiface.DynamoDBAPI, key string) (*dynamodb.GetItemOutput, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String("myTable"),
		Key: map[string]*dynamodb.AttributeValue{
			"myKey": {
				N: aws.String(key),
			},
		},
	}

	return svc.GetItem(input)
}
