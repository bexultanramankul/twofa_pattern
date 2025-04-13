package senders

import (
	"github.com/twilio/twilio-go"
	"twofa_pattern/internal/model"
)

type TwilioEmailService struct {
	client     *twilio.RestClient
	serviceSid string
}

func NewTwilioEmailService(client *twilio.RestClient, serviceSid string) *TwilioEmailService {
	return &TwilioEmailService{
		client:     client,
		serviceSid: serviceSid,
	}
}

func (s *TwilioEmailService) SendCode(user model.User, code string) error {
	//TODO: verify email code
	return nil
}

func (s *TwilioEmailService) CanSend(method model.ContactMethod) bool {
	return method == model.ContactMethodEmail
}
