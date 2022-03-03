package main

import (
	"fmt"
)

func main() {
	var varaible1 int // declare int type varialbe "varaible1"
	fmt.Println(varaible1) // print value 0
	var variable2 string
	fmt.Println(variable2) // print ""
	varaible1 = 123
	variable2 = "hello world"
	fmt.Println(varaible1, variable2) // print 123 "hello world"
	variable5, variable6 := "multiple", 2 
	fmt.Println(variable5, variable6) // print "multiple" 2
	var variable7, variable8, varialbe9 int // declare multiple variables of same type
	variable7, variable8, varialbe9 = 123, 456, 789 // assign multiple values to multiple variables
	fmt.Println(variable7, variable8, varialbe9) // print 123, 456, 789
	var (
		variable10 int 
		variable11 string
		variable12 bool
	) // declaring multiple variables of different types at  the same time.
	fmt.Println(variable10, variable11, variable12) // print 0, "", false
	variable10, variable11, variable12 = 123, "multiple", true
	fmt.Println(variable10, variable11, variable12) // print 123, "multiple", true
	// variable13, variable14 := 123, 456
	// fmt.Println(variable13) // error variable14 declared and not used
	variable15, _ := 123, 456 // use of "blnak identifer"
	fmt.Println(variable15) // print 123
	


 


}