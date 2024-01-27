Instructions
Write a program that displays, on the standard output, the content of a file given as argument.
---------Allowed functions-----------

fmt.*, os.*, io/ioutil.*, .., --allow-builtin

----------------------------------------
$ go run .
File name missing

$ echo "Almost there!!" > quest8.txt

$ go run . quest8.txt main.go
Too many arguments

$ go run . quest8.txt
Almost there!!
-----------------------------------