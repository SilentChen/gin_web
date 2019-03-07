package models

import (
	"database/sql"
	_ "gin/plugins/mysql"
	"os"
	"web/libs"
)

type Mysql struct {

}

var DBInstance *sql.DB

func init() {

	webRootPath, err := os.Getwd()
	libs.CheckErr(err)

	cfg := libs.LoadIniFile(webRootPath + "/conf/static.ini")

	db_user := cfg.Key("mysql.user").String()
	db_pwd  := cfg.Key("mysql.passwd").String()
	db_host := cfg.Key("mysql.host").String()
	db_port := cfg.Key("mysql.port").String()
	db_name := cfg.Key("mysql.name").String()
	db_char := cfg.Key("mysql.charset").String()

	dns := db_user + ":" + db_pwd + "@tcp(" + db_host + ":" + db_port + ")/" + db_name + "?charset=" + db_char + "&loc=Asia%2FShanghai"

	DBInstance, err_dopen := sql.Open("mysql", dns)

	libs.CheckErr(err_dopen)

	err_dping := DBInstance.Ping()

	libs.CheckErr(err_dping)

	db_idel,err	:= cfg.Key("mysql.max_idle_conns").Int()
	libs.CheckErr(err)

	db_open,err := cfg.Key("mysql.max_open_conns").Int()
	libs.CheckErr(err)

	DBInstance.SetMaxIdleConns(db_idel)

	DBInstance.SetMaxOpenConns(db_open)
}

func (_ *Mysql) test() {
	
}