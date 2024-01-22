package main

import "fmt"

type student struct {
  name string
  age int
}

//----------------------------
func changeName (s *student, name string) {
  s.name = name
}
//----------------------------
func main () {

 chris := student{name: "BOb", age: 22}
  chris.age = 35000
  name := "Chris"
  changeName(&chris,name)

  fmt.Println("name=",name,"=",chris)

  
}
//---------------------------------