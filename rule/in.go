package rule

import (
	"github.com/lanvard/support"
)

type In struct {
	with []interface{}
}

func (i In) Verify(value support.Value) error {
	if len(i.with) == 0 {
		return OptionWithIsRequiredError
	}
	for _, compare := range i.with {
		if value.Raw() == compare {
			return nil
		}
	}
	return SelectedIsInvalid
}

func (i In) With(with ...interface{}) In {
	i.with = with
	return i
}

