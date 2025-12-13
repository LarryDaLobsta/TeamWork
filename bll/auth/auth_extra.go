import (
	// "encoding/json"
	"fmt"
	"html"
	"log"
	"strings"
	// "github.com/gofiber/contrib/websocket"
	// dal "teamplayer/dal"
)




func ( response NewUserValidationError) Error() string {
	return e.Message
}

func ValidateNameEntry(field ,label, name string, maxNameLength int) error {
	if EmptyString(name) {
		return NewUserValidationError{
			Field:   field,
			Message: label + " is required. No empty entry.",
		} 
	}

	if CheckLength(maxNameLength, name) {
		return NewUserValidationError{
			Field:   field,
			Message: label + " must be " + strconv.Itoa(maxNameLength) + " characters or fewer."
		} 
	}

	if SpecialCharacter(name) {
		return NewUserValidationError{
			Field:   field,
			Message: label + " contains symbols, digits, and control characters.",
		} 
	}

	return nil 
}


func EmptyString(s string) bool {
	return strings.TrimSpace(s) == ""
}


func CheckLength(maxNameLength int, entry string) bool {
	return len([]rune(entry)) > maxNameLength
}


func SpecialCharacter(name string) bool {
	trimmedName := strings.TrimSpace(name)
	for _, r := range trimmedName {
		if unicode.IsLetter(r) || r == ' ' || r == '-' || r == '\'' || r == '’' {
			continue
		}
		return true
	}
	return false
}