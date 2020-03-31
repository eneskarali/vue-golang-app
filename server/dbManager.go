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
	user.Close()
	db.Close()

	userCreds := make(map[string]string)
	userCreds["username"] = username
	userCreds["password"] = password
	userCreds["name"] = name
	userCreds["surname"] = surname
	userCreds["postCount"] = strconv.Itoa(postCount)

	return userCreds

}

func createPost(postBy string, postText string, date int64) {

	db, err := sql.Open("sqlite3", "database")
	checkErr(err)

	stat, err := db.Prepare("INSERT INTO post(postBy, postText, date) values(?,?,?)")
	checkErr(err)

	res, err := stat.Exec(postBy, postText, date)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Print(id)

	stat.Close()
	db.Close()

}

func getPost(from int, to int) []postInfo {

	db, err := sql.Open("sqlite3", "database")
	checkErr(err)

	post, err := db.Query("select postBy, postText, date from post order by date desc limit " + strconv.Itoa(to) + " offset " + strconv.Itoa(from) + "")
	checkErr(err)

	var username string
	var name string
	var surname string
	var postCount int
	var postBy string
	var postText string
	var date int

	var posts []postInfo

	for post.Next() {
		err = post.Scan(&postBy, &postText, &date)

		user, err := db.Query("SELECT username, name, surname, postCount FROM user")
		checkErr(err)

		for user.Next() {
			err = user.Scan(&username, &name, &surname, &postCount)
			if username == postBy {
				break
			}
		}
		user.Close()

		posts = append(posts, postInfo{name, surname, postText, date, postCount})
	}
	post.Close()
	db.Close()

	return posts
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
