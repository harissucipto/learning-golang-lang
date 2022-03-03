package main // use of package keyword

import (
	"fmt"
	"time"
) // use of import keyword

type person struct { // use of type keyword
	name string
	age int
}

type interface_example interface {
	interface_function_int () int
	interface_function_str () string
} // use of interface keyword


func interface_function_str() (string) {
	return "returning form function" // use of return keyword
}

func main() { // func keyword
 var str = "hello world" // use of var keyword
 const num = 10 // const keyword
 arr := [6] int {10, 20, 30, 40, 50, 60}
 map_data := map[string] int {
	 "hello": 80,
	 "world": 90,
 } // declare map

 cs1 := make(chan string) // declare channel
 cs2 := make(chan string)
 go func() {
	 time.Sleep(1 * time.Second)
	 cs1 <- "one"
 }()
 fmt.Println(str, num, map_data, cs1)
 if num > 100 { // use of if keyword
	 fmt.Println("num is greater than 100")
 } else { // use of else keyword
	 fmt.Println("num is not greater than 100")
 }

 for pos, val := range arr { //use of for and range
	if (pos == 5) {
		break // break keyword
	}
	if (pos == 1) {
		continue // continue keyword
	}
	fmt.Printf("value at index %d is %d\n", pos, val)

 }
 
 fmt.Println(interface_function_str())
 switch num { // switch keyword
 	case 1: // use of case keyword
		fmt.Println("one")
 case 2:
		fmt.Println("two")
 case 3:
		fmt.Println("three")
 default: // default keyword
		fmt.Println("No match")

	switch num {
		case 1:
			fmt.Println("one")
		case 2:
			fmt.Println("two")
		case 10:
			fmt.Println("three")
		  fallthrough // use of fallthrough keyword
		default:
			fmt.Println("No match")
	} 

	 defer interface_function_str() // use of defer keyword

	 go interface_function_str() // use of go keyword
	 select {
	 		case msg1 := <-cs1: 
				fmt.Println("received", msg1)
	 		case msg2 := <- cs2:
				fmt.Println("received", msg2)
	 } //use of select keyword

	 goto gotlabel // use of goto keyword
	 gotlabel:
	 fmt.Println("inside goto statment")


 }


}


/*

Keywords are predefined reserved words specific to the programming language that have special meaning for the compiler. 
Golang supports 25 keywords

import : This keyword is used to import a package in a Go program.
type: This is used to create a new type as per the requirement.
const: This specifies a character, numeric, boolean, or string value.
var: This keyword is used to declare a variable containing a specific name and value.
func: This keyword is used to declare functions with a specific name, parameter, result type, and body.
map: This keyword declares the data of type map in a Go program.
chan: The chan keyword is used to define a channel in a Go program.
struct: A user-defined type "struct" is declared by using this keyword.
interface: This keyword is used to define a custom data type "interface" in Golang.
if: If the keyword specifies "if statement," it needs to evaluate the given condition.
else: else keyword executes the statement if the condition given in the block is false.
for: This keyword executes for a statement that repeatedly executes a block of code.
range: This keyword is used with for loop to iterate over a specific given data.
break: This defines the break statement inside a loop to terminate the current loop execution and resumes the program control next to the loop.
continue: This defines continue statement inside a loop to jump the program control on the next iteration.
goto: This executes the goto statement and transfers the program control to a labelled statement.
return: This executes the return statement by returning the result to the caller.
switch: This defines the switch statement in the Go program.
case: This defines various switch cases in a switch statement.
select: This keyword defines the select statement that works on multiple send/receive channel operations.
default: This executes the default case in a switch statement.
fallthrough: This executes the fallthrough statement that transfers the program control to the first statement of the case next to the currently executed case.
defer: This executes the defer statement that delays the function execution.
go: This keyword is used to handle Goroutines.

*/


/*
The identifier is the name of the class, functions, statement label, constant, package name or variables that uniquely identifies an objec
Go provides a special type of identifier called "blank identifier" that is denoted as _. Blank identifier works as a placeholder that can have any value, but we cannot read it in our program.

Valid
- _ = 12 // blank identifier
- var Version = 4.0 // exported identifers
- var name = "Ansh" // non-exported identifers
- var _12 = "hello world" // non exported identifiers


Invalid
- var 12 = "hello world" // should not start with digits
- var count$ = 23 // can't have special characters
- var default = 5.0 // can't use reserve keywords as identifers

*/