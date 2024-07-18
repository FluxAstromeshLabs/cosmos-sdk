package types

import (
	"github.com/cosmos/gogoproto/proto"
)

var EventStreamSingleton interface{}

type EventStreamI interface {
	EmitCosmosEvents(events ...proto.Message)
}

type FluxEventManager struct {
	beginBlockEvents []proto.Message
	txEvents         []proto.Message
	endBlockEvents   []proto.Message
}

func NewFluxEventManager() *FluxEventManager {
	return &FluxEventManager{
		beginBlockEvents: []proto.Message{},
		txEvents:         []proto.Message{},
		endBlockEvents:   []proto.Message{},
	}
}

func (fem *FluxEventManager) AddBeginBlockEvents(em EventManagerI) {
	fem.beginBlockEvents = append(fem.beginBlockEvents, em.TypedEvents()...)
}

func (fem *FluxEventManager) FlushBeginBlockEvents() {
	EventStreamSingleton.(EventStreamI).EmitCosmosEvents(fem.beginBlockEvents...)
	fem.beginBlockEvents = []proto.Message{}
}

func (fem *FluxEventManager) ClearBeginBlockEvents() {
	fem.beginBlockEvents = []proto.Message{}
}

func (fem *FluxEventManager) AddTxEvents(em EventManagerI) {
	fem.txEvents = append(fem.txEvents, em.TypedEvents()...)
}

func (fem *FluxEventManager) FlushTxEvents() {
	EventStreamSingleton.(EventStreamI).EmitCosmosEvents(fem.txEvents...)
	fem.txEvents = []proto.Message{}
}

func (fem *FluxEventManager) ClearTxEvents() {
	fem.txEvents = []proto.Message{}
}

func (fem *FluxEventManager) AddEndBlockEvents(em EventManagerI) {
	fem.endBlockEvents = append(fem.endBlockEvents, em.TypedEvents()...)
}

func (fem *FluxEventManager) FlushEndBlockEvents() {
	EventStreamSingleton.(EventStreamI).EmitCosmosEvents(fem.endBlockEvents...)
	fem.endBlockEvents = []proto.Message{}
}

func (fem *FluxEventManager) ClearEndBlockEvents() {
	fem.endBlockEvents = []proto.Message{}
}

var FluxEventManagerSingleton = NewFluxEventManager()
