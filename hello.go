package main

import "fmt"

func main() {
	fmt.Println("Hello, World")

	x := 5
	y := 7
	sum := x + y
	fmt.Println(sum)

	// To initialize an array with a numbered length and type do:
	var a [5]int
	// Assign variables at an index with:
	a[2] = 7
	// Using type inference, and passing values for each index of the array
	b := [5]int{5, 4, 3, 2, 1}
	fmt.Println(b)
	// To create an array without a fixed length called a slice don't pass a variable into the array initializer
	c := []int{3, 2, 1} // [3, 2, 1]
	// Being a slice you can now use built in functions to append variables to the slice
	c = append(c, 13)
	// Slices do not mutate the array, they create a new array(slice)
	fmt.Println(c) //  // [3, 2, 1, 13]

	// To create a map you have to use the built in make funciton
	verticies := make(map[string]int)
	// This creates a map with key value pairs where the key is a string and the value is the type of int
	// To access the keys you can index in to the map to recieve the value of the pair
	verticies["triangle"] = 3
	verticies["square"] = 4
	fmt.Println(verticies) // map[square:4 triangle:3]

	// for loops are the only type of loops in go
	// It contains 3 parts like many other languages the initializer, the condition and the iterator
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	} // 0, 1, 2, 3, 4

	// To create a while loop in go, you remove the initializer and the iterator from the loop declaration
	l := 0
	for l < 5 {
		fmt.Println(l)
		l++
	} // 0, 1, 2, 3, 4
	// you can also loop through an array or a slice using a for loop and a range
	arr := []string{"a", "b", "c"}
	// after intializing the array you can set variables to hold the index and the value of each index of the array/slice
	for index, value := range arr {
		fmt.Println("index:", index, "value:", value)
	}
	/*	index: 0 value: a
		index: 1 value: b
		index: 2 value: c */

	// this method also works with a map by creating a variable for the key and the value of the pair
	m := make(map[string]string)
	m["a"] = "alpha"
	m["b"] = "beta"

	for key, value := range m {
		fmt.Println("key:", key, "value:", value)
	} /*
		key: a value: alpha
		key: b value: beta */

	result := sumFunc(2, 3)
	fmt.Println(result) // 5

}

// when creating a new function you have to pass the types of the variables with the variable names
// and specify the return type of the function
func sumFunc(x int, y int) int {
	return x + y
} // then call it in the main function

// functions in go can also return more more than one thing
