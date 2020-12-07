package rule

import (
	"github.com/lanvard/errors"
	"github.com/lanvard/syslog/log_level"
	net "net/http"
)

// Validation Error
var ValidationError = errors.New("").Status(net.StatusUnprocessableEntity).Level(log_level.INFO)
var MustBePresentError = ValidationError.Wrap("the :attribute must be present")
var IsRequiredError = ValidationError.Wrap("the :attribute is required")
var MustBeAcceptedError = ValidationError.Wrap("the :attribute must be accepted")
var DateMustBeError = ValidationError.Wrap("the :attribute must be after :date")

// System Error
var NoOptionsGivenError = errors.New("options are required")
