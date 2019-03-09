package models

import (
	"database/sql"
	_ "gin/plugins/mysql"
	"os"
	"web/libs"
)

type Mysql struct {
	instance *sql.DB
}

var this Mysql

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

	this.instance, err = sql.Open("mysql", dns)
	libs.CheckErr(err)

	err_dping := this.instance.Ping()

	libs.CheckErr(err_dping)

	db_idel,err	:= cfg.Key("mysql.max_idle_conns").Int()
	libs.CheckErr(err)

	db_open,err := cfg.Key("mysql.max_open_conns").Int()
	libs.CheckErr(err)

	this.instance.SetMaxIdleConns(db_idel)

	this.instance.SetMaxOpenConns(db_open)
}

func (_ *Mysql) GetInstance() *sql.DB{
	return this.instance
}

func (_ *Mysql) GetRow(querySql string, record map[string]string) error {
	row, err := this.instance.Query(querySql)
	if nil != err {
		return err
	}

	columns, err := row.Columns()
	if nil != err {
		return err
	}

	scanArgs := make([]interface{}, len(columns))
	values   := make([]interface{}, len(columns))

	for j := range values {
		scanArgs[j] = &values[j]
	}

	row.Next()
	err = row.Scan(scanArgs...)
	for i, col := range values {
		if nil != col {
			record[columns[i]] = string(col.([]byte))
		}
	}

	return  nil
}

func (_ *Mysql) GetAll(querySql string, records *[]map[string]string) error {
	rows, err := this.instance.Query(querySql)
	defer rows.Close()

	if nil != err {
		return err
	}

	columns, err := rows.Columns()
	if nil != err {
		return err
	}

	scanArgs 	:= make([]interface{}, len(columns))
	values 		:= make([]interface{}, len(columns))
	record      := make(map[string]string)

	for j := range values {
		scanArgs[j] = &values[j]
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
		*records = append(*records, record)
	}

	return nil;
}
