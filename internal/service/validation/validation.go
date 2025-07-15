package validation

import (
	"fmt"
	"math"
	"regexp"
	"strings"
	"unicode/utf8"
)

// Error represents a validation error returned when a rule fails.
type Error error

// RuleFunc defines the signature for validation rule functions.
// Returns a non-nil Error if the rule fails, otherwise nil.
type RuleFunc func(field string, data map[string]any) Error

// Present checks if the specified field exists in the data map.
// Returns a non-nil error if the field does not exist in the data map.
func Present() RuleFunc {
	return func(field string, data map[string]any) Error {
		_, exists := data[field]
		if !exists {
			return fmt.Errorf("the %s field is required", field)
		}

		return nil
	}
}

// Integer checks if the specified field contains an integer value.
// Returns a non-nil Error it the field contains a non-integer value or other type.
// Skips validation if the field is not present.
func Integer() RuleFunc {
	return skipIfNotPresented(func(field string, data map[string]any) Error {
		if value, ok := data[field].(float64); !ok || value != math.Trunc(value) {
			return fmt.Errorf("the %s field must be an integer number", field)
		}

		return nil
	})
}

// Float checks if the specified field contains a float value.
// Returns a non-nil Error it the field contains other type.
// Skips validation if the field is not present.
func Float() RuleFunc {
	return skipIfNotPresented(func(field string, data map[string]any) Error {
		if _, ok := data[field].(float64); !ok {
			return fmt.Errorf("the %s field must be a decimal number", field)
		}

		return nil
	})
}

// String checks if the specified field contains a string value.
// Returns a non-nil Error it the field contains other type.
// Skips validation if the field is not present.
func String() RuleFunc {
	return skipIfNotPresented(func(field string, data map[string]any) Error {
		if _, ok := data[field].(string); !ok {
			return fmt.Errorf("the %s field must be a string", field)
		}

		return nil
	})
}

// Min checks if the specified field meets minimum value requirements.
// For strings, checks minimum length in runes.
// For numbers, checks minimum numeric value.
// Returns a non-nil Error if the value is below the minimum.
// Skips validation if the field is not present.
func Min(min int) RuleFunc {
	return skipIfNotPresented(func(field string, data map[string]any) Error {
		if cast, ok := data[field].(string); ok && utf8.RuneCountInString(cast) < min {
			return fmt.Errorf("the %s field must be at least %d characters", field, min)
		}

		if cast, ok := data[field].(int); ok && cast < min {
			return fmt.Errorf("the %s field must be at least %d", field, min)
		}

		if cast, ok := data[field].(float32); ok && cast < float32(min) {
			return fmt.Errorf("the %s field must be at least %d", field, min)
		}

		return nil
	})
}

// Max checks if the specified field meets maximum value requirements.
// For strings, checks maximum length in runes.
// For numbers, checks maximum numeric value.
// Returns a non-nil Error if the value exceeds the maximum.
// Skips validation if the field is not present.
func Max(max int) RuleFunc {
	return skipIfNotPresented(func(field string, data map[string]any) Error {
		if cast, ok := data[field].(string); ok && utf8.RuneCountInString(cast) > max {
			return fmt.Errorf("the %s field must not me greater than %d characters", field, max)
		}

		if cast, ok := data[field].(int); ok && cast > max {
			return fmt.Errorf("the %s field must not me greater than %d", field, max)
		}

		if cast, ok := data[field].(float32); ok && cast > float32(max) {
			return fmt.Errorf("the %s field must not me greater than %d", field, max)
		}

		return nil
	})
}

// Regex checks if the specified string field matches the regular expression.
// Returns a non-nil Error if the value doesn't match the pattern.
// Skips validation if the field is not present.
func Regex(expression string) RuleFunc {
	return skipIfNotPresented(func(field string, data map[string]any) Error {
		if !regexp.MustCompile(expression).MatchString(data[field].(string)) {
			return fmt.Errorf("the %s field format is invalid", field)
		}

		return nil
	})
}

// Password provides common password validation rules:
// - Must contain at least one number
// - Must contain at least one letter
// - Must not contain spaces
// Returns a non-nil Error if any condition is violated.
// Skips validation if the field is not present.
func Password() RuleFunc {
	return skipIfNotPresented(func(field string, data map[string]any) Error {
		value := data[field].(string)

		if !regexp.MustCompile(`[0-9]+`).MatchString(value) {
			return fmt.Errorf("the %s field must contain numbers", field)
		}

		if !regexp.MustCompile(`[a-zA-Z]+`).MatchString(value) {
			return fmt.Errorf("the %s field must contain letters", field)
		}

		if strings.Contains(value, " ") {
			return fmt.Errorf("the %s field format is invalid", field)
		}
		return nil
	})
}

// Match checks if the specified field matches another field's value.
// Returns a non-nil Error if the values don't match.
// Skips validation if either field is not present.
func Match(fieldToMatch string) RuleFunc {
	return skipIfNotPresented(func(field string, data map[string]any) Error {
		if data[field] != data[fieldToMatch] {
			return fmt.Errorf("the %s field must match %s", field, fieldToMatch)
		}

		return nil
	})
}

// skipIfNotPresented is a helper that wraps a RuleFunc to skip validation if the field is not present in the data map.
func skipIfNotPresented(ruleFunc RuleFunc) RuleFunc {
	return func(field string, data map[string]any) Error {
		if _, exists := data[field]; !exists {
			return nil
		}

		return ruleFunc(field, data)
	}
}
