package models

type Group struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Contacts    []int  `json:"contacts"`
}
