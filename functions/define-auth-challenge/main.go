package main

import (
	"errors"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(event events.CognitoEventUserPoolsDefineAuthChallenge) (events.CognitoEventUserPoolsDefineAuthChallenge, error) {

	// if user is not registered
	if event.Request.UserNotFound {
		event.Response.IssueTokens = false
		event.Response.FailAuthentication = true
		return event, errors.New("User does not exist")
	}

	var err error
	// wrong OTP event after 3 sessions?
	if len(event.Request.Session) >= 3 && !event.Request.Session[len(event.Request.Session)-1].ChallengeResult {
		event.Response.IssueTokens = false
		event.Response.FailAuthentication = true
		err = errors.New("Invalid OTP")

		// correct OTP
	} else if len(event.Request.Session) > 0 && event.Request.Session[len(event.Request.Session)-1].ChallengeResult {
		event.Response.IssueTokens = true
		event.Response.FailAuthentication = false

		// not yet received correct OTP
	} else {
		event.Response.IssueTokens = false
		event.Response.FailAuthentication = false
		event.Response.ChallengeName = "CUSTOM_CHALLENGE"
	}

	return event, err
}

func main() {
	lambda.Start(Handler)
}
