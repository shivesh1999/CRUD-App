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

type Users struct {
	ID       uint    `gorm: primary key;autoIncrement" json :"id"`
	Name     *string `json: name`
	Email    *string `json : email`
	Password *string `json :password`
}

func MigrateContacts(db *gorm.DB) error {
	err := db.AutoMigrate(&Contacts{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&Users{})
	return err
}
