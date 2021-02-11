package main

import (

"fmt"
"time"
)

func f(){
	fmt.Println("f funtion")
}

func main(){
//go routine
go f()

//time.Sleep makes the main funtion sleep for 1 sec as f()finishes 
time.Sleep(1 * time.Second)
fmt.Println("main funtion")

}
