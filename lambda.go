package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

type APIGatewayEvent struct {
	Name string `json:"name"`
}

type APIGatewayEventResponse struct {
	StatusCode int    `json:"statusCode"`
	Body       string `json:"body"`
}

func HandleRequest(ctx context.Context, name APIGatewayEvent) (APIGatewayEventResponse, error) {
	return APIGatewayEventResponse{StatusCode: 200, Body: name.Name}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
