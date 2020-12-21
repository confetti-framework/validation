package rule

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/support"
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
