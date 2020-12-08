package rule

import (
	"github.com/lanvard/support"
	"github.com/uniplaces/carbon"
)

type AfterOrEqual struct {
	Date     *carbon.Carbon
	Format   string
	TimeZone string
}

func (a AfterOrEqual) Verify(value support.Value) error {

	// if !after.GreaterThan(a.Date) {
	// 	return val_errors.WithAttribute(DateMustBeError, map[string]string{"date": a.Date.Format(a.Format)})
	// }
	return nil
}
