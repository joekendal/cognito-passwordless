package main_test

import (
	main "cognito-passwordless/pre-signup"
	"errors"
	"log"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestHandler(t *testing.T) {
	event, err := main.Handler(events.CognitoEventUserPoolsPreSignup{})
	if err != nil {
		log.Fatal(err)
	}
	if !event.Response.AutoConfirmUser || !event.Response.AutoVerifyPhone {
		log.Fatal(errors.New("Expected response to AutoConfirmUser and AutoVerifyPhone"))
	}
}
