package val_errors

import (
	"fmt"
	"github.com/lanvard/errors"
	"strings"
)

func FindError(errs []error, search string) error {
	for _, err := range errs {
		errField, _ := FindField(err)
		if search == errField {
			return err
		}
	}
	return nil
}

func FindField(err error) (string, bool) {
	var fieldHolder *withField

	ok := errors.As(err, &fieldHolder)
	if !ok {
		return "", false
	}

	return fieldHolder.field, true
}

func WithField(err error, field string) error {
	if err == nil {
		return nil
	}
	return &withField{
		err,
		field,
	}
}

type withField struct {
	cause error
	field string
}

func (w *withField) Error() string {
	err := w.cause.Error()
	if strings.Contains(err, "%") {
		return fmt.Sprintf(err, w.field)
	}
	return err
}

func (w *withField) Format(st fmt.State, verb rune) {
	errors.Format(st, verb, w)
}

func (w *withField) Unwrap() error {
	return w.cause
}
