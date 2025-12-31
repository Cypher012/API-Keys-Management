package main

import "fmt"

type Notifier interface {
	Send(to, msg string)
}

type EmailNotifier struct{}

func (e EmailNotifier) Send(to, msg string) {
	fmt.Printf("Sending email to %s: %s\n", to, msg)
}

type SMSNotifier struct{}

func (s SMSNotifier) Send(to, msg string) {
	fmt.Printf("Sending SMS to %s: %s\n", to, msg)
}

type PushNotifier struct{}

func (p PushNotifier) Send(to, msg string) {
	fmt.Printf("Sending push notification to %s: %s\n", to, msg)
}

type WhatsAppNotifier struct{}

func (w WhatsAppNotifier) Send(to, msg string) {
	fmt.Printf("Sending WhatsApp message to %s: %s\n", to, msg)
}

type SlackNotifier struct{}

func (s SlackNotifier) Send(to, msg string) {
	fmt.Printf("Sending Slack message to %s: %s\n", to, msg)
}

type UserService struct {
	notifier Notifier
}

func NewUserService(notifier Notifier) *UserService {
	return &UserService{notifier: notifier}
}

func (u *UserService) NotifyUser(to, msg string) {
	u.notifier.Send(to, msg)
}

func main() {
	emailService := NewUserService(EmailNotifier{})
	emailService.NotifyUser("mike@mail.com", "Welcome!")

	smsService := NewUserService(SMSNotifier{})
	smsService.NotifyUser("1234567890", "Your OTP is 1234")

	pushService := NewUserService(PushNotifier{})
	pushService.NotifyUser("device-id", "New alert")

	whatsappService := NewUserService(WhatsAppNotifier{})
	whatsappService.NotifyUser("+234000000", "Hello from WhatsApp")

	slackService := NewUserService(SlackNotifier{})
	slackService.NotifyUser("mike012", "Hello mike")
}
