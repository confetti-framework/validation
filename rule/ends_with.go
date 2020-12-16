package rule

import (
	"github.com/lanvard/support"
	"github.com/lanvard/validation/val_errors"
	"strings"
)

type Ends struct {
	with []string
}

func (e Ends) With(with ...string) Ends {
	e.with = with
	return e
}

func (e Ends) Verify(value support.Value) error {
	if len(e.with) == 0 {
		return OptionWithIsRequiredError
	}
	for _, expect := range e.with {
		if strings.HasSuffix(value.String(), expect) {
			return nil
		}
	}
	return val_errors.WithAttributes(
		MuseEndWithError,
		map[string]string{
			"expect": strings.Join(e.with, " or "),
			"input":  value.String(),
		},
	)
}
