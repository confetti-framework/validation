package rule

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/support"
)

type situation struct {
	rule    inter.Rule
	present bool
	value   support.Value
}

func (s situation) isValid() bool {
	return s.rule.Valid(s.present, s.value)
}

func (s situation) error() error {
	return s.rule.Error(s.present, s.value)
}
