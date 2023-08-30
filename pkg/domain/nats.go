package domain

type NatsMessage struct {
	Subject string
	Data    []byte
}
