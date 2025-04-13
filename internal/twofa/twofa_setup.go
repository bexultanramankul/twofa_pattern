package twofa

import (
	"github.com/twilio/twilio-go"
	"twofa_pattern/config"
	"twofa_pattern/internal/twofa/senders"
	"twofa_pattern/pkg/logger"
)

func NewAwsTwoFaSetup() *TwoFaSender {
	cfg := config.InitConfig()

	//sess := session.Must(session.NewSessionWithOptions(session.Options{
	//	SharedConfigState: session.SharedConfigEnable,
	//}))

	//snsClient := sns.New(sess)
	//sesClient := ses.New(sess)
	//cognitoClient := cognitoidentityprovider.New(sess)
	twilioClient := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: cfg.TwilioAccountSID,
		Password: cfg.TwilioAuthToken,
	})

	//awsSms := senders.NewSnsSmsService(snsClient, os.Getenv("AWS_SMS_SENDER_ID"))
	//awsEmail := senders.NewSesEmailService(sesClient, os.Getenv(cfg.AwsSesSenderEmail))
	//awsTotp := senders.NewCognitoTOTPService(cognitoClient, os.Getenv("AWS_COGNITO_USER_POOL_ID"))
	//twilioEmail := senders.NewTwilioEmailService(twilioClient, os.Getenv("TWILIO_VERIFY_SERVICE_SID"))
	twilioSms := senders.NewTwilioSmsService(twilioClient, cfg.TwilioFromNumber)
	googleEmail, err := senders.NewGoogleEmailService(
		cfg.GoogleSenderEmail,
		cfg.GoogleCredentialsPath,
		cfg.GoogleTokenPath,
	)
	if err != nil {
		logger.Log.Fatalf("Failed to create GoogleEmailService: %v", err)
	}

	return NewTwoFaSender(twilioSms, googleEmail)
}
