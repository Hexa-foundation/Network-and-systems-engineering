package main

import
(
"log"
"os"
"fmt"
)

	type FooReader struct{}
	
	//Reads reads data from stdin
	func(fooReader *FooReader)Read(b[]byte)(int,error){
		fmt.Print("in >")
		return os.Stdin.Read(b)
	}
	//FooWriter defines an io.Writer to write to Stdout
	type FooWriter struct{}
	//write data to Stdout
	func(fooWriter *FooWriter)Write(b[]byte)(int,error){
		fmt.Print("out >")
		return os.Stdout.Write(b)
	}
	func main(){
		
	//instantiate reader and writer
	var(
	reader FooReader
	writer FooWriter
	)
	//create buffer to hold input/output
	input := make([]byte,4096)
	
	//use reader to read input
	s,err := reader.Read(input)
	if err != nil{
		log.Fatalln("Unable to read data")
	}
	fmt.Println("Read %d bytes from stdin \n",s)
	
	//Use writer to write input
	s,err = writer.Write(input)
	if err != nil{
		log.Fatalln("Unable to write data")
	}
	fmt.Printf("Wrote %d bytes to stdout\n",s)
}

	
/* output

in >hello
Read %d bytes from stdin 
 6
out >hello
Wrote 4096 bytes to stdout


***Its a bit buggy but corrections can be made
*/ 
