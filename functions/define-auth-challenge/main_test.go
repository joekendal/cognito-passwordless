package main_test

import (
	main "cognito-passwordless/define-auth-challenge"
	"errors"
	"log"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestUserNotFound(t *testing.T) {
	event, _ := main.Handler(events.CognitoEventUserPoolsDefineAuthChallenge{
		Request: events.CognitoEventUserPoolsDefineAuthChallengeRequest{
			UserNotFound: true,
		},
	})
	if event.Response.IssueTokens || !event.Response.FailAuthentication {
		log.Fatal(errors.New("Expected UserNotFound to fail to issue token and to fail authentication"))
	}
}

func TestWrongOTP(t *testing.T) {
	event, _ := main.Handler(events.CognitoEventUserPoolsDefineAuthChallenge{
		Request: events.CognitoEventUserPoolsDefineAuthChallengeRequest{
			Session: []*events.CognitoEventUserPoolsChallengeResult{
				{ChallengeResult: false}, {ChallengeResult: false}, {ChallengeResult: false},
			},
		},
	})
	if event.Response.IssueTokens || !event.Response.FailAuthentication {
		log.Fatal(errors.New("Expected invalid OTP to fail authentication and not issue token"))
	}
}

func TestCorrectOTP(t *testing.T) {
	event, _ := main.Handler(events.CognitoEventUserPoolsDefineAuthChallenge{
		Request: events.CognitoEventUserPoolsDefineAuthChallengeRequest{
			Session: []*events.CognitoEventUserPoolsChallengeResult{
				{ChallengeResult: false}, {ChallengeResult: false}, {ChallengeResult: true},
			},
		},
	})
	if !event.Response.IssueTokens || event.Response.FailAuthentication {
		log.Fatal(errors.New("Expected correct OTP to approve authentication and issue token"))
	}
}

func TestNoneCorrectYet(t *testing.T) {
	event, _ := main.Handler(events.CognitoEventUserPoolsDefineAuthChallenge{
		Request: events.CognitoEventUserPoolsDefineAuthChallengeRequest{
			Session: []*events.CognitoEventUserPoolsChallengeResult{
				{ChallengeResult: false}, {ChallengeResult: false},
			},
		},
	})
	if event.Response.IssueTokens || event.Response.FailAuthentication || !(event.Response.ChallengeName == "CUSTOM_CHALLENGE") {
		log.Fatal(errors.New("Expected < 3 incorrect OTP to not fail authentication and not issue token"))
	}
}
