package rule

import (
	"github.com/confetti-framework/support"
	"github.com/confetti-framework/validation/rule"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_in_array_with_valid_inputs_same_type(t *testing.T) {
	value := support.NewValue("test_string")
	err := rule.
		InArray{}.
		Array("not_valid", "invalid", "test_string").
		Verify(value)

	require.Nil(t, err)
}

func Test_invalid_input_same_type(t *testing.T) {
	value := support.NewValue("test_string")
	err := rule.
		InArray{}.
		Array("not_valid", "invalid", "still_invalid").
		Verify(value)

	require.Equal(t, "the :attribute field does not exist in :other", err.Error())
}

type testInArrayStruct struct {
	HelloWorld uint
}

func Test_valid_input_different_types(t *testing.T) {
	value := support.NewValue("test_string")
	err := rule.
		InArray{}.
		Array(1234, 123.45, true, []string{"hello", "there"}, new(testInArrayStruct), "test_string").
		Verify(value)

	require.Nil(t, err)
}

func Test_invalid_input_different_types(t *testing.T) {
	value := support.NewValue("test_string")
	err := rule.
		InArray{}.
		Array(1234, 123.45, true, []string{"hello", "there"}, new(testInArrayStruct)).
		Verify(value)

	require.Equal(t, "the :attribute field does not exist in :other", err.Error())
}

func Test_empty_input(t *testing.T) {
	value := support.NewValue("test_string")
	err := rule.
		InArray{}.
		Array().
		Verify(value)

	require.Equal(t, "the :attribute field does not exist in :other", err.Error())
}
