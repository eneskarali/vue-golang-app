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

func getPost(from int, to int) [10]map[string]string {

	db, err := sql.Open("sqlite3", "database")
	checkErr(err)

	post, err := db.Query("select * from post order by date desc limit " + strconv.Itoa(to) + " offset " + strconv.Itoa(from) + "")
	checkErr(err)

	var username string
	var password string
	var postCount int
	var postId int
	var postBy string
	var postText string
	var date int
	var name string
	var surname string

	var posts [10]map[string]string
	var i int = 0

	for post.Next() {
		err = post.Scan(&postId, &postBy, &postText, &date)

		user, err := db.Query("SELECT * FROM user")
		checkErr(err)

		for user.Next() {
			err = user.Scan(&username, &password, &name, &surname, &postCount)
			if username == postBy {
				break
			}
		}
		user.Close()

		postCreds := make(map[string]string)
		postCreds["name"] = name
		postCreds["surname"] = surname
		postCreds["postText"] = postText
		postCreds["date"] = strconv.Itoa(date)
		postCreds["postCount"] = strconv.Itoa(postCount)

		posts[i] = postCreds
		i += 1
		if i == 10 {
			break
		}
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
