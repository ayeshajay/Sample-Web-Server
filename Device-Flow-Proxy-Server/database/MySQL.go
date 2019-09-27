package database

import (
"database/sql"
_ "github.com/go-sql-driver/mysql"
"log"
)

type Tag struct {
	CLIENT_ID   int    `json:"client_id"`
	USER_CODE string `json:"user_code"`
}

func main(){

	db, err := sql.Open("mysql", "user1:user1@(127.0.0.1:3306)/dbname")

	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	// Execute the query
	results, err := db.Query("SELECT id, name FROM example")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		var tag Tag
		// for each row, scan the result into our tag composite object
		err = results.Scan(&tag.CLIENT_ID, &tag.USER_CODE)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		log.Printf(tag.USER_CODE)
	}

}

