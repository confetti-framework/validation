package rule

import (
	"github.com/lanvard/support"
	"github.com/lanvard/validation/rule"
	"github.com/lanvard/validation/val"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_required_and_present(t *testing.T) {
	errors := val.Validate(
		support.NewValue(map[string]string{"age": "60"}),
		val.Verify("age", rule.Required{}),
	)
	require.Empty(t, errors)
}

func Test_required_and_not_present(t *testing.T) {
	errors := val.Validate(
		support.NewValue(nil),
		val.Verify("name", rule.Required{}),
	)
	require.EqualError(t, errors[0], "the name must be present")
}

func Test_required_and_present_but_empty(t *testing.T) {
	errors := val.Validate(
		support.NewValue(map[string]string{"age": ""}),
		val.Verify("age", rule.Required{}),
	)
	require.EqualError(t, errors[0], "the age is required")
}
