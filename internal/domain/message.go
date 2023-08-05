package domain

import "time"

// Message сущность для отправки сервисом
type Message struct {
	Payload   string
	CreatedAt time.Time
}
