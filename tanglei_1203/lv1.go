package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)
type DbWorker struct {
	//mysql data source name
	Dsn string
}
type User struct {
	ID int64 `db:"ID"`
	Name string  `db:"name"`
	//由于在mysql的users表中name没有设置为NOT NULL,所以name可能为null,在查询过程中会返回nil，如果是string类型则无法接收nil,但sql.NullString则可以接收nil值
	sex int `db:"sex"`
	password string `db:"password"`
}
var db *sql.DB

func initDB()(err error) {
	dbw := DbWorker{
		Dsn: "mysql8:TangLei123@tcp(127.0.0.1:3306)/student_mysql",
	}
	db, err = sql.Open("mysql",
		dbw.Dsn)
	if err != nil {
		panic(err)
		return
	}
	//defer db.Close()
	return nil
}

func main() {
	err := initDB() // 调用输出化数据库的函数
	if err != nil {
		fmt.Printf("err")
		return
	}
	r :=insertDB(db,1,2)
	InsertDB(db,"qwe")
	println(r)
}

func insertDB(db *sql.DB,a,b int) (result interface{})  {
	result, _ = db.Exec("insert test (ID,sex) values (?,?)", a, b)
	return result
}

func InsertDB(db *sql.DB,value string)  {
	stmt, err := db.Prepare("insert into test (sex) values (?)")
	if err != nil{
		log.Fatal(err)
	}
	stmt.Exec(value)

}

