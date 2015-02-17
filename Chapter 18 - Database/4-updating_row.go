package main

import (
  "database/sql"
  "fmt"
  _ "github.com/bmizerany/pq"
)

func main() {
  db, err := sql.Open("postgres", "user=go_app password=password dbname=go_example sslmode=disable")
  checkErr(err)

  stmt, err := db.Prepare("INSERT INTO users(username,departname) VALUES($1,$2) RETURNING uid")
  checkErr(err)

  res, err := stmt.Exec("RKReloaded", "GoLang")
  checkErr(err)

  stmt, err = db.Prepare("update users set username=$1 where uid=$9")
  checkErr(err)

  res, err = stmt.Exec("update", 3)
  checkErr(err)

  affected, err := res.RowsAffected()
  checkErr(err)

  fmt.Println(affected)

}

func checkErr(err error) {
  if err != nil {
    panic(err)
  }
}