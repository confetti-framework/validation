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
var DateMustBeAfterError = ValidationError.Wrap("the :attribute must be after :date, :input given")
var DateMustBeAfterOrEqualError = ValidationError.Wrap("the :attribute must be after or equal to :date, :input given")
var DateMustBeBeforeError = ValidationError.Wrap("the :attribute must be before :date, :input given")
var DateMustBeBeforeOrEqualError = ValidationError.Wrap("the :attribute must be before or equal to :date, :input given")
var DateMustBeEqualError = ValidationError.Wrap("the :attribute must be equal to :date, :input given")
var DateNotValidFormatError = ValidationError.Wrap("the :attribute is not a valid date (example :example), :input given")
var MuseBeABooleanError = ValidationError.Wrap("the :attribute must be a boolean, :input given")
var MuseEndWithError = ValidationError.Wrap("the :attribute must end with :expect, :input given")
var MustHaveAValue = ValidationError.Wrap("the :attribute field must have a value")
var SelectedIsInvalid = ValidationError.Wrap("the selected :attribute is invalid")
var MustBeAnInteger = ValidationError.Wrap("the :attribute must be an integer")
var MayNotBeGreaterThanError = ValidationError.Wrap("the :attribute may not be greater than :expect, :input given")
var MayNotHaveMoreThanItemsError = ValidationError.Wrap("the :attribute may not have more than :expect items, :input items given")

// System Error
var OptionDateIsRequiredError = errors.New("option Date is required")
var OptionWithIsRequiredError = errors.New("option With is required")
