package helpers

import "strings"

func FormatResponseValidation(err string) interface{} {
	boolPrefix := strings.Contains(err, ";")
	if boolPrefix {
		return strings.Split(err, ";")
	}
	return err
}
