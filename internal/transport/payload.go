package transport

import (
	"github.com/EugeneNail/gotus/internal/service/validation"
)

// A Payload is the interface that provides a set of validation rules based on payload type
type Payload interface {
	Rules() map[string][]validation.RuleFunc
}
