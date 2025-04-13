package model

type ContactMethod string

const (
	ContactMethodSMS   ContactMethod = "sms"
	ContactMethodEmail ContactMethod = "email"
	ContactMethodTOTP  ContactMethod = "totp"
)

type User struct {
	ID               string
	PreferredContact ContactMethod
	Email            string
	Phone            string
	TOTPSecret       string // Для Google Authenticator
}
