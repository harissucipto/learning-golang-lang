// referenced data types as they contain the address of the dynamically created objects in memeory

// ponter
package main

import "fmt"


func main() {
	// pointer
	// var name_of_pointer *type_of_ponter 
	var intval int= 12
	var pval = &intval
	fmt.Println(pval) // prints the address of the variable
	fmt.Println(*pval) // prints the value of the variable
	*pval = 78 // change the value of the variable
	fmt.Println(intval) // prints the new value of the variable

	// Array
	//  var name_of_array[size_of_array] type_of_array_elements
	// or
	// name_of_array := [size_of_array] type_of_array_elements{elements_of_array}

	var arrval[3] string 
	// arrval = ["abc", "def", "ghi"] // error "syntax error"
	arrval[0] = "abc"
	arrval[1] = "def"
	arrval[2] = "ghi" // will assign the value to the array by accessing the index
	fmt.Println(arrval)
	arrval2 := [3] string{"abc", "def", "ghi"} // will assign the value to the array by accessing the index
	fmt.Println(arrval2)


}