package types

type EventStreamI interface {
	ForwardEvents(events ...interface{})
	FinalizeEvents() error
}

var EventStreamSingleton EventStreamI

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
	EventStreamSingleton.ForwardEvents(fem.beginBlockEvents...)
	fem.beginBlockEvents = []interface{}{}
}

func (fem *FluxEventManager) ClearBeginBlockEvents() {
	fem.beginBlockEvents = []interface{}{}
}

func (fem *FluxEventManager) AddTxEvents(em EventManagerI) {
	fem.txEvents = append(fem.txEvents, em.GenericEvents()...)
}

func (fem *FluxEventManager) FlushTxEvents() {
	EventStreamSingleton.ForwardEvents(fem.txEvents...)
	fem.txEvents = []interface{}{}
}

func (fem *FluxEventManager) ClearTxEvents() {
	fem.txEvents = []interface{}{}
}

func (fem *FluxEventManager) AddEndBlockEvents(em EventManagerI) {
	fem.endBlockEvents = append(fem.endBlockEvents, em.GenericEvents()...)
}

func (fem *FluxEventManager) FlushEndBlockEvents() {
	EventStreamSingleton.ForwardEvents(fem.endBlockEvents...)
	fem.endBlockEvents = []interface{}{}
}

func (fem *FluxEventManager) ClearEndBlockEvents() {
	fem.endBlockEvents = []interface{}{}
}

func (fem *FluxEventManager) FinalizeEvents() error {
	return EventStreamSingleton.FinalizeEvents()
}

var FluxEventManagerSingleton = NewFluxEventManager()
