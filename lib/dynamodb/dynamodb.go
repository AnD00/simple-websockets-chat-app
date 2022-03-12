package dynamodb

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"github.com/pkg/errors"
)

func getTable(db *dynamo.DB, tableName string) dynamo.Table {
	return db.Table(tableName)
}

func connect() (*dynamo.DB, error) {
	config := aws.Config{
		Endpoint: aws.String(os.Getenv("DYNAMO_ENDPOINT")),
	}

	dynamoSession, err := session.NewSession(&config)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return dynamo.New(dynamoSession), nil
}
