//---------------------
Instructions
Create a new directory called point. *

The code below must be copied into a file called main.go inside the point directory.

The necessary changes must be applied so that the program works.

The function setPoint() must work with int.

//------------sample---------------
$ go run .
x = 42, y = 21
$
//---------------
github.com/01-edu/z01.PrintRune
//---------------
append(s S, x ...T) S  // T is the element type of S

s0 := []int{0, 0}
s1 := append(s0, 2)        // append a single element     s1 == []int{0, 0, 2}
s2 := append(s1, 3, 5, 7)  // append multiple elements    s2 == []int{0, 0, 2, 3, 5, 7}
s3 := append(s2, s0...)    // append a slice              s3 == []int{0, 0, 2, 3, 5, 7, 0, 0}
//---------------------------------------
s := append([]int{1, 2}, []int{3, 4}...)


//-----------------------------------