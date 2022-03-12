package main

import (
	"context"
	"fmt"

	"simple-websockets-chat-app/lib/apigw"
	"simple-websockets-chat-app/lib/dynamodb"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

func handler(_ context.Context, req *events.APIGatewayWebsocketProxyRequest) (apigw.Response, error) {
	fmt.Println("websocket disconnect")

	err := dynamodb.DeleteConnection(req.RequestContext.ConnectionID)
	if err != nil {
		fmt.Println(err)
		return apigw.InternalServerErrorResponse(), err
	}

	fmt.Println("websocket connection deleted from cache")

	return apigw.OkResponse(), nil
}
