package db

import (
	"fmt"
	"go_server/config"

	"github.com/jmoiron/sqlx"
)

// ServDB serves as a small wrapper around sqlx.DB so in the future we can add additional functions
type ServDB struct {
	*sqlx.DB
}

var persistentDb ServDB

// CreateMySQLHandler uses the configuration values to open a persistent database connection to MySQL
func CreateMySQLHandler(mysqlConfig config.MySQL) {
	var err error
	connectString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		mysqlConfig.Username, mysqlConfig.Password, mysqlConfig.Host,
		mysqlConfig.Database)
	persistentDb.DB, err = sqlx.Open("mysql", connectString)
	if err != nil {
		panic(err)
	}
}

// Db returns the handle to the configured persistent database connection
func Db() ServDB {
	return persistentDb
}

// func main() {

// 	//getInfo()
// }
