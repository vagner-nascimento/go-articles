package monkeypatching

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var SqlOpen = sql.Open

func OpenDd(user, pass, addr, db string) (*sql.DB, error) {
	connStr := fmt.Sprintf("%s:%s@%s/%s", user, pass, addr, db)

	return SqlOpen("mysql", connStr)
}
