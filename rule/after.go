package rule

import (
	"github.com/lanvard/support"
	"github.com/lanvard/validation/val_errors"
	"github.com/uniplaces/carbon"
	"github.com/vigneshuvi/GoDateFormat"
)

type After struct {
	Date     *carbon.Carbon
	Format   string
	TimeZone string
}

func (a After) Verify(value support.Value) error {
	if a.Date == nil {
		return OptionDateIsRequired.Wrap("the :attribute with rule.After")
	}

	if a.Format == "" {
		a.Format = carbon.DefaultFormat
	}

	a.Format = GoDateFormat.ConvertFormat(a.Format)

	if a.TimeZone == "" {
		a.TimeZone = "Local"
	}
	err := a.Date.SetTimeZone(a.TimeZone)
	if err != nil {
		return err
	}

	after, err := carbon.CreateFromFormat(a.Format, value.String(), a.TimeZone)
	if err != nil {
		return err
	}

	if !after.GreaterThan(a.Date) {
		return val_errors.WithAttribute(DateMustBeError, map[string]string{"date": a.Date.Format(a.Format)})
	}
	return nil
}
