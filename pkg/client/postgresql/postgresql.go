package postgresql

import (
	"github.com/jackc/pgx"
)

type Client interface {
	Exec(sql string, arguments ...interface{}) (commandTag pgx.CommandTag, err error)
	Query(sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(sql string, args ...interface{}) pgx.Row
	Begin() (pgx.Tx, error)
}

/*func NewClient() () {

}*/
