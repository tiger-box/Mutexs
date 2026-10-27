package main

type Notifier interface {
	Send(message string) error
}

type EmailNotifier struct {
}

type SMSNotifier struct {
}

func main() {

}
