package rule

import "github.com/lanvard/support"

type Accepted struct{}

func (a Accepted) NeedToBePresent() {}

func (a Accepted) Verify(value support.Value) error {
	if !value.Bool() {
		return MustBeAccepted
	}
	return nil
}
