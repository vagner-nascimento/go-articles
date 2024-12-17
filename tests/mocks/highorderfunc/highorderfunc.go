package highorderfunc

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type SqlOpener func(string, string) (*sql.DB, error)

func OpenDd(user, pass, addr, db string, open SqlOpener) (*sql.DB, error) {
	connStr := fmt.Sprintf("%s:%s@%s/%s", user, pass, addr, db)

	return open("mysql", connStr)
}
