package main

import (
	"fmt"

	architecture "github.com/cmd-ctrl-q/golang-general-architecture"
	"github.com/cmd-ctrl-q/golang-general-architecture/storage/mongo"
)

func main() {

	// create database
	mongdb := mongo.DB{}
	// psqldb := postgres.Psql{}

	// create objects
	p1 := architecture.Person{
		First: "Alice",
	}
	p2 := architecture.Person{
		First: "Bob",
	}
	ps := architecture.NewPersonService(mongdb)

	// store in mongo db
	ps.Put(1, p1)
	ps.Put(2, p2)

	fmt.Println(ps.Get(1)) // alice
	fmt.Println(ps.Get(2)) // Bob
	fmt.Println(ps.Get(3)) // nil
}
