package rule

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/support"
)

type Accepted struct{}

func (a Accepted) Requirements() []inter.Rule {
	return []inter.Rule{
		Present{},
	}
}

func (a Accepted) Verify(value support.Value) error {
	if !value.Bool() {
		return MustBeAcceptedError
	}
	return nil
}
