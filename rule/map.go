package rule

import (
	"github.com/lanvard/support"
	"reflect"
)

type Map struct{}

func (m Map) Verify(value support.Value) error {
	kind := support.Kind(value.Raw())
	if kind != reflect.Map {
		return MustBeAMap
	}
	return nil
}
