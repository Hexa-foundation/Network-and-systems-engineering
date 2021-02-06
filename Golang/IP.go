//takes an IP address and uses the net package to parse it
package main

import (

"fmt"
"net"
"os"
)
func main(){
	
      if len(os.Args) != 2{
		fmt.Fprintf(os.Stderr, "usage: %s ip-addr\n",os.Args[0])
        os.Exit(1)
}
name := os.Args[1]
addr := net.ParseIP(name)
if addr == nil {
	fmt.Println("Invalid IP address")
} else {
	fmt.Println("The address is",addr.String())
}
os.Exit(0)
	
}

/* Result
 $ ./IP 192.0.9.7
 $  The address is 192.0.9.7
 
 $  ./IP 256.0.0.0 
 $  Invalid IP address  
 */ 


