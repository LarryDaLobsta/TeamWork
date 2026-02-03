package helpers

import (
	"strings"
	extra "teamplayer/bll/auth"
)

// val := strings.ToLower(strings.TrimSpace(input))

func IsFormBlank(firstName string, lastName string, emailAddress string) bool {
	return extra.EmptyString(firstName) && extra.EmptyString(lastName) && extra.EmptyString(emailAddress)
	// return true
}

// normalize entries for comparison
func NormalizeField(entry string) string {
	return strings.ToLower(strings.TrimSpace(entry))
}

// see if the field is changed or not
func IsFieldChanged(entryOld string, entryNew string) bool {

	// this is set to return true if changed
	return entryOld != entryNew
}
