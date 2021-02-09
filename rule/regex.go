package rule

import (
	"regexp"

	"github.com/confetti-framework/support"
)

type Regex struct {
	Expression *regexp.Regexp
	Error      error
}

// Compiles Regular expression
// Returns error if given Expression is invalid
func (r Regex) Match(exp string) Regex {
	var err error

	r.Expression, err = regexp.Compile(exp)
	r.Error = err

	return r
}

func (r Regex) Verify(value support.Value) error {
	// Error compiling the regex
	if r.Error != nil {
		return MustCompileRegexError
		//return r.Error
	}
	if ok := r.Expression.MatchString(value.String()); !ok {
		//return fmt.Errorf("format invalid")
		return MustMatchRegexError
	}
	return nil
}
