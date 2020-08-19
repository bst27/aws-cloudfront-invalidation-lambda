package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"os"
)

func main() {
	// Allow cli invocation: ./app cli mySecretToken myDistributionId myInvalidationPath
	if len(os.Args) > 2 && os.Args[1] == "cli" {
		if len(os.Args) == 5 {
			rawBody, _ := json.Marshal(Body{
				SecretToken:      os.Args[2],
				DistributionId:   os.Args[3],
				InvalidationPath: os.Args[4],
			})

			resp, err := HandleRequest(context.Background(), Request{
				Body: string(rawBody),
			})

			fmt.Println(resp, err)
		}
	} else {
		lambda.Start(HandleRequest)
	}
}

type Request struct {
	Body string `json:"body"`
}

type Body struct {
	SecretToken      string
	DistributionId   string
	InvalidationPath string
}

type Response struct {
	Message string
	success bool
}

func HandleRequest(ctx context.Context, req Request) (Response, error) {
	body := &Body{}
	err := json.Unmarshal([]byte(req.Body), body)
	if err != nil {
		return Response{
			Message: "Invalid request",
			success: false,
		}, nil
	}

	// Deny access with empty token to enforce having a secret token defined
	if body.SecretToken == "" || body.SecretToken != os.Getenv("SECRET_TOKEN") {
		return Response{
			Message: "Forbidden",
			success: false,
		}, nil
	}

	if body.DistributionId == "" {
		return Response{
			Message: "Distribution ID is missing",
			success: false,
		}, nil
	}

	if body.InvalidationPath == "" {
		return Response{
			Message: "Invalidation path is missing",
			success: false,
		}, nil
	}

	err = invalidate(body.DistributionId, body.InvalidationPath)
	if err != nil {
		log.Println(err)

		return Response{
			Message: "Error",
			success: false,
		}, nil
	}

	return Response{
		Message: "OK",
		success: true,
	}, nil
}
