package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	id       int
	username string
}

func main() {
	db, err := sql.Open("mysql", "root:Hx123456789@tcp(127.0.0.1:3306)/safety")
	if err != nil {
		fmt.Printf(" Format failed, Error: %v \n", err)
		panic(err)
	}
	defer db.Close()

	var id, username string
	rows := db.QueryRow("select id, username from USER where id = 1") //获取一行数据
	rows.Scan(&id, &username)                                         //将rows中的数据存到id,name中
	fmt.Println(id, "--", username)
}
