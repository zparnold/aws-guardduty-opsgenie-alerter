package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"os"
)

type Response struct {
	Message string `json:"message"`
}

func Handler() (Response, error) {
	reqId, err := SendGenieAlert(os.Getenv("TEAM_NAME"), os.Getenv("TEAM_ID"))
	if err != nil {
		return Response{
			Message: "Something went wrong " + err.Error(),
		}, err
	} else {
		return Response{
			Message: "Alert Dispatched with ID: " + reqId,
		}, nil
	}

}

func main() {
	lambda.Start(Handler)
}
