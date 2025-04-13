package totp

import (
	"time"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

func GenerateSecret(account string) (string, error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "MyApp",
		AccountName: account,
	})
	if err != nil {
		return "", err
	}
	return key.Secret(), nil
}

func GenerateCode(secret string) (string, error) {
	code, err := totp.GenerateCode(secret, time.Now())
	if err != nil {
		return "", err
	}
	return code, nil
}

func VerifyCode(secret, code string) bool {
	valid, _ := totp.ValidateCustom(code, secret, time.Now(), totp.ValidateOpts{
		Period:    30,
		Skew:      1,
		Digits:    otp.DigitsSix,
		Algorithm: otp.AlgorithmSHA1,
	})
	return valid
}
