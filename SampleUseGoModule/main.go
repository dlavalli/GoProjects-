package main

import (
	"fmt"
	"github.com/dlavalli/GoProjects-/SampleGoModule/calc"
	"github.com/dlavalli/GoProjects-/SampleGoModule/geometry"
)

func main() {
	add1 := calc.Add(2, 3)
	geo1 := geometry.CubeVolume(5)
	fmt.Printf("Find add 2 + 3: %d,   Cube volume of 5: %d  \n", add1, geo1);
}