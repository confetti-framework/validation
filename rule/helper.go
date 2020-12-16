package rule

import (
	"github.com/lanvard/support"
	"github.com/uniplaces/carbon"
	"github.com/vigneshuvi/GoDateFormat"
)

func getComparableDate(date *carbon.Carbon, format, zone string) (*carbon.Carbon, error) {
	if date == nil {
		return nil, OptionDateIsRequiredError
	}
	date, err := setFormatAndZone(date, format, zone)
	if err != nil {
		return nil, err
	}
	result, err := generateDate(date.String(), format, zone)
	if err != nil {
		return nil, err
	}
	return setFormatAndZone(result, format, zone)
}

func setFormatAndZone(date *carbon.Carbon, format string, zone string) (*carbon.Carbon, error) {
	date.SetStringFormat(format)
	err := date.SetTimeZone(zone)
	if err != nil {
		return nil, err
	}
	return date, nil
}

func generateDate(value string, format string, zone string) (*carbon.Carbon, error) {
	err := Date{Format: format}.Verify(support.NewValue(value))
	if err != nil {
		return nil, err
	}

	date, err := carbon.CreateFromFormat(format, value, zone)
	if err != nil {
		return nil, err
	}

	return setFormatAndZone(date, format, zone)
}

func normalizeZone(zone string) string {
	if zone == "" {
		zone = "Local"
	}
	return zone
}

func normalizeFormat(format string) string {
	if format == "" {
		format = carbon.DefaultFormat
	}

	return GoDateFormat.ConvertFormat(format)
}
