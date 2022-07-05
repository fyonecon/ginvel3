package providers
// mysql文档：http://www.topgoer.com/%E6%95%B0%E6%8D%AE%E5%BA%93%E6%93%8D%E4%BD%9C/go%E6%93%8D%E4%BD%9Cmysql/mysql%E4%BD%BF%E7%94%A8.html

import (
	"database/sql"
	"fmt"
	"ginvel/internal/reader/framework_toml"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"strconv"
	"time"
)

type InitMySQL struct {}

var MysqlDB *sql.DB
var MysqlDBErr error
func (initMySQL *InitMySQL) InitMysql1() {
	log.Println("尝试连接MySQL1服务...")

	// get db config
	var dbConfig  map[string]string = framework_toml.GetMySQLConfig1()

	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&loc=Local&timeout=%s",
		dbConfig["DB_USER"],
		dbConfig["DB_PWD"],
		dbConfig["DB_HOST"],
		dbConfig["DB_PORT"],
		dbConfig["DB_NAME"],
		dbConfig["DB_CHARSET"],
		dbConfig["DB_TIMEOUT"],
	)

	MysqlDB, MysqlDBErr = sql.Open("mysql", dbDSN)

	if MysqlDBErr != nil {
		panic("database1 data source name error: " + MysqlDBErr.Error())
	}

	// max open connections
	dbMaxOpenConns, _ := strconv.Atoi(dbConfig["DB_MAX_OPEN_CONNS"])
	MysqlDB.SetMaxOpenConns(dbMaxOpenConns)

	// max idle connections
	dbMaxIdleConns, _ := strconv.Atoi(dbConfig["DB_MAX_IDLE_CONNS"])
	MysqlDB.SetMaxIdleConns(dbMaxIdleConns)

	// max lifetime of connection if <=0 will forever
	dbMaxLifetimeConns, _ := strconv.Atoi(dbConfig["DB_MAX_LIFETIME_CONNS"])
	MysqlDB.SetConnMaxLifetime(time.Duration(dbMaxLifetimeConns))

	if MysqlDBErr = MysqlDB.Ping(); nil != MysqlDBErr {
		log.Println("MySQL1数据库连接失败，MysqlDB2参数不可用，框架运行中断。", MysqlDBErr.Error())
		os.Exit(200)
	}else {
		log.Println("MySQL1已连接 >>> ")
	}

	return
}


var MysqlDB2 *sql.DB // 数据库2
var MysqlDBErr2 error
func (initMySQL *InitMySQL) InitMysql2() {
	log.Println("尝试连接MySQL2服务...")

	// get db config
	var dbConfig  map[string]string = framework_toml.GetMySQLConfig2()

	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&loc=Local&timeout=%s",
		dbConfig["DB_USER"],
		dbConfig["DB_PWD"],
		dbConfig["DB_HOST"],
		dbConfig["DB_PORT"],
		dbConfig["DB_NAME"],
		dbConfig["DB_CHARSET"],
		dbConfig["DB_TIMEOUT"],
	)

	MysqlDB2, MysqlDBErr2 = sql.Open("mysql", dbDSN)

	if MysqlDBErr2 != nil {
		panic("database2 data source name error: " + MysqlDBErr.Error())
	}

	// max open connections
	dbMaxOpenConns, _ := strconv.Atoi(dbConfig["DB_MAX_OPEN_CONNS"])
	MysqlDB.SetMaxOpenConns(dbMaxOpenConns)

	// max idle connections
	dbMaxIdleConns, _ := strconv.Atoi(dbConfig["DB_MAX_IDLE_CONNS"])
	MysqlDB.SetMaxIdleConns(dbMaxIdleConns)

	// max lifetime of connection if <=0 will forever
	dbMaxLifetimeConns, _ := strconv.Atoi(dbConfig["DB_MAX_LIFETIME_CONNS"])
	MysqlDB.SetConnMaxLifetime(time.Duration(dbMaxLifetimeConns))

	if MysqlDBErr2 = MysqlDB2.Ping(); nil != MysqlDBErr2 {
		log.Println("MySQL2数据库连接失败，MysqlDB2参数不起作用，框架继续运行", MysqlDBErr2.Error())
		//os.Exit(200)
	}else {
		log.Println("MySQL2已连接 >>> ")
	}

	return
}
