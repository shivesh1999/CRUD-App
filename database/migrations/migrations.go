package migrations

import "gorm.io/gorm"

type Contacts struct {
	ID           uint    `gorm: primary key;autoIncrement" json :"id"`
	Name         *string `json: name`
	MobileNumber *string `json : contactNumber`
	City         *string `json : city`
	Country      *string `json : country`
	Email        *string `json : email`
}

func MigrateContacts(db *gorm.DB) error {
	err := db.AutoMigrate(&Contacts{})
	return err
}
