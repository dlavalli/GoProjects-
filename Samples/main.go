package main

import (
	"fmt"
	"github.com/MyGoProjects/Samples/tools"
	"math"
)

// ##################### Interface / Struct - case 1 ##########################################
// Since using pointer to as receiver, modifications made on instances passed to functions affect
// the original instance passed
// ie: passed by refence (pointer)
// ############################################################################################
type valsetter interface {
	setVal(int)
}
type namesetter interface {
	valsetter
	setName(string)
}
type mystrid struct {
	val int
	name string
}

// This method expects a pointer as instance
func (s *mystrid) setVal(v int) {
	s.val = v
}

// This method expects a pointer as instance
func (s *mystrid) setName(n string) {
	s.name = n
}

// ##################### Interface / Struct - case 2 ##########################################
// Since not using pointer to as receiver, no modifications made on instances passed to functions
// ie: passed by copy
// ############################################################################################

type Shape interface {
	Area() float64
	Perimeter() float64
}
type Rect struct {
	width float64
	height float64
}
type Circle struct {
	radius float64
}
func (r Rect) Area() float64 {
	return r.width * r.height;
}
func (r Rect) Perimeter() float64 {
	return 2 * (r.width + r.height);
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius;
}
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius;
}


// ########################### Example with array / slice 1 ####################################

func printSlice(sl []int) {
	fmt.Printf("%#v :: len %d, cap %d \n", sl, len(sl), cap(sl))
}

// ########################### Example with map 1 ####################################



// ############################################################################################

func main() {
	// string ops
	r := tools.ReverseString("Hello world!")
	tools.PrintStringChars(r)

	// Fibonacci example
	f1 := tools.FibonacciRecur(4)
	fmt.Printf("Fibonnaci of: %d is %d\n\n", 4, f1)


	// ##################### Interface / Struct - case 1 ##########################################

	// My struct example
	ms := mystrid{val: 1, name: "test"}
	ms.setVal(2)
	ms.setName("nextTest")
	fmt.Printf("mystruct: %v\n", ms)

	// Using defined interface but need a ref because the methods receiver defined above expect a pointer to mystrid instance
	var v valsetter
	v = &ms
	fmt.Printf("Using %v %T\n", v, v)

	// Using defined interface but need a ref because the methods receiver defined above expect a pointer to mystrid instance
	var vn namesetter
	vn = &ms
	fmt.Printf("Using %v %T \n\n", vn, vn)

	res := tools.Max(10, 15)
	fmt.Printf("Max between %d and %d is: %d\n", 10, 15, res)

	// ##################### Interface / Struct - case 2 ##########################################

    var sh1, sh2 Shape
	r1 := Rect{width: 10, height: 10}
	c1 := Circle{radius: 5}
	sh1 = r1
	sh2 = c1
	fmt.Printf("Rect width: %d height: %d area: %d perimeter: %d\n", 10, 10, int(sh1.Area()), int(sh1.Perimeter()))
	fmt.Printf("Circle radius: %d area: %f perimeter: %f \n", 5, sh2.Area(), sh2.Perimeter())

	// ########################### Example with array / slice 1 ####################################
	var sl []int
	for i := 0; i < 17; i++ {
		sl = append(sl, i)
	}
	printSlice(sl)

	arr := []int {1,2,3,4,5,6,7,8,9}
	fmt.Printf("The array sum is: %d\n", tools.SumArray(arr));

	// ########################### Example with map 1 ####################################
	mv := map[string]int{
		"a" : 1,
		"b" : 2,
		"c" : 3,
	}

	nv := make(map[string]int)
	nv["a"] = 1
	nv["b"] = 2
	nv["c"] = 3

	var ov map[string]int
	ov = make(map[string]int)
	ov["a"] = 1
	ov["b"] = 2
	ov["c"] = 3

	// Order of key returned is not garanied and can change from one call to the next
	for key, val := range mv {
		fmt.Printf("An mv map entry: key : %s value: %d\n", key, val);
	}
	for key, val := range nv {
		fmt.Printf("An nv map entry: key : %s value: %d\n", key, val);
	}
	for key, val := range ov {
		fmt.Printf("An ov map entry: key : %s value: %d\n", key, val);
	}

	// ########################### Function pointer / Closures ####################################
}
