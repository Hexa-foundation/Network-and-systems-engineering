package main

import (
"fmt"

)

func main() {

var k,l,m int
var arr[10][10][10] int


for k = 0; k < 10; k++ {
for l=0; l < 10; l++ {
for m=0; m < 10; m++ {
	
arr[k][l][m] = 1
fmt.Println("Element value ",k,l,m," is", arr[k][l][m])
}
}
}
}

/* result
Element value  0 0 0  is 1
Element value  0 0 1  is 1
Element value  0 0 2  is 1
Element value  0 0 3  is 1
Element value  0 0 4  is 1
Element value  0 0 5  is 1
Element value  0 0 6  is 1
Element value  0 0 7  is 1
Element value  0 0 8  is 1
*/




