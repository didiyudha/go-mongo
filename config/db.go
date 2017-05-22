package config

import (
	"fmt"

	_ "github.com/lib/pq"
	"gopkg.in/mgo.v2"
)

// DB database
var DB *mgo.Database

// Books mongo collection
var Books *mgo.Collection

func init() {
	s, err := mgo.Dial("mongodb://localhost/bookstore")
	if err != nil {
		panic(err)
	}
	if err = s.Ping(); err != nil {
		panic(err)
	}
	DB = s.DB("bookstore")
	Books = DB.C("books")
	fmt.Println("You are connected to mongodb")
}
