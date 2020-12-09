package rule

import (
	"github.com/lanvard/support"
	"github.com/lanvard/validation/val_errors"
	"github.com/uniplaces/carbon"
)

type BeforeOrEqual struct {
	Date     *carbon.Carbon
	Format   string
	TimeZone string
}

func (b BeforeOrEqual) Verify(value support.Value) error {
	format := normalizeFormat(b.Format)
	zone := normalizeZone(b.TimeZone)
	compareTo, err := getComparableDate(b.Date, format, zone)
	if err != nil {
		return err
	}
	input, err := generateDate(value.String(), format, zone)
	if err != nil {
		return err
	}

	if !input.LessThanOrEqualTo(compareTo) {
		return val_errors.WithAttributes(
			DateMustBeBeforeOrEqualError,
			map[string]string{"date": compareTo.String(), "input_date": input.String()},
		)
	}

	return nil
}
