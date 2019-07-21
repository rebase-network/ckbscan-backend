package model

import (
	"ckbscan-backend/conf"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Init() {

	dbCfg := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		conf.Conf.PostgreSQL.Username,
		conf.Conf.PostgreSQL.Password,
		conf.Conf.PostgreSQL.Host,
		conf.Conf.PostgreSQL.Port,
		true,
		"Local")

	db, err := gorm.Open("postgres", dbCfg)
	if err != nil {
		panic(err)
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(0)
}
