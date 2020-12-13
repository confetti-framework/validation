package rule

import (
	"github.com/lanvard/support"
	"github.com/lanvard/validation/val_errors"
	"github.com/uniplaces/carbon"
)

type Date struct {
	Format string
}

func (d Date) Verify(value support.Value) error {
	format := normalizeFormat(d.Format)
	_, err := carbon.CreateFromFormat(format, value.String(), "")
	if err != nil {
		return val_errors.WithAttributes(
			DateNotValidFormatError,
			map[string]string{
				"example": format,
				"input":   value.String(),
			},
		)
	}

	return nil
}
