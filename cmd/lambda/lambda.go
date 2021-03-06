package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type APIGatewayEvent struct {
	Name string `json:"name"`
}

type Response struct {
	Success bool `json:"success"`
}

func HandleRequest(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Headers: ", event.Headers)
	body := Response{Success: true}
	response, err := json.Marshal(body)
	if err != nil {
		errorResponse, _ := json.Marshal(Response{Success: false})
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: string(errorResponse)}, nil
	}
	fmt.Println("APIGatewayProxyResponse: ", events.APIGatewayProxyResponse{StatusCode: 200, Body: string(response)})
	return events.APIGatewayProxyResponse{StatusCode: 200, Body: string(response)}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
