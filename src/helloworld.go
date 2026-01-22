package main

func main() {
	// they have their own zero value
	// var name string // empty string
	// var age int // 0
	// var permission bool // false
	// p := true // infers the type
	// var money float32 // 0.0

	// formating verbs
	// %T for type
	// %v for value
	// %q for a single-quoted character literal safely escaped with Go syntax.
	// %f decimal point
	// %t true of false

	// fmt.Printf("%q %v %t %f", name, age, p, money)
	// fmt.Printf("%q %v %v %v", name, age, p, money)

	// we can declare multiple variable in the same line
	// name, age := "mehedi", 20
	// fmt.Printf("%v %v", name, age)

	// type conversion
	// name := "mehedi"
	// number := int('a') // gives me character code
	// number := int('1') // nah this also gives me character code
	// number_in_string := "20"
	// number := []byte(number_in_string)
	// rune type is int32 and represents single unicode
	// println(number)

	// Constansts
	const PI = 3.14 // imutable :D
	// constants in go does his things in compile time.
	// that means if it need to calculate it will happen in compile time
	// and the it will only be that static value
	// in this case 3.14
	// TODO: read this text: https://www.reddit.com/r/golang/comments/1gll74w/go_constants_beyond_basics/

	// string formating
	// good := fmt.Sprintf("This is pi %5.1f", PI)
	// fmt.Print(good)

	// conditional
	// if true {

	// } else if false {

	// } else {

	// }

	// wow, NOTE: it will scoped to the condition only
	// if foo := len("something"); foo > 5 {

	// }

	// print(sub(1, 2))

	//go lang is pass by value -> (value get copied) not pass by reference -> (value doesn't get copied)

	// ignore the "_"
	// x, _ := multiple_return_values()
	// print(x)

}

// func sub(x int, y int) int {
// 	return x - y
// }

// succint (clear and concise) way to define this function
// func sub(x, y int, z string) int {

// 	return x - y
// }

// func multiple_return_values() (int, int) {
// 	return 1, 2
// }

// func get_coords() (x, y int) {
// 	// gets intialized with their respected zero values
// 	return // returns them // but dont do this this hurts readability
// }

// the top function is same as the bottom one
// func get_coords() (int, int) {
// 	var x int
// 	var y int
// 	return x,y
// }
