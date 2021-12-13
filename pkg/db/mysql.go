package db

import (
	"fmt"
	"gorm.io/driver/mysql"

	"gorm.io/gorm"
	"time"
)

var Gorm *gorm.DB

// InitGorm 初始化mysqlOrm
func InitGorm(username string, password string, host string, dbName string, maxIdle int, maxOpen int) error {
	connInfo := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username,
		password,
		host,
		dbName)
	var err error
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       connInfo, // DSN data source name
		DefaultStringSize:         256,      // string 类型字段的默认长度
		DisableDatetimePrecision:  true,     // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,     // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,     // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,    // 根据版本自动配置, &gorm.Config{})
	}), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	// 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(maxIdle)
	// 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(maxOpen)
	// 设置了连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Hour)
	initDb(db)
	Gorm = db
	return nil
}

func initDb(db *gorm.DB) {
	//var err error
	//// 创建表时添加后缀
	//err = db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&po.User{})
	//if err != nil {
	//	log.Error("sys", "迁移数据库失败")
	//}
}
