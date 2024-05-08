package models

import "time"

// Bundles represents the bundles table.
type (
	Bundles struct {
		BundleID  uint      `gorm:"primaryKey;autoIncrement" json:"bundle_id"` // Primary key
		Name      string    `gorm:"type:float" json:"btc_usd_price"`
		UpdatedAt time.Time `gorm:"" json:"updated_at"`
		CreatedAt time.Time `gorm:"" json:"created_at"`
	}
)

// TableName overrides the table name used by GORM to `bundles`
func (Bundles) TableName() string {
	return "bundles"
}
