package main

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/rs/zerolog/log"
	"os"
)

const (
	AWS_ACCCESS_KEY_ID    = "AWS_ACCESS_KEY_ID"
	AWS_SECRET_ACCESS_KEY = "AWS_SECRET_ACCESS_KEY"
	AWS_SESSION_TOKEN     = "AWS_SESSION_TOKEN"
	CORNERSTONE_AWS_ARN   = "CORNERSTONE_AWS_ARN"
)

func main() {
	os.Unsetenv(AWS_ACCCESS_KEY_ID)
	os.Unsetenv(AWS_SECRET_ACCESS_KEY)
	os.Unsetenv(AWS_SESSION_TOKEN)


	svc := sts.New(session.New())

	input := &sts.GetSessionTokenInput{
		DurationSeconds: aws.Int64(60*60*24),
		SerialNumber:    aws.String(getARN()),
		TokenCode:       aws.String(getMFACode()),
	}

	result, err := svc.GetSessionToken(input)
	if err != nil {
		log.Fatal().Err(err).Msg("could not get session token")
	}

	toPrint := fmt.Sprintf(
		"export %s=%s\nexport %s=%s\nexport %s=%s", 
		AWS_ACCCESS_KEY_ID, *result.Credentials.AccessKeyId, 
		AWS_SECRET_ACCESS_KEY, *result.Credentials.SecretAccessKey,
		AWS_SESSION_TOKEN, *result.Credentials.SessionToken)
	
	fmt.Println(toPrint)
	}


func getARN() string {
	v := os.Getenv(CORNERSTONE_AWS_ARN)
	if v == "" {
		log.Fatal().Err(errors.New(fmt.Sprintf("%s not set", CORNERSTONE_AWS_ARN))).Msg("")
	}

	return v
}

func getMFACode() string {
	if len(os.Args) < 2 {
		log.Fatal().Err(errors.New("could not get MFA code")).Msg("did you pass the MFA code as a CLI arg?")
	}
	v := os.Args[1]
	return v
}
