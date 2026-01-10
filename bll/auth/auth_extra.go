package bll


import (
	// "encoding/json"
	"strings"
	"strconv"
	"net/mail"
	"unicode"
	"golang.org/x/crypto/bcrypt"
	"errors"
	models "teamplayer/models"
	viewmodels "teamplayer/models/viewmodels"
	"github.com/google/uuid"
	// "github.com/gofiber/contrib/websocket"
	// dal "teamplayer/dal"
)

// for the email make a list of acceptable email extensions
// create use ful check for the extensions


func ValidateNameEntry(SignUpField ,label, name string, maxNameLength int) error {
	if EmptyString(name) {
		return &models.NewUserValidationError{
			SignUpField:   SignUpField,
			ValidationMessage: label + " is required. No empty entry.",
		} 
	}

	if CheckMaxLength(maxNameLength, name) {
		return &models.NewUserValidationError{
			SignUpField:   SignUpField,
			ValidationMessage: label + " must be " + strconv.Itoa(maxNameLength) + " characters or fewer.",
		} 
	}

	if SpecialCharacter(name) {
		return &models.NewUserValidationError{
			SignUpField:   SignUpField,
			ValidationMessage: label + " contains symbols, digits, and control characters.",
		} 
	}

	return nil 
}

func ValidateEmailEntry(SignUpField ,label, emailAddress string, maxEmailLength int) error {
	email := strings.ToLower(strings.TrimSpace(emailAddress))
	if EmptyString(email) {
		return &models.NewUserValidationError{
			SignUpField:   SignUpField,
			ValidationMessage: label + " is required. No empty entry.",
		} 
	}

	if CheckMaxLength(maxEmailLength, email) {
		return &models.NewUserValidationError{
			SignUpField:   SignUpField,
			ValidationMessage: label + " must be " + strconv.Itoa(maxEmailLength) + " characters or fewer.",
		} 
	}

	if IsInvalidEmail(email) {
		return &models.NewUserValidationError{
			SignUpField:   SignUpField,
			ValidationMessage: label + " is not a valid email address.",
		} 
	}

	return nil 
}

func ValidateUsernameEntry(SignUpField, label, userName string, minUserNameLength, maxUserNameLength int) error {
	username := strings.ToLower(strings.TrimSpace(userName))	

	if EmptyString(username) {
		return &models.NewUserValidationError{
			SignUpField:   SignUpField,
			ValidationMessage: label + " is required. No empty entry.",
		} 
	}

	if CheckUserNameLength(minUserNameLength , maxUserNameLength, username) {
		return &models.NewUserValidationError{
			SignUpField:   SignUpField,
			ValidationMessage: label + " must be " + strconv.Itoa(minUserNameLength) + " to "+ strconv.Itoa(maxUserNameLength) + " characters.",
		} 
	}

	if IsInvalidUserNameCharacters(username) {
		return &models.NewUserValidationError{
			SignUpField:   SignUpField,
			ValidationMessage: label + " is not a valid username.",
		} 
	}

	return nil 
}

func ValidatePasswordEntry(SignUpField, label, password string, minPasswordLength int) error {
	if EmptyString(password) {
		return &models.NewUserValidationError{
			SignUpField:   SignUpField,
			ValidationMessage: label + " is required. No empty entry.",
		} 
	}

	if CheckMinLength(minPasswordLength, password) {
		return &models.NewUserValidationError{
			SignUpField:   SignUpField,
			ValidationMessage: label + " must be at least " + strconv.Itoa(minPasswordLength) + " characters.",
		} 
	}

	if IsInvalidPasswordCharacters(password) {
		return &models.NewUserValidationError{
			SignUpField:   SignUpField,
			ValidationMessage: label + " contains invalid control characters.",
		} 
	}

	return nil 
}

func EmptyString(s string) bool {
	return strings.TrimSpace(s) == ""
}

func IsInvalidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err != nil
}

func CheckMaxLength(maxNameLength int, entry string) bool {
	return len([]rune(entry)) > maxNameLength
}

func CheckMinLength(minPasswordLength int, entry string) bool {
	return len([]rune(entry)) < minPasswordLength
}

func CheckUserNameLength(minUserNameLength , maxNameLength int, entry string) bool {
	return len([]rune(entry)) > maxNameLength ||len([]rune(entry)) < minUserNameLength 
}


func SpecialCharacter(name string) bool {
	for _, r := range name {
		if unicode.IsLetter(r) || r == ' ' || r == '-' || r == '\'' || r == '’' {
			continue
		}
		return true
	}
	return false
}

func IsInvalidUserNameCharacters(userName string) bool {
	for _, r := range userName {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' || r == '.' || r == '-' {
			continue
		}
		return true
	}
	return false	
}

func IsInvalidPasswordCharacters(password string) bool {
	for _, r := range password {
		if unicode.IsControl(r){
			return true
		}
	}
	return false	
}

func HashPassword(password string) (string,error) {
	// Cost: bcrypt.DefaultCost (10) is okay; 12 is a common stronger default.
	const cost = 12

	hashBytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", err
	}
	return string(hashBytes), nil
}

func CreateAccError(SignUpErr error, SignUpVM viewmodels.SignUpViewModel) viewmodels.SignUpViewModel {
	var ve *models.NewUserValidationError
	if errors.As(SignUpErr, &ve) {
		switch ve.SignUpField {
		case "first_name":
			SignUpVM.Errors.FirstName = ve.ValidationMessage
		case "last_name":
			SignUpVM.Errors.LastName = ve.ValidationMessage
		case "email":
			SignUpVM.Errors.Email = ve.ValidationMessage
		case "username":
			SignUpVM.Errors.UserName = ve.ValidationMessage
		case "password":
			SignUpVM.Errors.Password = ve.ValidationMessage
		}
	}

	return SignUpVM
}





func IsValidUUID(userUUID uuid.UUID) bool {
	// if true no UUID is present
	return userUUID == uuid.Nil
}