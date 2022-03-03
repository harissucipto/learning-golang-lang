// Primitve data types
// dibangun dengan bahasa pemrograman dan selalu memiliki nilai tunggal

package main

import "fmt"

func main() {
	// Boolean
	var is_Red bool = true
	fmt.Printf("%v\n", is_Red)
	
	// Numberic
	var int32val int32 = 100
	var valint int = 76 // treats as int32 or int64 depends on implementation
	fmt.Println(int32val, valint)
	// var int64val int64 = 343;
	// var int16val int16 = int64val // error "cannot use int64val (type int64) as type int16 in assignment"
	var int16val int16 = int16(int32val) // works
	fmt.Println(int16val) // print 100
	// var int64val int64 = int32val // error "cannot use int32val (type int32) as type int64 in assignment"

	// String
	var strval string = "Grow with Golang"
	var strcopied string = strval 
	fmt.Println("original string is:", strval)
	fmt.Println("copied string is:", strcopied)
	fmt.Println("Address of strval:", &strval)
	fmt.Println("Address of strcopied:", &strcopied)


}