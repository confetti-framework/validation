package rule

import "github.com/lanvard/errors"

var MustBePresent = errors.New("field %s must be present")
var IsRequiredError = errors.New("field %s is required")
var MustBeAccepted = errors.New("field %s must be accepted")
