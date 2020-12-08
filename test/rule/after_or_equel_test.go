package rule

// func Test_after_or_equal_or_equal_field_not_present(t *testing.T) {
// 	errs := val.Validate(
// 		nil,
// 		val.Verify("start_date", rule.AfterOrEqual{}),
// 	)
// 	require.Len(t, errs, 0)
// }
//
// func Test_after_or_equal_field_no_options(t *testing.T) {
// 	value := support.NewValue("2021")
// 	err := rule.AfterOrEqual{}.Verify(value)
// 	require.EqualError(t, err, "option Date is required")
// }
//
// func Test_after_or_equal_tomorrow(t *testing.T) {
// 	value := support.NewValue(carbon.Now().String())
// 	err := rule.AfterOrEqual{Date: carbon.Now().SubDay()}.Verify(value)
// 	require.Nil(t, err)
// }
//
// func Test_after_or_equal_with_equal_date(t *testing.T) {
// 	date := carbon.Now()
// 	value := support.NewValue(date.String())
// 	err := rule.AfterOrEqual{Date: date}.Verify(value)
// 	require.Nil(t, err)
// }
//
// func Test_after_or_equal_not_after(t *testing.T) {
// 	value := support.NewValue(carbon.Now().String())
// 	err := rule.AfterOrEqual{Date: carbon.Now().AddDay()}.Verify(value)
// 	require.NotNil(t, err)
// 	require.Regexp(t, `the :attribute must be after \d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}`, err.Error())
// }
//
// func Test_after_or_equal_with_input_format(t *testing.T) {
// 	input := carbon.Now()
// 	input.SetStringFormat(carbon.HourMinuteFormat)
// 	value := support.NewValue(input.String())
//
// 	err := rule.AfterOrEqual{
// 		Date:   carbon.Now().AddDay(),
// 		Format: carbon.HourMinuteFormat,
// 	}.Verify(value)
//
// 	require.NotNil(t, err)
// 	require.Regexp(t, `the :attribute must be after \d{2}:\d{2}`, err.Error())
// }
//
// func Test_after_or_equal_with_clear_format(t *testing.T) {
// 	input := carbon.Now()
// 	input.SetStringFormat(carbon.HourMinuteFormat)
// 	value := support.NewValue(input.String())
//
// 	err := rule.AfterOrEqual{
// 		Date:     carbon.Now().AddDay(),
// 		Format:   "HH:MM",
// 		TimeZone: "UTC",
// 	}.Verify(value)
//
// 	require.NotNil(t, err)
// 	require.Regexp(t, `the :attribute must be after \d{2}:\d{2}`, err.Error())
// }
//
// func Test_after_or_equal_with_timezone(t *testing.T) {
// 	input := carbon.Now()
// 	_ = input.SetTimeZone("UTC")
// 	value := support.NewValue(input.String())
//
// 	err := rule.AfterOrEqual{
// 		Date:     carbon.Now().AddSeconds(5),
// 		TimeZone: "UTC",
// 	}.Verify(value)
//
// 	require.NotNil(t, err)
// 	require.Contains(t, err.Error(), `the :attribute must be after`)
// }
