package main

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type APIGatewayEvent struct {
	Name string `json:"name"`
}

type HttpHelper struct {
	Success bool `json:"success"`
}

func HandleRequest(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	body := HttpHelper{Success: true}
	response, err := json.Marshal(body)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500}, err
	}
	return events.APIGatewayProxyResponse{StatusCode: 200, Body: string(response)}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
