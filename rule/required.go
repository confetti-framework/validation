package rule

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/support"
)

type Required struct{}

func (r Required) Require() []inter.Rule {
	return []inter.Rule{
		Present{},
	}
}

func (r Required) Verify(value support.Value) error {
	if !value.Filled() {
		return IsRequiredError
	}
	return nil
}
