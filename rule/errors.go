package rule

import (
	"github.com/lanvard/errors"
	"github.com/lanvard/syslog/log_level"
	net "net/http"
)

var ValidationError = errors.New("").Status(net.StatusUnprocessableEntity).Level(log_level.INFO)
var MustBePresent = ValidationError.Wrap("field %s must be present")
var IsRequiredError = ValidationError.Wrap("field %s is required")
var MustBeAccepted = ValidationError.Wrap("field %s must be accepted")
