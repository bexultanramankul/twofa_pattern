package senders

import (
	"github.com/aws/aws-sdk-go/service/sns"
	"twofa_pattern/internal/model"
)

type AwsSnsSmsService struct {
	snsClient *sns.SNS
	senderID  string
}

func NewSnsSmsService(snsClient *sns.SNS, senderID string) *AwsSnsSmsService {
	return &AwsSnsSmsService{
		snsClient: snsClient,
		senderID:  senderID,
	}
}

func (s *AwsSnsSmsService) SendCode(user model.User, code string) error {
	//TODO: verify sms code

	//_, err := s.snsClient.Publish(&sns.PublishInput{
	//	Message:     aws.String("Your verification code: " + code),
	//	PhoneNumber: aws.String(user.Phone),
	//	MessageAttributes: map[string]*sns.MessageAttributeValue{
	//		"AWS.SNS.SMS.SenderID": {
	//			DataType:    aws.String("String"),
	//			StringValue: aws.String(s.senderID),
	//		},
	//	},
	//})
	//return err

	return nil
}

func (s *AwsSnsSmsService) CanSend(method model.ContactMethod) bool {
	return method == model.ContactMethodSMS
}
