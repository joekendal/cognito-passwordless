package main_test

import (
	main "cognito-passwordless/create-auth-challenge"
	"log"
	"os"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sns/snsiface"
)

type mockSvc struct {
	snsiface.SNSAPI
}

func (m *mockSvc) Publish(*sns.PublishInput) (*sns.PublishOutput, error) {
	return &sns.PublishOutput{}, nil
}

func TestHandler(t *testing.T) {
	testPhone := "+11111111111"
	event, err := main.Handler(events.CognitoEventUserPoolsCreateAuthChallenge{
		Request: events.CognitoEventUserPoolsCreateAuthChallengeRequest{
			UserAttributes: map[string]string{
				"phone_number": testPhone,
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	// check response
	if event.Response.PublicChallengeParameters["phone"] != testPhone {
		log.Fatal("Expected phone number in public challenge params")
	}
	if event.Response.PrivateChallengeParameters["secretCode"] == "" {
		log.Fatal("Expected pass code to be set in private challenge params")
	}
	if event.Response.ChallengeMetadata == "" {
		log.Fatal("Expected challenge metadata to be set")
	}
}

func TestMain(t *testing.M) {
	main.Svc = &mockSvc{}

	code := t.Run()

	os.Exit(code)
}
