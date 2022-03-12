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
	fmt.Println("websocket connect")

	err := dynamodb.PutConnection(req.RequestContext.ConnectionID, req.QueryStringParameters["room"])
	if err != nil {
		fmt.Println(err)
		return apigw.InternalServerErrorResponse(), err
	}

	fmt.Println("websocket connection cached")

	return apigw.OkResponse(), nil
}
