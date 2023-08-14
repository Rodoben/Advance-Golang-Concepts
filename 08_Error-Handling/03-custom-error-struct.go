package main

import (
	"fmt"
	"log"
)

func main() {

	c := customErr{}
	foo(&c)

	c1 := customErr1{
		info: "Success",
		m: map[int]string{
			200: "Succesful transaction 200",
			201: "Succesful transaction 201",
			202: "Succesful transaction 202",
			203: "Succesful transaction203",
		},
	}
	bar(&c1)

	_, c2 := sqrt(-10.8)
	if c2 != nil {
		log.Println(c2)
	}

}

// Error interface Error() string
//------------------------------------------------------------------------
type customErr struct {
}

func (c *customErr) Error() string {
	return "Error 404:"
}

func foo(err error) {
	log.Println("I am printing error here", err)
}

// type Assertion
//------------------------------------------------------------------------
type customErr1 struct {
	info string
	m    map[int]string
}

func (c *customErr1) Error() string {
	return fmt.Sprintf("I hold info: %s and a map %v", c.info, c.m)
}

func bar(err error) {
	fmt.Println("I am Bar", err.(*customErr1).info)
	fmt.Println("I am Bar", err.(*customErr1).m)

	for i, v := range err.(*customErr1).m {
		fmt.Println(i, v)
	}
}

// Struct an Error
//---------------------------------------------------------------------------

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (s *sqrtError) Error() string {

	//fmt.Println(s)
	return fmt.Sprintf("Lat: %v ---Long: %v----Error: %v", s.lat, s.long, s.err)
}

func sqrt(f float32) (float32, error) {
	if f < 0 {
		// e := errors.New("more coffee needed")
		e := fmt.Errorf("more coffee needed - value was %v", f)
		return 0, &sqrtError{"50.2289 N", "99.4656 W", e}
	}
	return 42, nil
}
