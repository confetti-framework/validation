package rule

import (
	"github.com/uniplaces/carbon"
	"github.com/vigneshuvi/GoDateFormat"
)

func getComparableDate(date *carbon.Carbon, format, zone string) (string, error) {
	if date == nil {
		return "", OptionDateIsRequired
	}
	date.SetStringFormat(format)
	err := date.SetTimeZone(zone)
	if err != nil {
		return "", err
	}
	return date.String(), nil
}

func normalizeDate(value string, format string, zone string) (*carbon.Carbon, error) {
	return carbon.CreateFromFormat(format, value, zone)
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
