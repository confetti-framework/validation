package rule

import (
	"github.com/lanvard/support"
	"github.com/lanvard/validation/rule"
	"github.com/lanvard/validation/val"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_present(t *testing.T) {
	errors := val.Validate(
		support.NewValue(map[string]string{"age": "60"}),
		val.Verify("age", rule.Present{}),
	)
	require.Empty(t, errors)
}

func Test_not_present(t *testing.T) {
	errors := val.Validate(
		support.NewValue(map[string]string{"age": "60"}),
		val.Verify("name", rule.Present{}),
	)
	require.EqualError(t, errors[0], "field name must be present")
}
