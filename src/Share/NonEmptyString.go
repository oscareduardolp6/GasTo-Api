package share

import "strings"

type NonEmptyString string

func (val NonEmptyString) Value() string {
	return string(val)
}

type EmptyString struct{}

func (err EmptyString) Error() string {
	return "The string cannot be empty"
}

func CreateNonEmptyString(str string) (NonEmptyString, error) {
	cleaned := strings.TrimSpace(str)
	if cleaned == "" {
		return "", EmptyString{}
	}
	return NonEmptyString(str), nil
}
