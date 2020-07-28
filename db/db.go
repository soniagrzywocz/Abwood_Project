package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go_server/config"
	"go_server/models"

	"github.com/jmoiron/sqlx"
)

type ServDB struct {
	*sqlx.DB
}

var persistentDb ServDB

func CreateMySQLHandler(mysqlConfig config.MySQL) {
	var err error
	connectString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&time_zone=%s",
		mysqlConfig.Username, mysqlConfig.Password, mysqlConfig.Host,
		mysqlConfig.Database, mysqlConfig.Timezone)
	persistentDb.DB, err = sqlx.Open("mysql", connectString)
	if err != nil {
		panic(err)
	}
}

// function that selects struct info from the table
func GetInfo(name string) (err error) {
	db, err := sql.Open("mysql", "user:password@/dbname")
	defer db.Close()
	if err != nil {
		fmt.Println(err.Error())
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
	}

	var c models.Contact
	err = db.QueryRow("select name, email, message from contacts WHERE name = ?", name).Scan(&c.Name, &c.Email, &c.Message)

	if err != nil {
		fmt.Println(err.Error())
	}

	//fmt.Printf("Name: %s\n Email: %s\n Message: %s\n", c.Name, c.Email, c.Message)

	b, err := json.Marshal(c)

	fmt.Println(b)
	//instance of a contact

	//marshal a JSON-encoded version of m using json.Marshal
	// b, err := json.Marshal(m)

	return err
}

// func main() {

// 	//getInfo()
// }
