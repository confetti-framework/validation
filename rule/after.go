package rule

import (
	"github.com/lanvard/support"
	"github.com/lanvard/validation/val_errors"
	"github.com/uniplaces/carbon"
)

type After struct {
	Date     *carbon.Carbon
	Format   string
	TimeZone string
}

func (a After) Verify(value support.Value) error {
	format := normalizeFormat(a.Format)
	zone := normalizeZone(a.TimeZone)
	compareTo, err := getComparableDate(a.Date, format, zone)
	if err != nil {
		return err
	}
	input, err := generateDate(value.String(), format, zone)
	if err != nil {
		return err
	}

	if !input.GreaterThan(compareTo) {
		return val_errors.WithAttributes(
			DateMustBeAfterError,
			map[string]string{"date": compareTo.String(), "input_date": input.String()},
		)
	}

	return nil
}
