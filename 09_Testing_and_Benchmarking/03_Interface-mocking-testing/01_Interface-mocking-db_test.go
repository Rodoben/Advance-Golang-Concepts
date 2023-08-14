package main

import (
	"testing"
)

func TestGetUser(t *testing.T) {
	md := &MockDataStore{
		users: map[int]User{
			2: {ID: 2, First: "Jenny", Last: "Ronald"},
		},
	}

	s := &Service{
		ds: md,
	}

	u, err := s.GetUser(2)
	if err != nil {
		t.Errorf("got error: %v", err)
	}
	if u.First != "Jenny" {
		t.Errorf("got error:  got  %v   want %v", u.First, "Jenny")
	}

}

func TestSaveUser(t *testing.T) {
	md := &MockDataStore{
		users: map[int]User{},
	}

	u1 := User{
		ID:    1,
		First: "Ronald",
		Last:  "Benjamin",
	}

	s := &Service{
		ds: md,
	}

	err := s.SaveUser(u1)
	if err != nil {
		t.Errorf("Got Error: unable to insert User %v", u1)
	}

	if md.users[1].First != "Ronald" {
		t.Errorf("got error:  got  %v   want %v", md.users[1].First, "Ronald")
	}
	if md.users[1].Last != "Benjamin" {
		t.Errorf("got error:  got  %v   want %v", md.users[1].First, "Benjamin")
	}
}
