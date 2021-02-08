package main

import (

 "fmt"
 "os"
 "net"

)

func main(){
	if len(os.Args) != 2 {
		
		fmt.Fprintf(os.Stderr, "Usage: %s hostname\n",os.Args[0])
		fmt.Println("Usage:", os.Args[0],"hostname")
		os.Exit(1)
	}
	name := os.Args[1]
	addr, err := net.ResolveIPAddr("ip",name)
	
	if err != nil {
		
		fmt.Println("Resolution error",err.Error())
		os.Exit(1)
	}
	fmt.Println("Resolved address is",addr.String())
	os.Exit(0)
	
}

/* Result
  A DNS lookup will be performed

$ ./IPAddressREsolver google.com
$ Resolved address is 216.58.223.110

*/
