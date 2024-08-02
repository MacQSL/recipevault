package model

// Cookbook omitting audit columns
type Cookbook struct {
	Cookbook_id int     `json:"cookbook_id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}