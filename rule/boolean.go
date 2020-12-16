package rule

import (
	"github.com/lanvard/support"
	"github.com/lanvard/support/str"
	"github.com/lanvard/validation/val_errors"
)

type Boolean struct{}

func (b Boolean) Verify(value support.Value) error {
	if str.InSlice(value.Raw(), "0", "1", 0, 1, true, false) {
		return nil
	}

	input := value.String()
	if value.Raw() == nil {
		input = "nil"
	}

	return val_errors.WithAttributes(
		MuseBeABooleanError,
		map[string]string{"input": input},
	)
}
