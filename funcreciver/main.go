package main

import "fmt"

type Database struct {
	user string
	id   int
}

type server struct {
	db *Database
}

func (s *server) GetUser() string {
	user := s.db.user
	return user
}

func main() {
	server := &server{
		db: &Database{
			user: "Hemlo",
			id:   1,
		},
	}
	// if using func receiver then always have to put pointed when declaring the new struct
	user := server.GetUser()
	fmt.Println(user)
}
