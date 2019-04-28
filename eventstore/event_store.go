package eventstore

import (
	"github.com/google/uuid"
	"github.com/evalsocket/gophers-utility/domain"
)

// EventStore methods allow to save, load events and event streams
type EventStore interface {
	Store([]*domain.Event) error
	Get(uuid.UUID) (*domain.Event, error)
	FindAll() []*domain.Event
	GetStream(uuid.UUID, string) []*domain.Event
}
