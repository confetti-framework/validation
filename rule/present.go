package rule

import "github.com/lanvard/support"

type Present struct{}

func (p Present) NeedToBePresent() {}

func (p Present) Verify(support.Value) error {
	return MustBePresent
}
