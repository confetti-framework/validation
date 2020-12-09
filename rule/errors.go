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
var DateMustBeAfterError = ValidationError.Wrap("the :attribute must be after :date, :input_date given")
var DateMustBeAfterOrEqualError = ValidationError.Wrap("the :attribute must be after or equal :date, :input_date given")

// System Error
var OptionDateIsRequired = errors.New("option Date is required")
