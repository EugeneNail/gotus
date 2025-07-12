package api

type Payload interface {
	Validate() bool
}
