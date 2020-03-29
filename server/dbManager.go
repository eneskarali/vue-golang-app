package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func getUser(uname string) map[string]string {
	fmt.Printf("")

	db, err := sql.Open("sqlite3", "database")
	checkErr(err)

	user, err := db.Query("SELECT * FROM user")
	checkErr(err)

	var username string
	var password string
	var name string
	var surname string
	var postCount int

	for user.Next() {
		err = user.Scan(&username, &password, &name, &surname, &postCount)
		if username == uname {
			break
		}
	}

	db.Close()

	userCreds := make(map[string]string)
	userCreds["username"] = username
	userCreds["password"] = password
	userCreds["name"] = name
	userCreds["surname"] = surname
	userCreds["postCount"] = strconv.Itoa(postCount)

	return userCreds

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
