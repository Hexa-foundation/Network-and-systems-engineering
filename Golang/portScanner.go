// Port scanner,link used scanme.nmap.org
package main

import (

"fmt"
"net"
)

func main(){

for i:=1; i<=1024; i++{
	
	//Sprintf inside the net package transforms the int to a string
	address := fmt.Sprintf("scanme.nmap.org:%d",i)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		//port is closed or filtered
		continue
}
conn.Close()
fmt.Printf("%d open\n",i)

}
}


/* Result
 $ go build portScanner 
 $ ./portScanner
 
 $ 22 open
 $ 80 open
  
 */

