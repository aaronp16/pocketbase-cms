package models

var _ Model = (*CMSConfig)(nil)

type CMSConfig struct {
	BaseModel

	PrimaryColour   	string `db:"primaryColour" json:"primaryColour"`
	SecondaryColour 	string `db:"secondaryColour" json:"secondaryColour"`
	Logo   				int `db:"Logo" json:"Logo"`
}

func (m *CMSConfig) TableName() string {
	return "_cms_configs"
}