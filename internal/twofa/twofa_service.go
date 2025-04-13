package twofa

import "twofa_pattern/internal/model"

type TwoFaService interface {
	SendCode(user model.User, code string) error
	CanSend(preferredContact model.ContactMethod) bool
}
