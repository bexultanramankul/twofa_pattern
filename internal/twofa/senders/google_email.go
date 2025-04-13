package senders

import (
	"context"
	"encoding/base64"
	"fmt"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
	"os"

	"twofa_pattern/internal/model"
	"twofa_pattern/internal/util"
)

type GoogleEmailService struct {
	srv  *gmail.Service
	from string
}

func NewGoogleEmailService(fromEmail, credentialsPath, tokenPath string) (*GoogleEmailService, error) {
	ctx := context.Background()

	b, err := os.ReadFile(credentialsPath)
	if err != nil {
		return nil, fmt.Errorf("unable to read credentials.json: %w", err)
	}

	config, err := google.ConfigFromJSON(b, gmail.GmailSendScope)
	if err != nil {
		return nil, fmt.Errorf("unable to parse client secret: %w", err)
	}

	tok, err := util.TokenFromFile(tokenPath)
	if err != nil {
		return nil, fmt.Errorf("unable to get token: %w", err)
	}

	client := config.Client(ctx, tok)

	srv, err := gmail.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return nil, fmt.Errorf("unable to create Gmail client: %w", err)
	}

	return &GoogleEmailService{
		srv:  srv,
		from: fromEmail,
	}, nil
}

func (s *GoogleEmailService) SendCode(user model.User, code string) error {
	var message gmail.Message

	emailTo := user.Email
	subject := "Verification Code"
	body := fmt.Sprintf("Your verification code is: %s", code)

	raw := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s",
		s.from, emailTo, subject, body)

	message.Raw = base64.URLEncoding.EncodeToString([]byte(raw))

	_, err := s.srv.Users.Messages.Send("me", &message).Do()
	if err != nil {
		return fmt.Errorf("unable to send email: %w", err)
	}
	return nil
}

func (s *GoogleEmailService) CanSend(method model.ContactMethod) bool {
	return method == model.ContactMethodEmail
}
