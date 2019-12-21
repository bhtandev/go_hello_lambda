package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/tidwall/gjson"
)


func main() {
	lambda.Start(Handler)
}


func Handler(ctx context.Context, evt json.RawMessage) (events.APIGatewayProxyResponse, error) {
	param := gjson.Get(string(evt), "param1").String()
	token := gjson.Get(string(evt), "token")

	fmt.Println("Hello!")

	returnVal := events.APIGatewayProxyResponse{
		Body:       "Hello Lambda with" + param,
		StatusCode: 200,
	}

	if token.String() != os.Getenv("SECRET_KEY_TO_CALL_THIS_LAMBDA") {
		returnVal.Body = "Error - wrong token"
		returnVal.StatusCode = 401
	}

	return returnVal, nil
}
