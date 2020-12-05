package rule

import "github.com/lanvard/support"

type Present struct{}

func (p Present) Valid(present bool, _ support.Value) bool {
	return present
}

func (p Present) Error(_ bool, _ support.Value) error {
	return MustBePresent
}
