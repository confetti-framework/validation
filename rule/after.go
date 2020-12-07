package rule

import (
	"github.com/lanvard/support"
)

type After struct {
}

func (a After) Verify(value support.Value) error {
	return nil
}
