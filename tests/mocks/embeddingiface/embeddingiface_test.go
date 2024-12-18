package embeddingiface_test

import (
	"articles/tests/mocks/embeddingiface"
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type dynamoDBClient struct {
	dynamodbiface.DynamoDBAPI

	expect    *dynamodb.GetItemOutput
	expectErr error
}

func (dc *dynamoDBClient) GetItem(*dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	return dc.expect, dc.expectErr
}

func TestGetItem(t *testing.T) {
	tests := []struct {
		name        string
		description string
		dynamoCli   *dynamoDBClient
		expect      *dynamodb.GetItemOutput
		expectErr   error
	}{
		{
			name:        "get item success",
			description: "should return a dynamodb.GetItemOutput without errors",
			dynamoCli: &dynamoDBClient{
				expect: &dynamodb.GetItemOutput{
					Item: make(map[string]*dynamodb.AttributeValue),
				},
				expectErr: nil,
			},
			expect: &dynamodb.GetItemOutput{
				Item: make(map[string]*dynamodb.AttributeValue),
			},
			expectErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Log(test.description)

			item, err := embeddingiface.GetItem(test.dynamoCli, "mykey")

			if !reflect.DeepEqual(test.expect, item) {
				t.Errorf("expected %v, got %v", test.expect, item)
			}

			if err != test.expectErr {
				t.Errorf("expected error %v, got %v", test.expectErr, err)
			}
		})
	}
}
