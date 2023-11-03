package initial

import (
	"gorm.io/gorm"

	"gorm.io/driver/postgres"
)

var PG_CONN = "postgresql://yomi:yomi123@postgres:5432/main?sslmode=disable"

var CrudDB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(postgres.Open(PG_CONN), &gorm.Config{TranslateError: true})

	if err != nil {
		return
	}

	CrudDB = db
}
