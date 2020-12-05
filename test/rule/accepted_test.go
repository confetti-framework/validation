package rule

import (
	"github.com/lanvard/support"
	"github.com/lanvard/validation/rule"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_accepted_field_not_present(t *testing.T) {
	situation := situation{
		rule:    rule.Accepted{},
		present: false,
		value:   support.NewValue(nil),
	}
	require.False(t, situation.isValid())
	require.Equal(t, rule.IsRequiredError, situation.error())
}

func Test_accepted_field_present_but_empty_string(t *testing.T) {
	situation := situation{
		rule:    rule.Accepted{},
		present: true,
		value:   support.NewValue(nil),
	}
	require.False(t, situation.isValid())
	require.Equal(t, rule.MustBeAccepted, situation.error())
}

func Test_accepted_field_present_with_string_yes(t *testing.T) {
	situation := situation{
		rule:    rule.Accepted{},
		present: true,
		value:   support.NewValue("yes"),
	}
	require.True(t, situation.isValid())
	require.Nil(t, situation.error())
}

func Test_accepted_field_present_with_string_on(t *testing.T) {
	situation := situation{
		rule:    rule.Accepted{},
		present: true,
		value:   support.NewValue("on"),
	}
	require.True(t, situation.isValid())
	require.Nil(t, situation.error())
}
