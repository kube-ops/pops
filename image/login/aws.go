package login

import "encoding/base64"
import "github.com/aws/aws-sdk-go/aws"
import "github.com/aws/aws-sdk-go/aws/session"
import "github.com/aws/aws-sdk-go/service/ecr"
import "github.com/fsouza/go-dockerclient"
import "github.com/sirupsen/logrus"
import "strings"

// GetAWSCredentials retrieve ecr credentials from AWS
func GetAWSCredentials() docker.AuthConfiguration {
	awsConfig := aws.NewConfig().WithRegion("eu-west-1")
	ecrSvc := ecr.New(session.Must(session.NewSession(awsConfig)))
	tokenOutput, err := ecrSvc.GetAuthorizationToken(&ecr.GetAuthorizationTokenInput{})
	if err != nil {
		logrus.Fatal(err)
	}
	if len(tokenOutput.AuthorizationData) != 1 {
		logrus.Fatal("Error retrieving ECR token")
	}
	authData := *tokenOutput.AuthorizationData[0]
	decoded, err := base64.StdEncoding.DecodeString(*authData.AuthorizationToken)
	if err != nil {
		logrus.Fatal(err)
	}
	splited := strings.Split(string(decoded), ":")
	if len(splited) != 2 {
		logrus.Fatal("Error parsing ECR token response")
	}

	return docker.AuthConfiguration{Username: splited[0], Password: splited[1], Email: *authData.ProxyEndpoint}
}
