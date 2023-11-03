package initial

import (
	"gorm.io/gorm"

	"gorm.io/driver/postgres"
)

var PG_CONN = "postgresql://yomi:yomi123@localhost:5432/main"

var CrudDB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(postgres.Open(PG_CONN), &gorm.Config{})

	if err != nil {
		return
	}

	CrudDB = db
}
