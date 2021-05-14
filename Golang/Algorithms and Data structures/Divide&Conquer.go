package main

import "fmt"
//divide and conquer
//The idea is to solve the sub-problems & combine to form one 1 solution

func fibonacci (k int) int{
	
	if k <= 1{
		return 1
	}
	return fibonacci(k-1)+fibonacci(k-2)
}

func main(){
var m int = 5

for m=0; m<8; m++{
	
	var fib = fibonacci(m)
	fmt.Println(fib) 	
	
}
	
}
/* Result
1
1
2
3
5
8
13
21

*/

