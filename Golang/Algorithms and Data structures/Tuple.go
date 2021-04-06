//A tuple is defined as a finite sorted list of elements
package main

import "fmt"

func PowerSeries(a int)(int, int){
	
	return a*a, a*a*a
}
	
	
func main(){
	var square int
	var cube int
	square, cube = PowerSeries(3)
	
	fmt.Println("Square", square, "Cube",cube)
	
	
}
/*Result

Square 9 Cube 27

*/
