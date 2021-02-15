// Port scanner,link used scanme.nmap.org
package main

import (

"fmt"
"net"
)

func main(){

for i:=1; i<=1024; i++{
	
	//Sprintf inside the net package transforms the int to a string
	//go routine
	go func(j int){
	address := fmt.Sprintf("scanme.nmap.org:%d",j)
	
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return
		
}
conn.Close()
fmt.Printf("%d open\n",j)
}(i)

}
}

// the program exits almost immediately
/* A single go routine is launched per connection & the main go routine doesnt
 wait fot the connection to take place therefore the ocde compeletes & exits as soon as the for loop is done 
*/

