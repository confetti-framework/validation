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
	for _, field := range keys {
		value, err := input.GetE(field)
		present := err == nil

		for _, rule := range getAllRequiredRules(verification) {
			err := verifyRule(field, present, value, rule)
			if err != nil {
				result = append(result, errors.WithStack(err))
				break
			}
		}
	}

	return result
}

func verifyRule(field string, present bool, value support.Value, rule inter.Rule) error {
	if !needToVerify(present, rule) {
		return nil
	}

	return val_errors.WithAttribute(rule.Verify(value), map[string]string{"attribute": field})
}

func getAllRequiredRules(verification Verification) []inter.Rule {
	var result []inter.Rule
	for _, baseRule := range verification.Rules {
		if baseRule, ok := baseRule.(inter.RuleWithRequirements); ok {
			result = append(result, baseRule.Require()...)
		}

		result = append(result, baseRule)
	}
	return result
}

func needToVerify(present bool, rule inter.Rule) bool {
	_, isPresentRule := rule.(rules.Present)

	// If the field is neither present nor required to be present,
	// we do not need to validate further
	if !isPresentRule && !present {
		return false
	}

	// If we only need to check if the value is present, and the value
	// is present, we do not need to validate further
	if isPresentRule && present {
		return false
	}

	return true
}
