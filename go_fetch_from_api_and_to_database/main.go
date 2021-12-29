package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ilibs/gosql/v2"
)

type Nodee struct {
	UserId    int    `json:"userId" db:"userId"`
	Id        int    `json:"id" db:"id"`
	Title     string `json:"title" db:"title"`
	Completed bool   `json:"completed" db:"completed"`
}

func (u *Nodee) TableName() string {
	return "todolist"
}

func (u *Nodee) PK() string {
	return "id"
}

func main() {

	configs := make(map[string]*gosql.Config)

	configs["default"] = &gosql.Config{
		Enable:  true,
		Driver:  "mysql",
		Dsn:     "root:lolnoyou@tcp(127.0.0.1:3306)/gotestland",
		ShowSql: true,
	}

	//connection database
	gosql.Connect(configs)

	// fetch data
	data, err := http.Get("https://jsonplaceholder.typicode.com/todos")
	if err != nil {
		panic(err)
	}

	// response handling
	response, err := io.ReadAll(data.Body)
	if err != nil {
		panic(err)
	}
	var SerializedData []Nodee

	// unserilization
	if err := json.Unmarshal(response, &SerializedData); err != nil {
		panic(err)
	}

	for _, todo := range SerializedData {
		fmt.Println(todo)
		gosql.Model(&todo).Create()

	}

	// for i := 0; i < len(SerializedData); i++ {
	// 	gosql.Model(&Todolist{
	// 		Id:        SerializedData[i].Id,
	// 		UserId:    SerializedData[i].UserId,
	// 		Title:     SerializedData[i].Title,
	// 		Completed: SerializedData[i].Completed}).Create()

	// } // 5

	fmt.Println("all")
}
