package main

import (
	"fmt"
	"os"
)

func main() {
//-------------------------------------------
	file, err := os.Open("readme.txt")

	if err != nil {
		fmt.Println("err=", err)
		fmt.Println("err.Error()=", err.Error())
		fmt.Println("file=",file)
	} else {

		fmt.Println("file=", file)
    
    arr := make([]byte,10)
    file.Read(arr)
    fmt.Println("arr=",string(arr))
    //fmt.Println("file.Name()=", file.)
		fmt.Println("file.Name()=", file.Name())
    fmt.Println("file.Chdir()=", file.Chdir())
    file.Close()
	}
  //----------------------------
  fmt.Println()
  
  dir, errdir := os.Open("folder1")

  if errdir != nil {
    fmt.Println("errdir=", errdir)
    fmt.Println("errdir.Error()=", errdir.Error())
    fmt.Println("dir=",dir)
    
  } else {

    fmt.Println("dir=", dir)
    
    arrDir := make([]byte,10)
    file.Read(arrDir)
    fmt.Println("arrDir=",string(arrDir))
    
    fmt.Println("dir.Name()=", dir.Name())
    fmt.Println("dir.Chdir()=", dir.Chdir())
    dir.Close()
  }

  //-----------------------------------

}
