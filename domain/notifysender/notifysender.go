package notifysender

import (
	"github.com/markossilva/go-app-ddd/domain"
	"github.com/markossilva/go-app-ddd/infra"
)

type Service struct {
	sender infra.EmailSender
}

func NewService(sender infra.EmailSender) Service {
	return Service{
		sender: sender,
	}
}

func (s Service) SendNotifications(notifys []domain.Notification) {}
