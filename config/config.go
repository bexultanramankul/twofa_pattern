package config

import (
	"os"
	"twofa_pattern/pkg/logger"

	"github.com/joho/godotenv"
)

type Config struct {
	// AWS Configuration
	AwsAccessKeyID       string
	AwsSecretAccessKey   string
	AwsRegion            string
	AwsSmsSenderId       string
	AwsSesSenderEmail    string
	AwsCognitoUserPoolId string

	// Twilio Configuration
	TwilioAccountSID string
	TwilioAuthToken  string
	TwilioServiceSID string
	TwilioFromNumber string

	// SendGrid Configuration (optional)
	SendGridApiKey string

	GoogleSenderEmail     string
	GoogleCredentialsPath string
	GoogleTokenPath       string
}

func InitConfig() *Config {
	if err := godotenv.Load(); err != nil {
		logger.Log.Println("No .env file found, using system environment variables")
	}

	return &Config{
		// AWS Config
		AwsAccessKeyID:       getEnv("AWS_ACCESS_KEY_ID", ""),
		AwsSecretAccessKey:   getEnv("AWS_SECRET_ACCESS_KEY", ""),
		AwsRegion:            getEnv("AWS_REGION", "us-east-1"),
		AwsSmsSenderId:       getEnv("AWS_SMS_SENDER_ID", ""),
		AwsSesSenderEmail:    getEnv("AWS_SES_SENDER_EMAIL", ""),
		AwsCognitoUserPoolId: getEnv("AWS_COGNITO_USER_POOL_ID", ""),

		// Twilio Config
		TwilioAccountSID: getEnv("TWILIO_ACCOUNT_SID", ""),
		TwilioAuthToken:  getEnv("TWILIO_AUTH_TOKEN", ""),
		TwilioServiceSID: getEnv("TWILIO_SERVICE_SID", ""),
		TwilioFromNumber: getEnv("TWILIO_FROM_NUMBER", ""),

		// SendGrid Config
		SendGridApiKey: getEnv("SENDGRID_API_KEY", ""),

		GoogleSenderEmail:     getEnv("GOOGLE_SENDER_EMAIL", ""),
		GoogleCredentialsPath: getEnv("GOOGLE_CREDENTIALS_PATH", ""),
		GoogleTokenPath:       getEnv("GOOGLE_TOKEN_PATH", ""),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
