package aws

import (
	"encoding/base64"
	"errors"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
)

type TokenRetriever interface {
	GetToken(region string) (string, error)
}

type DefaultTokenRetriever struct {
}

func NewDefaultTokenRetriever() *DefaultTokenRetriever {
	return &DefaultTokenRetriever{}
}

func (r *DefaultTokenRetriever) GetToken(region string) (string, error) {
	s := session.Must(session.NewSession(&aws.Config{Region: aws.String(region)}))
	svc := ecr.New(s)
	input := &ecr.GetAuthorizationTokenInput{}
	result, err := svc.GetAuthorizationToken(input)
	if err != nil {
		return "", err
	}

	var tokenBytes []byte

	if tokenBytes, err = base64.StdEncoding.DecodeString(*result.AuthorizationData[0].AuthorizationToken); err != nil {
		return "", err
	}

	token := strings.Split(string(tokenBytes), ":")

	if len(token) != 2 {
		return "", errors.New("token returned from AWS is not valid")
	}

	return token[1], nil
}
