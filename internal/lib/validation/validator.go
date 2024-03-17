package validator

import (
	"errors"
	"regexp"
	"time"
	"unicode/utf8"

	"github.com/go-openapi/strfmt"
)

var (
	ValidSex  = []string{"male", "female"}
	ValidRole = []string{"user", "admin"}

	ValidEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	ErrIncorrectTime = errors.New("invalid time, it should be less then now")
)

func Matches(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}

func MinChars(value string, n int) bool {
	return utf8.RuneCountInString(value) >= n
}

func MaxChars(value string, n int) bool {
	return utf8.RuneCountInString(value) <= n
}

func StringValueBetween(value string, mn, mx int) bool {
	if utf8.RuneCountInString(value) < mn {
		return false
	}
	if utf8.RuneCountInString(value) > mx {
		return false
	}

	return true
}

func PermittedValue[T comparable](value T, permittedValues ...T) bool {
	for i := range permittedValues {
		if value == permittedValues[i] {
			return true
		}
	}
	return false
}

func ParseTime(check *strfmt.Date) (*time.Time, error) {
	now := time.Now()

	strTime := check.String()

	tm, err := time.Parse("2006-01-02", strTime)
	if err != nil {
		return nil, err
	}

	if tm.After(now) {
		return nil, ErrIncorrectTime
	}

	return &tm, nil
}
