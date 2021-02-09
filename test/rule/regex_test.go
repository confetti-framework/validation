package rule

import (
	"github.com/confetti-framework/support"
	"github.com/confetti-framework/validation/rule"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_regex_with_valid_expression_valid_input(t *testing.T) {
	value := support.NewValue("abcd")
	err := rule.Regex{}.
		Match("[a-z]{1,}").
		Verify(value)

	require.Nil(t, err)
}

func Test_regex_with_valid_advanced_expression_valid_input(t *testing.T) {
	value := support.NewValue("conf_etti@github.com")
	err := rule.
		Regex{}.
		Match(`[\w]{3,15}@[a-z]{3,5}\.com`).
		Verify(value)

	require.Equal(t, "the :attribute format is invalid", err.Error())
}

func Test_regex_with_invalid_expression(t *testing.T) {
	value := support.NewValue("abcd")
	err := rule.
		Regex{}.
		Match(`\k`).
		Verify(value)

	require.Equal(t, "the :expect isn't a valid regex", err.Error())
}

func Test_regex_with_valid_expression_invalid_input(t *testing.T) {
	value := support.NewValue("aa@abc.com.br")
	err := rule.
		Regex{}.
		Match(`[\w]{3,15}@[a-z]{3,5}\.com`).
		Verify(value)

	require.Equal(t, "the :attribute format is invalid", err.Error())
}
