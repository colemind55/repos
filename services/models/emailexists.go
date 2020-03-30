package models

import (
	"allsup.assessment/api/services/db"
)

// Email ...
type Email struct {
	Address string
	CaseID string
}

// Validate if the email address already exists ...
func Validate(email string, caseID string) (bool, error) {
	
	isExists, err := db.ValidateEmailAddress(email, caseID)

	if (isExists) {
		return true, nil
	} 

	return false, err
}
