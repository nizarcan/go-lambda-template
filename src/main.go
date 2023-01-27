package main

// write a lambda function that takes a custom struct as input and returns a custom struct as output

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	firstName string `json:"firstName"`
	lastName  string `json:"lastName"`
}

type MyResponse struct {
	fullName string `json:"fullName"`
}

func HandleRequest(ctx context.Context, name MyEvent) (MyResponse, error) {
	fullName := fmt.Sprintf("%s %s", name.firstName, name.lastName)
	return MyResponse{fullName: fullName}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
