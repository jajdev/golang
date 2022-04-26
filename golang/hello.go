package main

import (
	"fmt"
	"reflect"
	"errors"
	"math"
)

func main() {
	// Hello World!
	var hello string = "Hello World"
	fmt.Println(hello)

	// Variables
	var x int = 5
	y := 11 //Shorthand for variable declaration
	sum := sum(x, y)
	var b bool = true
	var gpa float64 = 3.6
	var rune1 rune = 'B'

	// Displaying rune and its type
    fmt.Printf("Rune 1: %c; Unicode: %U; Type: %s \n", rune1, rune1, reflect.TypeOf(rune1))
      
	fmt.Printf("Sum = %d \n", sum)

	// if statements
	if sum > 10 {
		fmt.Println("Sum is greater than 10")
	} else if sum == 10 {
		fmt.Println("Sum is equal to 10")
	} else {
		fmt.Println("Sum is less than 10")		
	}
	if b {
		fmt.Printf("True! GPA = %f \n", gpa)
	}

	//Arrays
	var array [5]int //array length is fixed to it's type
	array[2] = 88
	fmt.Println(array)

	//Slices
	slice := []string{"This", "is", "a", "slice"} // Initialized as []T or []T{} or  []T{value1, value2, value3, ...value n}
	slice = append(slice, "!!") //slice has dyanmic sizing, so append can be used to add elements at the end of the slice
	fmt.Println(slice)
	fmt.Printf("Length of the slice: %d \n", len(slice))

	//Maps
	vertices := make(map[string]int) // Initialized as map[key]value
	vertices["triangle"] = 3
	vertices["square"] = 4
	vertices["hexagon"] = 6
	fmt.Println(vertices)
	delete(vertices, "triangle")
	fmt.Println(vertices)
	fmt.Printf("Vertices in a square: %d \n", vertices["square"])

	//Loops
	for i:= 0; i < 5; i++ { //for loop
		fmt.Println(i)
	}

	i := 0
	for i < 5 { //while loop
		fmt.Println(i)
		i++
	}

	for index, value := range slice { //loop over an array
		fmt.Println("index:", index, "value:", value)
	}

	for key, value := range vertices { //loop over a map
		fmt.Println("key:", key, "value:", value)
	}

	//Function call
	result, err := sqrt(-16)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	//Using struct
	p := person{name: "Jimmy", age:40}
	fmt.Println(p)
	fmt.Printf("Age: %d \n", p.age)

	//Pointers
	fmt.Println("Pointer:", &i)
	i = 7
	inc_PassByValue(i)
	fmt.Println(i)
	inc_PassByReference(&i)
	fmt.Println(i)
}

//Functions in the format
//func functionName(param type...) returnType {}
func sum(x int, y int) int {
	return x + y
}

func sqrt(x float64) (float64, error) { //multiple return types are allowed
	if x < 0 {
		return 0, errors.New("Cannot calculate sqrt of negative numbers")
	}
	return math.Sqrt(x), nil //In Go, nil is the zero value for pointers, interfaces, maps, slices, channels and function types, representing an uninitialized value.
}

func inc_PassByValue (x int) {
	x++
}

func inc_PassByReference (x int) {
	*x++
}

//Struct
type person struct {
	name string
	age int

}




