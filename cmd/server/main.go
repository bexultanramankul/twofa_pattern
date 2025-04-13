package main

import (
	"fmt"
	"log"
	"twofa_pattern/config"
	"twofa_pattern/internal/model"
	"twofa_pattern/internal/twofa"
	"twofa_pattern/pkg/logger"
)

func main() {
	logger.InitLogger()
	config.InitConfig()

	// Инициализация 2FA sender (AWS версия)
	sender := twofa.NewAwsTwoFaSetup()

	// Пример пользователя с SMS
	smsUser := model.User{
		ID:               "user-123",
		PreferredContact: model.ContactMethodSMS,
		Phone:            "+77478867173",
	}

	// Пример пользователя с Email
	emailUser := model.User{
		ID:               "user-456",
		PreferredContact: model.ContactMethodEmail,
		Email:            "bexultan.ramankul@gmail.com",
	}

	// Тестовый код (в реальном приложении генерируется случайно)
	testCode := "654321"

	// Отправка SMS
	fmt.Println("Sending SMS code...")
	if err := sender.SendCode(smsUser, testCode); err != nil {
		log.Printf("Failed to send SMS: %v", err)
	} else {
		fmt.Println("SMS sent successfully!")
	}

	// Отправка Email
	fmt.Println("\nSending Email code...")
	if err := sender.SendCode(emailUser, testCode); err != nil {
		log.Printf("Failed to send Email: %v", err)
	} else {
		fmt.Println("Email sent successfully!")
	}
}
