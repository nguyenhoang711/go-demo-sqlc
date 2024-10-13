package initialize

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/hoangnguyen/demo-sqlc/global"
	"go.uber.org/zap"
)

func checkErrorPanicc(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

func InitMysqlc() {
	m := global.Config.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", m.Username, m.Password, m.Host, m.Port, m.Dbname)
	db, err := sql.Open("mysql", dsn)
	checkErrorPanicc(err, "InitMysql initialization error")
	global.Logger.Info("Initializing MySQL successfully")
	global.Mdbc = db
}

func SetPoolc() {
	//mo san ket noi --> user su dung thi lay
	m := global.Config.Mysql
	sqlDb, err := global.Mdb.DB()
	if err != nil {
		fmt.Printf("mysql error %v", err)
	}
	sqlDb.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns))
	sqlDb.SetMaxOpenConns(m.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifeTime))
}