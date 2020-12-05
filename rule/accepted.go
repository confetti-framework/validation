package rule

import "github.com/lanvard/support"

type Accepted struct{}

func (a Accepted) Valid(present bool, value support.Value) bool {
	if !present {
		return false
	}
	return value.Bool()
}

func (a Accepted) Error(present bool, value support.Value) error {
	if a.Valid(present, value) {
		return nil
	}
	if !present {
		return IsRequiredError
	}
	return MustBeAccepted
}
