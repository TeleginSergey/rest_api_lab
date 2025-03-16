package services

import (
	"time"

	"cmd/main.go/internal/models"
)

var contacts = []models.Contact{
	{
		ID:         1,
		Username:   "johndoe",
		GivenName:  "John",
		FamilyName: "Doe",
		Phone: []struct {
			TypeID      int `json:"typeid"`
			CountryCode int `json:"countrycode"`
			Operator    int `json:"operator"`
			Number      int `json:"number"`
		}{
			{
				TypeID:      1,
				CountryCode: 1,
				Operator:    123,
				Number:      5551234,
			},
		},
		Email:     []string{"john.doe@example.com"},
		Birthdate: time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		ID:         2,
		Username:   "janedoe",
		GivenName:  "Jane",
		FamilyName: "Doe",
		Phone: []struct {
			TypeID      int `json:"typeid"`
			CountryCode int `json:"countrycode"`
			Operator    int `json:"operator"`
			Number      int `json:"number"`
		}{
			{
				TypeID:      2,
				CountryCode: 1,
				Operator:    456,
				Number:      5555678,
			},
		},
		Email:     []string{"jane.doe@example.com"},
		Birthdate: time.Date(1995, 5, 15, 0, 0, 0, 0, time.UTC),
	},
}

func GetContacts() []models.Contact {
	return contacts
}

func CreateContact(newContact models.Contact) {
	contacts = append(contacts, newContact)
}

func UpdateContact(id int, key string, value string) (*models.Contact, error) {
	for i, contact := range contacts {
		if contact.ID == id {
			switch key {
			case "username":
				contacts[i].Username = value
			case "givenname":
				contacts[i].GivenName = value
			case "familyname":
				contacts[i].FamilyName = value
			case "email":
				contacts[i].Email = []string{value}
			case "birthday":
				birthdate, err := time.Parse(time.RFC3339, value)
				if err != nil {
					return nil, err
				}
				contacts[i].Birthdate = birthdate
			default:
				return nil, nil
			}
			return &contacts[i], nil
		}
	}
	return nil, nil
}

func DeleteContact(id int) bool {
	for i, contact := range contacts {
		if contact.ID == id {
			contacts = append(contacts[:i], contacts[i+1:]...)
			return true
		}
	}
	return false
}