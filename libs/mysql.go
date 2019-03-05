package libs

import (
	"database/sql"
	"fmt"
)

var db *sql.DB

func instance() {
	fmt.Printf(db)
	if !db {

	}
}
