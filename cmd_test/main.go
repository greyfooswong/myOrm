package main

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"myOrm"
)

func main() {
	engine, _ := myOrm.NewEngine("sqlite3", "gee.db")
	defer engine.Close()
	s := engine.NewSession()
	_, _ = s.Raw("drop table if exists User;").Exec()
	_, _ = s.Raw("create table User(Name text);").Exec()
	_, _ = s.Raw("create table User(Name text);").Exec()
	result, _ := s.Raw("insert into User(`Name`) values (?),(?)", "Tom", "Sam").Exec()
	count, _ := result.RowsAffected()
	fmt.Printf("Exec success, %d affected \n", count)
}
