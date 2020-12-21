package rule

import (
	"github.com/confetti-framework/support"
)

type Filled struct{}

func (f Filled) Verify(value support.Value) error {
	if !value.Filled() {
		return MustHaveAValueError
	}
	return nil
}
