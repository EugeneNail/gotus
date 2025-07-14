package transport

// A Payload is the interface that provides a set of validation rules based on payload type
type Payload interface {
	Rules()
}
