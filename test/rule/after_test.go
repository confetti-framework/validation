package rule

import (
	"github.com/lanvard/support"
	"github.com/lanvard/validation/rule"
	"github.com/lanvard/validation/val"
	"github.com/stretchr/testify/require"
	"github.com/uniplaces/carbon"
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
	value := support.NewValue("2021")
	err := rule.After{}.Verify(value)
	require.EqualError(t, err, "the :attribute with rule.After: options are required")
}

func Test_after_tomorrow(t *testing.T) {
	value := support.NewValue(carbon.Now().String())
	err := rule.After{Date: carbon.Now().SubDay()}.Verify(value)
	require.Nil(t, err)
}

func Test_after_not_after(t *testing.T) {
	value := support.NewValue(carbon.Now().String())
	err := rule.After{Date: carbon.Now().AddDay()}.Verify(value)
	require.NotNil(t, err)
	require.Regexp(t, `the :attribute must be after \d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}`, err.Error())
}

func Test_after_with_input_format(t *testing.T) {
	input := carbon.Now()
	input.SetStringFormat(carbon.HourMinuteFormat)
	value := support.NewValue(input.String())

	err := rule.After{
		Date:   carbon.Now().AddDay(),
		Format: carbon.HourMinuteFormat,
	}.Verify(value)

	require.NotNil(t, err)
	require.Regexp(t, `the :attribute must be after \d{2}:\d{2}`, err.Error())
}

func Test_after_with_timezone(t *testing.T) {
	input := carbon.Now()
	_ = input.SetTimeZone("UTC")
	value := support.NewValue(input.String())

	err := rule.After{
		Date:     carbon.Now().AddSeconds(5),
		TimeZone: "UTC",
	}.Verify(value)

	require.NotNil(t, err)
	require.Contains(t, err.Error(), `the :attribute must be after`)
}
