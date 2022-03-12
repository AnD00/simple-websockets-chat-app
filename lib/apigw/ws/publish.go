package ws

import (
	"context"
	"fmt"
	"simple-websockets-chat-app/lib/dynamodb"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi"
	"github.com/aws/aws-sdk-go/aws/awserr"
)

func Publish(cli *apigatewaymanagementapi.Client, ctx context.Context, id string, data []byte) error {
	_, err := cli.PostToConnectionRequest(&apigatewaymanagementapi.PostToConnectionInput{
		Data:         data,
		ConnectionId: aws.String(id),
	}).Send(ctx)

	handleError(err, id)

	return err
}

func handleError(err error, id string) error {
	if err == nil {
		return err
	}

	if aerr, ok := err.(awserr.Error); ok {
		switch aerr.Code() {
		case aws.ErrCodeSerialization:
			fmt.Println("delete stale connection details from cache")
			return deleteConnectionId(id)
		case apigatewaymanagementapi.ErrCodeGoneException:
			fmt.Println("delete stale connection details from cache")
			return deleteConnectionId(id)
		default:
			return err
		}
	}

	return err
}

func deleteConnectionId(id string) error {
	err := dynamodb.DeleteConnection(id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("websocket connection deleted from cache")

	return err
}
