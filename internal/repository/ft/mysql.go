package ft

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init(dsn string) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	fmt.Println(db, err)
}
