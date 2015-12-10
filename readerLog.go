package main

import (
//  "flag" 
  "fmt"
  "os"
)


func readLogFile(filename string) {

 file, err := os.Open(filename)
    if err != nil {
	fmt.Print("Not found: "+filename)
    }
    defer file.Close()
}
