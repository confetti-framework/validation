package rule

import (
	"github.com/uniplaces/carbon"
	"github.com/vigneshuvi/GoDateFormat"
)

func normalizeInputDate(value string, format string, zone string) (*carbon.Carbon, error) {
	return carbon.CreateFromFormat(format, value, zone)
}

func normalizeCompareDate(date *carbon.Carbon, zone string) (*carbon.Carbon, error) {
	if date == nil {
		return nil, OptionDateIsRequired
	}
	err := date.SetTimeZone(zone)
	if err != nil {
		return nil, err
	}
	return date, nil
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
