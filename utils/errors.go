package utils

import "strings"

func SplitValidationErrors(errors string) []string {
	return strings.Split(errors, "; ")
}
