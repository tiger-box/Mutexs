package main

import (
	"errors"
	"fmt"
)

type Notifier interface {
	Send(message string) error
}

type EmailNotifier struct {
	email string
}

type SMSNotifier struct {
	sms string
}

func (s *SMSNotifier) Send(message string) error {
	if s.sms == "" {
		return fmt.Errorf("empty sms")
	}
	fmt.Println(s.sms, message)
	return nil
}

func (e *EmailNotifier) Send(message string) error {

	fmt.Println(e.email, message)
	if e.email == "" {
		return fmt.Errorf("empty email")
	}
	return nil
}

func NotifyAll(notifier []Notifier, message string) error {
	var err error
	for _, v := range notifier {
		if v.Send(message) != nil {

			err = errors.Join(err, v.Send(message))
			fmt.Println(err)

		}
	}
	return err
}

func main() {

	var sms Notifier = &SMSNotifier{sms: "+79293715581"}
	var email Notifier = &EmailNotifier{email: "@mailsDogs"}

	notifiers := []Notifier{sms, email}

	err := NotifyAll(notifiers, "Заказ готов")

	if err != nil {
		fmt.Println("error")
	}

}
