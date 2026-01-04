package event

import "time"

type TransactionCreated struct {
	Name    string
	Payload interface{}
}

func NewTransactionCreatedEvent() *TransactionCreated {
	return &TransactionCreated{
		Name: "TransactionCreated",
	}
}

func (e *TransactionCreated) GetName() string {
	return e.Name
}

func (e *TransactionCreated) GetPayload() interface{} {
	return e.Payload
}

func (e *TransactionCreated) SetPayload(payload interface{}) error {
	e.Payload = payload
	return nil
}

func (e *TransactionCreated) GetDateTime() time.Time {
	return time.Now()
}
