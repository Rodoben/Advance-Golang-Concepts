package main

import (
	"fmt"
	"log"
)

type User struct {
	ID    int
	First string
	Last  string
}
type MockDataStore struct {
	users map[int]User
}

func (md MockDataStore) GetUser(id int) (User, error) {
	users, ok := md.users[id]
	if !ok {
		return User{}, fmt.Errorf("No Records found with id %v", id)
	}
	return users, nil
}

func (md MockDataStore) SaveUser(u User) error {
	_, ok := md.users[u.ID]
	if ok {
		return fmt.Errorf("User %v already exists", u.ID)
	}
	md.users[u.ID] = u
	return nil
}

type DataStore interface {
	GetUser(id int) (User, error)
	SaveUser(u User) error
}

type Service struct {
	ds DataStore
}

func (s Service) GetUser(id int) (User, error) {
	return s.ds.GetUser(id)
}

func (s Service) SaveUser(u User) error {
	return s.ds.SaveUser(u)
}

func main() {
	db := MockDataStore{
		users: make(map[int]User),
	}

	srvc := Service{
		ds: db,
	}

	u1 := User{
		ID:    1,
		First: "James",
		Last:  "Benjamin",
	}

	err := srvc.SaveUser(u1)
	if err != nil {
		log.Fatalf("error %s", err)
	}

	u1Returned, err := srvc.GetUser(u1.ID)
	if err != nil {
		log.Fatalf("error %s", err)
	}

	fmt.Println(u1)
	fmt.Println(u1Returned)

}
