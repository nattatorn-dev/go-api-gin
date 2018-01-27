package service

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

// CreateConfig func for aws
func CreateConfig(id, secret, token string) *aws.Config {
	AnonymousCredentials := credentials.NewStaticCredentials(id, secret, token)
	config := &aws.Config{
		Region:      aws.String("ap-southeast-1"),
		Credentials: AnonymousCredentials,
	}
	return config
}
