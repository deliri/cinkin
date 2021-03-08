package utils

import "regexp"

// this regular expression is from the .NET library itself
var Re = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// ValidateEmail is used to compare a form email input against a regular expression
func ValidateEmail(email string) bool {
	return Re.MatchString(email)
}
