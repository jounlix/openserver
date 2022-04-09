package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type list struct{
	id int `json:"id"`
	info string `json:"id"`
}

func api(c *gin.Context) {
	db, err := sql.Open("mysql", "root:root@/test")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Exec("insert into test.list (info) values (?)", "JOOOOON SINAAAA!!!")

	c.JSON(200, gin.H {
		"message": "",
	})
}

func all(c *gin.Context) {
	db, err := sql.Open("mysql", "root:root@/test")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from test.list")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	for rows.Next(){
		l := list{}
		rows.Scan(&l.id, &l.info)
		lists = append(lists, l)
	}

	c.JSON(200, lists)
}

func main() {
	r := gin.Default()
	r.GET("/api", api)
	r.GET("/api", all)
	r.Run("0.0.0.0:9090")
}