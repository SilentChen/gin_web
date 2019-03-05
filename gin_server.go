package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"gin"
	_ "gin/plugins/mysql"
	"io/ioutil"
	"os"
	"web/controllers/admin"
)

type static_mysql struct {
	Host string `json:"host"`
	User string `json:"user"`
	Passwd string `json:"passwd"`
	DbName string `json:"dbname""`
	Write bool `json:"write"`
}
type static struct {
	App string `json:"base_app"`
	Debug bool `json:"base_debug"`
	Key string `json:"base_key"`
	Mysql map[string][]static_mysql `json:"db_mysql"`
}
var cfg static

func db_intit() (*sql.DB, error) {
	return sql.Open("mysql", "root:qwer@tcp(192.168.33.200:3306)/qz_center?charset=utf8")
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func pool(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM users limit 1")
	defer rows.Close()
	checkErr(err)

	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for j := range values {
		scanArgs[j] = &values[j]
	}

	record := make(map[string]string)
	for rows.Next() {
		//将行数据保存到record字典
		err = rows.Scan(scanArgs...)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
	}

	fmt.Println(record)
}

func load_conf() {
	fd, err := os.Open("conf/static.json")
	if nil != err {
		return
	}
	staticConf, err := ioutil.ReadAll(fd)
	checkErr(err)

	if err := json.Unmarshal([]byte(staticConf), &cfg); err == nil {
		fmt.Println(cfg.Mysql["master"][0].Host)
	} else {
		fmt.Println(err)
	}
	fmt.Println(staticConf)
}

func main() {
	load_conf()

	db, err := db_intit()
	if nil != err {
		return
	}

	pool(db)

	fmt.Println(os.Getenv("GOPATH"))
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/admin", admin.Index)
	r.Run() // listen and serve on 0.0.0.0:8080
}