package dynamoutils

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

const region = "us-east-1"
const endpoint = "http://dynamodb:8000"
const table = "test-table"

func newClient() *dynamodb.DynamoDB {
	return NewDynamoClient(region, endpoint)
}

func TestCreateDynamoTableTest(t *testing.T) {
	t.Parallel()

	client := newClient()

	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{{
			AttributeName: aws.String("key"),
			AttributeType: aws.String("S"),
		}},
		KeySchema: []*dynamodb.KeySchemaElement{{
			AttributeName: aws.String("key"),
			KeyType:       aws.String("HASH"),
		}},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(5),
			WriteCapacityUnits: aws.Int64(5),
		},
		TableName: aws.String(table),
	}

	err := CreateDynamoTableIfNotExists(client, input)
	if err != nil {
		t.Fatalf(fmt.Sprintf("Error creating initial table %v", err))
	}

	err = CreateDynamoTable(client, input)
	if err == nil {
		t.Fatalf(fmt.Sprintf("should have had an error re-creating initial table %v", err))
	}

	err = CreateDynamoTableIfNotExists(client, input)
	if err != nil {
		t.Fatalf(fmt.Sprintf("should have had ingored error re-creating initial table %v", err))
	}
}
