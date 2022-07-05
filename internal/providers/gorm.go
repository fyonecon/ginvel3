package providers
// GORM文档：https://learnku.com/docs/gorm/v2/advanced_query/9757
// GORM文档：https://www.cnblogs.com/zisefeizhu/category/1747066.html

import (
	"database/sql"
	"fmt"
	"ginvel/internal/reader/framework_toml"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type InitGORM struct {}

var GDB *gorm.DB // 数据库1
var gErr1 error
func (initGORM *InitGORM) InitGorm1() {
	dbConfig := framework_toml.GetMySQLConfig1()
	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&loc=Local&timeout=%s",
		dbConfig["DB_USER"],
		dbConfig["DB_PWD"],
		dbConfig["DB_HOST"],
		dbConfig["DB_PORT"],
		dbConfig["DB_NAME"],
		dbConfig["DB_CHARSET"],
		dbConfig["DB_TIMEOUT"],
	)
	// 连接现有MySQL
	sqlDB, sErr := sql.Open("mysql", dbDSN)
	GDB, gErr1 = gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	if sErr != nil {log.Println("GORM1现有数据库连接失败，GORM1功能将不可用。。。", sErr)}else {log.Println("尝试连接GORM1服务... ")}

	if gErr1 != nil {log.Println("GORM1数据库连接失败，GDB参数不可用，框架运行中断。", gErr1)}else {log.Println("GORM1已连接现有数据库驱动 >>> ")}

	return
}

var GDB2 *gorm.DB // 数据库2
var gErr2 error
func (initGORM *InitGORM) InitGorm2() {
	dbConfig := framework_toml.GetMySQLConfig2()
	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&loc=Local&timeout=%s",
		dbConfig["DB_USER"],
		dbConfig["DB_PWD"],
		dbConfig["DB_HOST"],
		dbConfig["DB_PORT"],
		dbConfig["DB_NAME"],
		dbConfig["DB_CHARSET"],
		dbConfig["DB_TIMEOUT"],
	)
	// 连接现有MySQL
	sqlDB, sErr := sql.Open("mysql", dbDSN)
	GDB2, gErr2 = gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	if sErr != nil {log.Println("GORM2现有数据库连接失败，GORM2功能将不可用。。。", sErr)}else {log.Println("尝试连接GORM2服务... ")}

	if gErr1 != nil {log.Println("GORM2数据库连接失败，GDB2参数不起作用，框架继续运行", gErr1)}else {log.Println("GORM2已连接现有数据库驱动 >>> ")}

	return
}
