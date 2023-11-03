package migrate

import (
	"golang/crud-basic/initial"
	"golang/crud-basic/model"
)

func Migrate() {
	initial.CrudDB.AutoMigrate(&model.Book{})
}
