package models

import (
	"fmt"
	"log"

	// import database driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/picspider/pkg/setting"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

var engine *xorm.Engine

// Setup initializes the databse instance
func Setup() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name)
	engine, err = xorm.NewEngine(setting.DatabaseSetting.Type, dsn)

	if err != nil {
		log.Fatalf("models.Setup, database conn err: %v", err)
	}

	engine.SetMapper(names.GonicMapper{})

	err = engine.Sync2(new(PhotoAlbum), new(Photo))
	if err != nil {
		log.Fatalf("models.Sycn2, database table create err: %v", err)
	}
}
