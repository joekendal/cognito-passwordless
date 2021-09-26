package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(event events.CognitoEventUserPoolsVerifyAuthChallenge) (events.CognitoEventUserPoolsVerifyAuthChallenge, error) {
	expectedAnswer := event.Request.PrivateChallengeParameters["secretCode"]

	if event.Request.ChallengeAnswer == expectedAnswer {
		event.Response.AnswerCorrect = true
	} else {
		event.Response.AnswerCorrect = false
	}

	return event, nil
}

func main() {
	lambda.Start(Handler)
}
