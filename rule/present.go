package rule

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/support"
)

type Present struct{}

func (p Present) Requirements() []inter.Rule {
	return []inter.Rule{
		Present{},
	}
}

func (p Present) Verify(support.Value) error {
	return MustBePresentError
}
