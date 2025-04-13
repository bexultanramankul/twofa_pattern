package senders

import (
	"github.com/aws/aws-sdk-go/service/ses"
	"twofa_pattern/internal/model"
)

type AwsSesEmailService struct {
	sesClient *ses.SES
	sender    string
}

func NewSesEmailService(sesClient *ses.SES, senderEmail string) *AwsSesEmailService {
	return &AwsSesEmailService{
		sesClient: sesClient,
		sender:    senderEmail,
	}
}

func (s *AwsSesEmailService) SendCode(user model.User, code string) error {
	//TODO: verify email code

	//_, err := s.sesClient.SendEmail(&ses.SendEmailInput{
	//	Destination: &ses.Destination{
	//		ToAddresses: []*string{aws.String(user.Email)},
	//	},
	//	Message: &ses.Message{
	//		Body: &ses.Body{
	//			Text: &ses.Content{
	//				Data: aws.String("Your verification code: " + code),
	//			},
	//		},
	//		Subject: &ses.Content{
	//			Data: aws.String("Your Verification Code"),
	//		},
	//	},
	//	Source: aws.String(s.sender),
	//})
	//return err

	return nil
}

func (s *AwsSesEmailService) CanSend(method model.ContactMethod) bool {
	return method == model.ContactMethodEmail
}
