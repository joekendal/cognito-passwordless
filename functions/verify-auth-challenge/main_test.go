package main_test

import (
	main "cognito-passwordless/verify-auth-challenge"
	"log"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestHandler(t *testing.T) {
	event, err := main.Handler(events.CognitoEventUserPoolsVerifyAuthChallenge{
		Request: events.CognitoEventUserPoolsVerifyAuthChallengeRequest{
			PrivateChallengeParameters: map[string]string{
				"secretCode": "1234",
			},
			ChallengeAnswer: "1234",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	if !event.Response.AnswerCorrect {
		log.Fatal("Expected Response.AnswerCorrect to be true")
	}
}
