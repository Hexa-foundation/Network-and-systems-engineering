package main

import (

"fmt"
"container/list"
)

func main(){
	
var intList list.List
//PushBack adds values to the list
intList.PushBack(1)
intList.PushBack(2)
intList.PushBack(3)

for element := intList.Front(); element != nil; element=element.Next(){
	fmt.Println(element.Value.(int))
}
	
	
	
}
/*result
1
2
3
*/
