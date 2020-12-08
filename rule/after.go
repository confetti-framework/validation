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

	compareDate, err := normalizeCompareDate(a.Date, zone)
	if err != nil {
		return err
	}

	after, err := normalizeInputDate(value.String(), format, zone)
	if err != nil {
		return err
	}

	if !after.GreaterThan(compareDate) {
		return val_errors.WithAttribute(DateMustBeError, map[string]string{"date": compareDate.Format(format)})
	}

	return nil
}
