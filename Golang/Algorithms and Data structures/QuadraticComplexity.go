//processing time is directly proportional to the square of elements
package main

import "fmt"
func main(){
	var k,l int
	
	for k = 1; k <= 10;k++{
		fmt.Println("Multiplication Table",k)
		for l=1; l<=10; l++{
			var x int = l *k
			fmt.Println(x)
		}
	}
	
}

/*
 Result
 Multiplication Table 1
1
2
3
4
5
6
7
8
9
10
Multiplication Table 2
2
4
6
8
10
12
14
16
18
20
 
 */
