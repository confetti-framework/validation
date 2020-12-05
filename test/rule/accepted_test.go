package rule

import (
	"errors"
	"github.com/lanvard/validation/rule"
	"github.com/lanvard/validation/val"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_accepted_field_not_present(t *testing.T) {
	errs := val.Validate(
		nil,
		val.Verify("title", rule.Accepted{}),
	)
	require.Len(t, errs, 1)
	require.EqualError(t, errs[0], "field title must be present")
}

func Test_accepted_field_present_but_empty_string(t *testing.T) {
	errs := val.Validate(
		map[string]string{"title": ""},
		val.Verify("title", rule.Accepted{}),
	)
	require.Len(t, errs, 1)
	require.True(t, errors.Is(errs[0], rule.MustBeAccepted))
}

func Test_accepted_field_present_with_string_yes(t *testing.T) {
	errs := val.Validate(
		map[string]string{"title": "yes"},
		val.Verify("title", rule.Accepted{}),
	)
	require.Len(t, errs, 0)
}

func Test_accepted_field_present_with_string_on(t *testing.T) {
	errs := val.Validate(
		map[string]string{"title": "on"},
		val.Verify("title", rule.Accepted{}),
	)
	require.Len(t, errs, 0)
}
