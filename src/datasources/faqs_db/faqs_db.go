package faqs_db

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ssoql/faq-chat-bot/src/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var (
	Client *gorm.DB
)

func init() {
	if config.IsProduction() || config.IsDevelop() {
		Client = getDbClient()
	}
}

func getDbClient() (db *gorm.DB) {
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

	return database
}

func MigrateData(model interface{}) {
	if err := Client.AutoMigrate(model); err != nil {
		panic("failed to migrate data into DB!")
	}
}

func GetDbMock() (*gorm.DB, sqlmock.Sqlmock) {
	// get db and mock
	sqlDB, mock, err := sqlmock.New(
		sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp),
	)
	if err != nil {
		log.Fatalf("[sqlmock new] %s", err)
	}
	defer sqlDB.Close()

	// create dialector
	dialector := mysql.New(mysql.Config{
		Conn:       sqlDB,
		DriverName: "mysql",
	})

	// a SELECT VERSION() query will be run when gorm opens the database
	// so we need to expect that here
	columns := []string{"version"}
	mock.ExpectQuery("SELECT VERSION()").WithArgs().WillReturnRows(
		mock.NewRows(columns).FromCSVString("1"),
	)

	// open the database
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Fatalf("[gorm open] %s", err)
	}

	return db, mock
}
