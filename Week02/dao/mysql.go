package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

const (
	USER_NAME = "root"
	PASS_WORD = "cpicocr2020"
	HOST      = "10.53.4.204"
	PORT      = "3306"
	DATABASE  = "cpic_ocr"
	CHARSET   = "utf8"
)

var (
	MysqlDB    *sql.DB
	MysqlDBErr error
)

func init() {
	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", USER_NAME, PASS_WORD, HOST, PORT, DATABASE, CHARSET)

	// 打开连接失败
	MysqlDB, MysqlDBErr = sql.Open("mysql", dbDSN)
	//defer MysqlDb.Close();
	if MysqlDBErr != nil {
		log.Println("dbDSN: " + dbDSN)
		panic("数据源配置不正确: " + MysqlDBErr.Error())
	}

	// 最大连接数
	MysqlDB.SetMaxOpenConns(100)
	// 闲置连接数
	MysqlDB.SetMaxIdleConns(20)
	// 最大连接周期
	MysqlDB.SetConnMaxLifetime(100 * time.Second)

	if MysqlDBErr = MysqlDB.Ping(); nil != MysqlDBErr {
		panic("数据库链接失败: " + MysqlDBErr.Error())
	}
}
