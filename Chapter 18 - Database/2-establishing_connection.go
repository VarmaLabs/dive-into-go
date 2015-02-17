package main

import (
  "database/sql"
  "fmt"
  _ "github.com/bmizerany/pq"
)

func main() {
  db, err := sql.Open("postgres", "user=go_app password=password dbname=go_example sslmode=disable")
  checkErr(err)
}

func checkErr(err error) {
  if err != nil {
    panic(err)
  }
}