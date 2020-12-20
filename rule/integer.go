package rule

import (
	"github.com/lanvard/support"
)

type Integer struct{}

func (i Integer) Verify(value support.Value) error {
	switch value.Raw().(type) {
	case nil, bool, float32, float64:
		return MustBeAnInteger
	}

	_, err := value.IntE()
	if err != nil {
		return MustBeAnInteger
	}

	return nil
}
