package dynamoutils

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// NewDynamoClient returns a new client used to communicate with DynamoDB
func NewDynamoClient(region, endpoint string) *dynamodb.DynamoDB {
	awscfg := aws.NewConfig().
		WithRegion(region).
		WithEndpoint(endpoint)
	return dynamodb.New(session.Must(session.NewSession(awscfg)), awscfg)
}

// CreateDynamoTable will create a new DynamoDB table. This function will error if the table
// has already been created (see CreateDynamoTableIfNotExists:)
func CreateDynamoTable(client *dynamodb.DynamoDB, schema *dynamodb.CreateTableInput) error {
	return createTable(client, schema, false)
}

// CreateDynamoTableIfNotExists will create a DynamoDB table if it does not already exist
func CreateDynamoTableIfNotExists(client *dynamodb.DynamoDB, schema *dynamodb.CreateTableInput) error {
	return createTable(client, schema, true)
}

func createTable(client *dynamodb.DynamoDB, schema *dynamodb.CreateTableInput, ignoreExistingTableError bool) error {
	_, err := client.CreateTable(schema)

	if err != nil {
		aerr, ok := err.(awserr.Error)
		if ok && ignoreExistingTableError && aerr.Code() == dynamodb.ErrCodeResourceInUseException {
			err = nil
		}
	}

	return err
}
