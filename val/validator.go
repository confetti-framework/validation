package val

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/errors"
	"github.com/lanvard/support"
	"github.com/lanvard/validation/val_errors"
)

func Validate(input interface{}, verifications ...Verification) []error {
	result := []error{}
	value, err := support.NewValueE(input)
	if err != nil {
		return append(result, err)
	}
	for _, verification := range verifications {
		result = append(result, verifyVerification(value, verification)...)
	}
	return result

}

func verifyVerification(input support.Value, verification Verification) []error {
	var result []error
	keys := support.GetSearchableKeysByOneKey(verification.Field, input)
	for _, key := range keys {
		value, err := input.GetE(key)
		present := err == nil

		for _, rule := range verification.Rules {
			err = verifyRule(key, present, value, rule)
			if err != nil {
				result = append(result, errors.WithStack(err))
			}
		}
	}

	return result
}

func verifyRule(key string, present bool, value support.Value, rule inter.Rule) error {
	if !rule.Valid(present, value) {
		return val_errors.WithField(rule.Error(present, value), key)
	}
	return nil
}
