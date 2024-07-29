package types

var EventStreamSingleton interface{}

type EventStreamI interface {
	ForwardEvents(events ...interface{})
}

type FluxEventManager struct {
	beginBlockEvents []interface{}
	txEvents         []interface{}
	endBlockEvents   []interface{}
}

func NewFluxEventManager() *FluxEventManager {
	return &FluxEventManager{
		beginBlockEvents: []interface{}{},
		txEvents:         []interface{}{},
		endBlockEvents:   []interface{}{},
	}
}

func (fem *FluxEventManager) AddBeginBlockEvents(em EventManagerI) {
	fem.beginBlockEvents = append(fem.beginBlockEvents, em.GenericEvents()...)
}

func (fem *FluxEventManager) FlushBeginBlockEvents() {
	EventStreamSingleton.(EventStreamI).ForwardEvents(fem.beginBlockEvents...)
	fem.beginBlockEvents = []interface{}{}
}

func (fem *FluxEventManager) ClearBeginBlockEvents() {
	fem.beginBlockEvents = []interface{}{}
}

func (fem *FluxEventManager) AddTxEvents(em EventManagerI) {
	fem.txEvents = append(fem.txEvents, em.GenericEvents()...)
}

func (fem *FluxEventManager) FlushTxEvents() {
	EventStreamSingleton.(EventStreamI).ForwardEvents(fem.txEvents...)
	fem.txEvents = []interface{}{}
}

func (fem *FluxEventManager) ClearTxEvents() {
	fem.txEvents = []interface{}{}
}

func (fem *FluxEventManager) AddEndBlockEvents(em EventManagerI) {
	fem.endBlockEvents = append(fem.endBlockEvents, em.GenericEvents()...)
}

func (fem *FluxEventManager) FlushEndBlockEvents() {
	EventStreamSingleton.(EventStreamI).ForwardEvents(fem.endBlockEvents...)
	fem.endBlockEvents = []interface{}{}
}

func (fem *FluxEventManager) ClearEndBlockEvents() {
	fem.endBlockEvents = []interface{}{}
}

var FluxEventManagerSingleton = NewFluxEventManager()
