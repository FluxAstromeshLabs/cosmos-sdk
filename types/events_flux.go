package types

import (
	"github.com/cosmos/gogoproto/proto"
)

var EventStreamSingleton interface{}

type EventStreamI interface {
	EmitCosmosEvents(events ...proto.Message)
}

type FluxEventManager struct {
	TypedEvents []proto.Message
}

func NewFluxEventManager() *FluxEventManager {
	return &FluxEventManager{
		TypedEvents: []proto.Message{},
	}
}

func (fem *FluxEventManager) AddTypedEvents(em EventManagerI) {
	fem.TypedEvents = append(fem.TypedEvents, em.TypedEvents()...)
}

func (fem *FluxEventManager) Flush() {
	EventStreamSingleton.(EventStreamI).EmitCosmosEvents(fem.TypedEvents...)
	fem.TypedEvents = []proto.Message{}
}

var FluxEventManagerSingleton = NewFluxEventManager()
