package main

import (
	"database/sql"
	"fmt"
)

func main() {
	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	//Execute the query
	rows, err := db.Query("SELECT * FROM table")
	if err != nil {
		panic(err.Error())
	}

	// get columns name
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}

	// make slice for values
	values := make([]sql.RawBytes, len(columns))

	//copying references into slice for rows.Scan
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	//fetch rows
	for rows.Next() {
		// get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}

		// print each column as string
		var value string
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(columns[i], ":", value)
		}

	}
	if err = rows.Err(); err != nil {
		panic(err.Error())
	}
}
