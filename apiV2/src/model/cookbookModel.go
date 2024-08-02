package model

import "database/sql"

// Cookbook omitting audit columns
type Cookbook struct {
	Cookbook_id int            `json:"cookbook_id"`
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
}

// Convenience method used with sql#Row.Scan
func (c *Cookbook) PtrFields() []any {
	return []any{&c.Cookbook_id, &c.Name, &c.Description}
}
