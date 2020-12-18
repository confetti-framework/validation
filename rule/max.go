package rule

import (
	"github.com/lanvard/support"
	"github.com/lanvard/validation/val_errors"
	"github.com/spf13/cast"
	"reflect"
	"strconv"
)

type Max struct {
	Len int
}

func (m Max) Verify(value support.Value) error {
	input := value.Raw()
	amount := m.getAmount(input)
	if amount > m.Len {
		return val_errors.WithAttributes(
			m.getError(input),
			map[string]string{
				"expect": strconv.Itoa(m.Len),
				"input":  strconv.Itoa(amount),
			},
		)
	}
	return nil
}

func (m Max) getAmount(input interface{}) int {
	switch support.Type(input) {
	case reflect.Slice, reflect.Map, reflect.Array:
		return reflect.ValueOf(input).Len()
	default:
		return cast.ToInt(input)
	}
}

func (m Max) getError(input interface{}) error {
	switch support.Type(input) {
	case reflect.Slice, reflect.Map, reflect.Array:
		return MayNotHaveMoreThanItemsError
	default:
		return MayNotBeGreaterThanError
	}
}
