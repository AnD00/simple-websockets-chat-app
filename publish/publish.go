package main

import (
	"context"
	"fmt"
	"time"

	"simple-websockets-chat-app/lib/apigw"
	"simple-websockets-chat-app/lib/apigw/ws"
	"simple-websockets-chat-app/lib/dynamodb"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, req *events.APIGatewayWebsocketProxyRequest) (apigw.Response, error) {
	fmt.Println("websocket publish")

	apiClient := apigw.NewAPIGatewayManagementClient(req.RequestContext.DomainName, req.RequestContext.Stage)

	input, err := new(ws.InputEnvelop).Decode([]byte(req.Body))
	if err != nil {
		fmt.Println(err)
		return apigw.BadRequestResponse(), err
	}

	output := &ws.OutputEnvelop{
		Data:     input.Data,
		Received: time.Now().Unix(),
	}

	data, err := output.Encode()
	if err != nil {
		fmt.Println(err)
		return apigw.InternalServerErrorResponse(), err
	}

	conns, err := dynamodb.GetAllConnections(input.Room)
	if err != nil {
		fmt.Println(err)
		return apigw.InternalServerErrorResponse(), err
	}

	fmt.Println("websocket connections read from cache")

	sender := req.RequestContext.ConnectionID
	for _, conn := range conns {
		id := conn.ConnectionID
		if id == sender {
			continue
		}
		ws.Publish(apiClient, ctx, id, data)
	}

	return apigw.OkResponse(), nil
}
