package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var db = new(sql.DB)

func InitDb() *sql.DB {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}

	return db
}

type Tag struct {
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (s Tag) insert() (sql.Result, error) {
	stmt, err := db.Prepare("insert into blog_tag(name,state) values(?,?)")
	if err != nil {
		return nil, err
	}
	res, err := stmt.Exec(s.Name, s.State)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func main() {
	db = InitDb()
	defer db.Close()
	tag := Tag{Name: "zhangsan", State: 2}
	result, err := tag.insert()
	if err != nil {
		fmt.Println(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("操作数据的id：", id)

}
