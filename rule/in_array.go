package rule

import (
	"github.com/confetti-framework/support"
)

type InArray struct {
	inArray []interface{}
	err     error
}

// Receives slice of interface
func (i InArray) Array(values ...interface{}) InArray {
	i.inArray = values
	return i
}

// Compare each slice item with the searched value
// All of them in interface
func (i InArray) Verify(value support.Value) error {
	for _, val := range i.inArray {
		if value.Source() == val {
			return nil
		}
	}
	return MustBeInArrayError
}
