package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

type APIGatewayEvent struct {
	Name string `json:"name"`
}

type HttpHelper struct {
	Success bool `json:"success"`
}

type APIGatewayEventResponse struct {
	StatusCode int        `json:"statusCode"`
	Body       HttpHelper `json:"body"`
}

func HandleRequest(ctx context.Context, name APIGatewayEvent) (APIGatewayEventResponse, error) {
	return APIGatewayEventResponse{StatusCode: 200, Body: HttpHelper{Success: true}}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
