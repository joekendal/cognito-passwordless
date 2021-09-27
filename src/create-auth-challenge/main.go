package main

import (
	"log"

	"crypto/rand"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sns/snsiface"
)

var (
	Svc snsiface.SNSAPI
)

func init() {
	sess := session.New()
	Svc = sns.New(sess)
}

// Send code over SMS via Amazon Simple Notification Service (SNS)
func SendSMS(message string, phoneNumber string) error {
	_, err := Svc.Publish(&sns.PublishInput{
		Message:     &message,
		PhoneNumber: &phoneNumber,
	})
	return err
}

// Generate random One-time Passcode (OTP)
func GenerateOTP() (string, error) {
	buffer := make([]byte, 6)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}
	for i := 0; i < 6; i++ {
		buffer[i] = "1234567890"[int(buffer[i])%10]
	}
	return string(buffer), nil
}

func Handler(event events.CognitoEventUserPoolsCreateAuthChallenge) (events.CognitoEventUserPoolsCreateAuthChallenge, error) {
	phoneNumber := event.Request.UserAttributes["phone_number"]

	var err error
	var secretCode string
	if len(event.Request.Session) == 0 {
		// generate a new secret login code and send it to user
		secretCode, err = GenerateOTP()
		if err != nil {
			// crypto failure
			log.Print(err)
			return event, err
		}
		err = SendSMS(secretCode, phoneNumber)
		if err != nil {
			// sms failure
			log.Print(err)
		}
	} else {
		// re-use code generated in previous challenge
		previousChallenge := event.Request.Session[len(event.Request.Session)-1]
		secretCode = previousChallenge.ChallengeMetadata
	}

	event.Response.PublicChallengeParameters = map[string]string{
		"phone": phoneNumber,
	}

	// add the login code to the private challenge params
	// so it can be verified by the "Verify Auth Challenge Response" trigger
	event.Response.PrivateChallengeParameters = map[string]string{
		"secretCode": secretCode,
	}

	// add the login code to the session so it is available
	// in a next invocation of the "Create Auth Challenge" trigger
	event.Response.ChallengeMetadata = secretCode

	return event, err
}

func main() {
	lambda.Start(Handler)
}
