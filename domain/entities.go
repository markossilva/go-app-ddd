package domain

type Notification struct {
	ID      int
	Subject string
	Message string
}

func NewNotification() Notification {
	return Notification{}
}

func (n Notification) MethodWhenNecessary() {}
