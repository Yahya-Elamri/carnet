package utils

import (
	"fmt"
	"regexp"

	"gorm.io/gorm"
)

func IsValidEmail(email string) bool {
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

func IsValidUsername(username string) bool {
	const usernameRegex = `^[a-zA-Z0-9_-]+$`
	re := regexp.MustCompile(usernameRegex)
	return re.MatchString(username)
}

func IsFieldUnique(fieldName, fieldValue string, db *gorm.DB, tables ...string) struct {
	Message string
	Class   string
	Valid   bool
} {
	var validation struct {
		Message string
		Class   string
		Valid   bool
	}

	// Validate email format if the field is email
	if fieldName == "email" && !IsValidEmail(fieldValue) {
		validation.Message = "Write a valid email: example@email.com"
		validation.Class = "text-red-700"
		validation.Valid = false
		return validation
	}

	if fieldName == "username" && !IsValidUsername(fieldValue) {
		validation.Message = "Write a valid username: example123"
		validation.Class = "text-red-700"
		validation.Valid = false
		return validation
	}

	// Iterate over the tables to check for field uniqueness
	for _, table := range tables {
		var count int64
		err := db.Table(table).Where(fmt.Sprintf("%s = ?", fieldName), fieldValue).Count(&count).Error
		if err != nil {
			validation.Message = "Error retry again"
			validation.Class = "text-red-700"
			validation.Valid = false
			return validation
		}
		if count > 0 {
			validation.Message = fmt.Sprintf("%s already exists", fieldValue)
			validation.Class = "text-red-700"
			validation.Valid = false
			return validation
		}
	}

	// If field value is not found in any of the tables
	validation.Message = fmt.Sprintf("%s is good", fieldValue)
	validation.Class = "text-green-700"
	validation.Valid = true
	return validation
}
