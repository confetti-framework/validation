package rule

import (
	"regexp"

	"github.com/confetti-framework/support"
)

type Regex struct {
	Expression *regexp.Regexp
	err        error
}

// Compiles Regular expression
// Returns error in Verify() if given Expression is invalid
func (r Regex) Match(exp string) Regex {
	var err error

	r.Expression, err = regexp.Compile(exp)
	r.err = err

	return r
}

func (r Regex) Verify(value support.Value) error {
	// Error compiling the regex
	if r.err != nil {
		return MustCompileRegexError
	}
	if ok := r.Expression.MatchString(value.String()); !ok {
		return MustMatchRegexError
	}
	return nil
}
