package rule

import (
	"github.com/lanvard/validation/rule"
	"github.com/lanvard/validation/val"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_after_field_not_present(t *testing.T) {
	errs := val.Validate(
		nil,
		val.Verify("date", rule.After{}),
	)
	require.Len(t, errs, 0)
}

func Test_after_field_no_options(t *testing.T) {
	errs := val.Validate(
		nil,
		val.Verify("date", rule.After{}),
	)
	require.Len(t, errs, 1)
}
