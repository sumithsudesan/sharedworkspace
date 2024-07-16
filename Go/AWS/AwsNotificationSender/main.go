package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

//go get github.com/aws/aws-lambda-go/lambda

func handleRequest(ctx context.Context, snsEvent events.SNSEvent) (events.APIGatewayProxyResponse, error) {
	var request RecipientInfo
	fmt.Printf("[INFO] Message = %s \n", snsEvent.Records[0].SNS.Message)

	decoder := json.NewDecoder(strings.NewReader(snsEvent.Records[0].SNS.Message))
	err := decoder.Decode(&request)
	if err != nil {
		fmt.Printf("[ERROR] Failed to decode request. message : %s\n", snsEvent.Records[0].SNS.Message)
		return events.APIGatewayProxyResponse{}, err
	}
	fmt.Printf("[INFO] Sending Notification, request :%v\n", request)

	chanelFactory := ChannelFactory{}
	// Send notification
	err = chanelFactory.Send(request, InitConfig())
	if err != nil {
		fmt.Printf("[ERROR] Failed to send notiication, err=%v\n", err)
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{Body: string(""), StatusCode: 200}, nil
}

func main() {
	lambda.Start(handleRequest)
}
