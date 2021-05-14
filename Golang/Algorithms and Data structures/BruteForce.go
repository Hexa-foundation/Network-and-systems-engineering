package main

import "fmt"

//find element method given array & k element
func findElement(arr[10] int,k int)bool{
	var i int
	for i=0; i<10; i++{
		if arr[i] == k{
			return true
		}
	}
	return false
}

func main(){
	
	var arr = [10]int{4,6,7,3,10,9,5,7,8}
	var check bool = findElement(arr,10)
	
	fmt.Println(check)
	
	var check2 bool = findElement(arr,12)
	fmt.Println(check2)
	
}

/* result
 true
 false
  
 */

