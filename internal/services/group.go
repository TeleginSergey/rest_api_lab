package services

import "cmd/main.go/internal/models"

var groups = []models.Group{
	{
		ID:          1,
		Title:       "Friends",
		Description: "Close friends group",
		Contacts:    []int{1, 2},
	},
}

func GetGroups() []models.Group {
	return groups
}

func CreateGroup(newGroup models.Group) {
	groups = append(groups, newGroup)
}

func UpdateGroup(id int, key string, value string) (*models.Group, error) {
	for i, group := range groups {
		if group.ID == id {
			switch key {
			case "title":
				groups[i].Title = value
			case "description":
				groups[i].Description = value
			default:
				return nil, nil
			}
			return &groups[i], nil
		}
	}
	return nil, nil
}

func DeleteGroup(id int) bool {
	for i, group := range groups {
		if group.ID == id {
			groups = append(groups[:i], groups[i+1:]...)
			return true
		}
	}
	return false
}