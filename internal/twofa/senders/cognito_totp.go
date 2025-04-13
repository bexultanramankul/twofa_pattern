package senders

import (
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"twofa_pattern/internal/model"
)

type AwsCognitoTOTPService struct {
	cognitoClient *cognitoidentityprovider.CognitoIdentityProvider
	userPoolID    string
}

func NewCognitoTOTPService(client *cognitoidentityprovider.CognitoIdentityProvider, userPoolID string) *AwsCognitoTOTPService {
	return &AwsCognitoTOTPService{
		cognitoClient: client,
		userPoolID:    userPoolID,
	}
}

func (s *AwsCognitoTOTPService) SendCode(user model.User, _ string) error {
	//TODO: verify totp code
	return nil
}

func (s *AwsCognitoTOTPService) CanSend(method model.ContactMethod) bool {
	return method == model.ContactMethodTOTP
}
