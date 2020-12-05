package rule

import "github.com/lanvard/support"

type Required struct{}

func (r Required) NeedToBePresent() {}

func (r Required) Verify(value support.Value) error {
	if !value.Filled() {
		return MustBePresent
	}
	return nil
}
