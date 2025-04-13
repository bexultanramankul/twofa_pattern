package senders

import (
	"github.com/twilio/twilio-go"
	verify "github.com/twilio/twilio-go/rest/verify/v2"
	"twofa_pattern/config"
	"twofa_pattern/internal/model"
)

type TwilioSmsService struct {
	client     *twilio.RestClient
	fromNumber string
}

func NewTwilioSmsService(client *twilio.RestClient, fromNumber string) *TwilioSmsService {
	return &TwilioSmsService{
		client:     client,
		fromNumber: fromNumber,
	}
}

func (s *TwilioSmsService) SendCode(user model.User, code string) error {
	params := &verify.CreateVerificationParams{}
	params.SetTo(user.Phone)
	params.SetChannel("sms")

	_, err := s.client.VerifyV2.CreateVerification(config.InitConfig().TwilioServiceSID, params)
	return err
}

func (s *TwilioSmsService) CanSend(method model.ContactMethod) bool {
	return method == model.ContactMethodSMS
}
