package models

import "time"

type Contact struct {
	ID         int    `json:"id"`
	Username   string `json:"username"`
	GivenName  string `json:"givenname"`
	FamilyName string `json:"familyname"`
	Phone      []struct {
		TypeID      int `json:"typeid"`
		CountryCode int `json:"countrycode"`
		Operator    int `json:"operator"`
		Number      int `json:"number"`
	} `json:"phone"`
	Email     []string  `json:"email"`
	Birthdate time.Time `json:"birthday"`
}
