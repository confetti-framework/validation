package rule

import "github.com/lanvard/support"

type Required struct{}

func (r Required) Valid(present bool, value support.Value) bool {
	if !present {
		return false
	}
	return value.Filled()
}

func (r Required) Error(present bool, _ support.Value) error {
	return MustBePresent
}
