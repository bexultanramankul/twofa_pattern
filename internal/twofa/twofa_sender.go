package twofa

import (
	"errors"
	"twofa_pattern/internal/model"
	"twofa_pattern/pkg/logger"
)

type TwoFaSender struct {
	services []TwoFaService
}

func NewTwoFaSender(services ...TwoFaService) *TwoFaSender {
	return &TwoFaSender{
		services: services,
	}
}

func (s *TwoFaSender) SendCode(user model.User, code string) error {
	logger.Log.Info("Sending 2FA code to user ", user.ID)

	for _, service := range s.services {
		if service.CanSend(user.PreferredContact) {
			return service.SendCode(user, code)
		}
	}

	return errors.New("no 2FA service available for user")
}
