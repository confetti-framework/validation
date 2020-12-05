package val

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/errors"
	"github.com/lanvard/support"
	rules "github.com/lanvard/validation/rule"
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

	// If the key contains an asterisk, then there are more keys we need to verify
	keys := support.GetSearchableKeysByOneKey(verification.Field, input)
	for _, key := range keys {
		value, err := input.GetE(key)
		present := err == nil

		for _, rule := range verification.Rules {
			err := verifyRule(key, present, value, rule)
			if err != nil {
				result = append(result, errors.WithStack(err))
				break
			}
		}
	}

	return result
}

func verifyRule(key string, present bool, value support.Value, rule inter.Rule) error {

	// Verify that the field must be present
	if needToBePresent(rule) && !present {
		presentRule := rules.Present{}
		return val_errors.WithField((presentRule).Verify(value), key)
	}

	if !needToValidateRule(present, rule) {
		return nil
	}

	return val_errors.WithField(rule.Verify(value), key)
}

func needToValidateRule(present bool, rule inter.Rule) bool {
	// If the field is neither present nor required,
	// we do not need to validate it further
	if !needToBePresent(rule) && !present {
		return false
	}

	// If we only need to check if the value
	// is present, then this rule is valid now
	_, isPresentRule := rule.(rules.Present)
	if isPresentRule {
		return false
	}

	return true
}

func needToBePresent(rule inter.Rule) bool {
	_, needToBePresent := rule.(inter.RuleNeedToBePresent)
	return needToBePresent
}