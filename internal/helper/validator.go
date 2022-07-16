package helper

import (
	"regexp"
	"strings"
)

// simple check for email
func ValidateEmail(email string) bool {
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return re.MatchString(email)
}

// check if names:
// 1) have spaces
// 2) is less than 2 or more than 40 symbols
func ValidateNames(name string) bool {
	if len(name) < 1 || len(name) > 40 {
		return false
	}

	if strings.ContainsRune(name, ' ') {
		return false
	}

	return true
}

// check if names:
// 1) have spaces
// 2) is less than 7 or more than 40 symbols
func ValidatePassword(password string) bool {
	if len(password) < 7 || len(password) > 40 {
		return false
	}

	if strings.ContainsRune(password, ' ') {
		return false
	}

	return true
}

// check for empty
func ValidateNonEmpty(text ...string) bool {
	for _, s := range text {
		if len(strings.TrimSpace(s)) == 0 {
			return false
		}
	}

	return true
}

// check if names:
// 1) have spaces
// 2) empty
// 2) more than 40 symbols
func ValidateTags(tags ...string) bool {
	for _, tag := range tags {
		if len(tag) < 1 || len(tag) > 40 {
			return false
		}

		if strings.ContainsRune(tag, ' ') {
			return false
		}
	}

	return true
}
