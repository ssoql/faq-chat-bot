package faqs_db

import (
	"fmt"
	"github.com/ssoql/faq-chat-bot/src/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Client *gorm.DB
)

func init() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.GetDbUser(),
		config.GetDbPassword(),
		config.GetDbHost(),
		config.GetDbPort(),
		config.GetDbSchema(),
	)
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect to database!")
	}

	Client = database
}

func MigrateData(model interface{}) {
	if err := Client.AutoMigrate(model); err != nil {
		panic("failed to migrate data into DB!")
	}
}
