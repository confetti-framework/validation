package rule

import (
	"github.com/lanvard/support"
)

type Filled struct{}

func (f Filled) Verify(value support.Value) error {
	if !value.Filled() {
		return MustHaveAValue
	}
	return nil
}
