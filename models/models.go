package models

import (
	"fmt"
	"github.com/k0k1a/go-gin-example/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

var db *gorm.DB

type Model struct {
	ID         int            `gorm:"primaryKey" json:"id,omitempty"`
	CreatedOn  int            `gorm:"autoCreateTime" json:"created_on,omitempty"`
	ModifiedOn int            `gorm:"autoUpdateTime" json:"modified_on,omitempty"`
	DeletedOn  gorm.DeletedAt `json:"deleted_on,omitempty"`
}

func init() {
	var (
		err                                       error
		dbName, user, password, host, tablePrefix string
	)
	//dbType = sec.Key("TYPE").String()
	dbName = setting.DatabaseSetting.Name
	user = setting.DatabaseSetting.User
	password = setting.DatabaseSetting.Password
	host = setting.DatabaseSetting.Host
	tablePrefix = setting.DatabaseSetting.TablePrefix

	//MYSQL dsn格式： {username}:{password}@tcp({host}:{port})/{Dbname}?charset=utf8&parseTime=True&loc=Local
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, dbName)

	//设置数据库表前缀
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   tablePrefix,
			SingularTable: true,
		},
	})

	if err != nil {
		log.Println(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Println(err)
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
}
