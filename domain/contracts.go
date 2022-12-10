package domain

type MessageSender interface {
	SendMessage(msg string) error
}

type Message struct{}
type SendReport struct{}
