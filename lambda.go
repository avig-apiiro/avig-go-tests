package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"os"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	//required query
	address := request.QueryStringParameters["address"]

	//Replace with api key
	key := os.Getenv("GOOGLE_API_KEY")

	geoParams := map[string]string{
		"address": address,
		"key":     key,
	}

	//obtains find place response to be processed

	jsonString, _ := json.Marshal(geoParams)

	//Returning response with AWS Lambda Proxy Response
	return events.APIGatewayProxyResponse{Body: string(jsonString), StatusCode: 200}, nil
}

func lambdamain() {
	lambda.Start(Handler)
}
