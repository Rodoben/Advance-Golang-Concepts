package main

import "fmt"

type Person struct {
	first string
}

func main() {

	i

	//-----------------------------------------------------------------
	p := Person{
		first: "Jenny",
	}

	fmt.Println(p)

	changeNameStruct(p)
	fmt.Println(p)

	changeNamePtrStruct(&p)
	fmt.Println(p)

	p = changeNameReturnStruct(p)
	fmt.Println(p)
	//-------------------------------------------------------------------

}

func changeInt(i int, f float64, b bool, s string) {
	i = 1
	f = 10.23
	b = true
	s = "string1"
}

func changeIntReturn(i int) (int, float64, bool, string) {
	i = 1
	f := 10.23
	b := true
	s := "string1"
	return i, f, b, s
}

func changeIntptr(i *int) {
	i1 := 3
	i = &i1
}

// ----------------- struct
func changeNameStruct(p Person) {
	p.first = "changed"
}
func changeNameReturnStruct(p Person) Person {
	p.first = "changedReturn"
	return p
}
func changeNamePtrStruct(p *Person) {
	p.first = "ChangedPtr"
}

// -------------------slice
func changeSlice(x []int) {
	x[0] = 99
}
func changeSliceReturn(x []int) []int {
	x[0] = 99
	return x
}
func changeSlicePtr(x []*int) { // no need og pointer here since slice is already pointing to the backing array.but we can put the  * after the array []*int , this is array of pointer to int
	x1 := 99
	x[0] = &x1
}

//---------------------map

func changeMap(m map[string]int) {
	m["one"] = 1
}
func changeMapReturn(m map[string]int) map[string]int {
	m["one"] = 1
	return m
}
func changeMapPtr(m *map[string]int) {
	m1 := 1200023
	fmt.Printf("I am under Map pointer %v and Type %T", m1)
	(*m)["one"] = m1

}
func changeMapPtr1(m map[string]*int) {
	m1 := 1200023
	fmt.Printf("I am under Map pointer %v and Type %T", m1)
	m["one"] = &m1

}

//--------------------array

func changeArray(n [5]int) {
	n[1] = 100
}
func changeArrayReturn(n [5]int) [5]int {
	n[1] = 100
	return n
}
func changeArrayPtr(n *[5]int) {
	(*n)[1] = 23
}
func changeArrayPtr1(n [5]*int) {
	n1 := 1001
	n[1] = &n1
}

//------------------channels

func changeChannelPtr(c <-chan *int) {

}
func changeChannelReturn(c <-chan *int) <-chan *int {
	c1 := make(chan *int)

	go func() {
		for i := 0; i < 10; i++ {
			c1 <- &i
		}
		close(c1)
	}()

	return c1
}
func changeChannelReturnChanSlice(c <-chan *[]int) <-chan *[]int {
	c1 := make(chan *[]int)

	go func() {
		c1 <- &[]int{1, 2, 3}
		close(c1)
	}()

	return c1
}

func changeChannelReturnSliceChan(c []chan *int) []chan *int {
	c1 := make([]chan *int, 2)

	go func() {
		c1[0] = make(chan *int)
		n1 := 12
		n2 := 24
		c1[0] <- new(int)
		c1[0] <- new(int)
		c1[0] <- &n1

		c1[1] = make(chan *int)
		c1[1] <- new(int)
		c1[1] <- new(int)
		c1[1] <- &n2
		close(c1[0])
		close(c1[1])

	}()

	return c1
}

// func changeChannelReturnSlice(c <-chan []chan *int) <-chan []chan *int {
// 	c1 := make(chan []chan *int, 1)

// 	go func() {
// 		c1 <- make([]chan *int, 2)

// 		c1[0] = make(chan *int)
// 		n1 := 12
// 		n2 := 24
// 		c1[0] <- new(int)
// 		c1[0] <- new(int)
// 		c1[0] <- &n1

// 		c1[1] = make(chan *int)
// 		c1[1] <- new(int)
// 		c1[1] <- new(int)
// 		c1[1] <- &n2

// 		close(c1[0])
// 		close(c1[1])
// 		close(c1)
// 	}()

// 	return c1
// }
