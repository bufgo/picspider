package setting

import (
	"log"

	"github.com/go-ini/ini"
)

// Database model
type Database struct {
	Type     string
	User     string
	Password string
	Host     string
	Name     string
}

// DatabaseSetting database info
var DatabaseSetting = &Database{}
var cfg *ini.File

// Setup initialize the configuration instance
func Setup() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("database", DatabaseSetting)
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.mapTo %s err: %v", section, err)
	}
}
