/*
--------tools-------------------

github.com/01-edu/z01.PrintRune, os.Args,
-------------what they ask------------------

Create a new directory called boolean.

The code below must be copied into a file called main.go inside of the boolean directory.

The necessary changes must be applied for the program to work.
---------example------------

$ go run . "not" "odd"
I have an even number of arguments
$ go run . "not even"
I have an odd number of arguments


$ go run . "a" "2" "Wb" "Qxds3"
$

Expected :

$ go run . "a" "2" "Wb" "Qxds3"
I have an even number of arguments
$

-----sudo-code-------

1- read arg
2- check if arg is even or odd

2-1 
result 1 = print "I have an even number of arguments"
result 2 = print "I have an odd number of arguments"

---------result 1 :-------
1- case 1 : "not" "odd"

----------result 2---------
case 1: "not even"


//----------------------
*/
// ---------------------------
//  takes int as argument and "it passed the argumet to even func"return true if it is negative
// it seems even function return a int value , is seems 1 one the value is 1

func isEven(nbr int) bool {
  if even(nbr) == 1 {
    return true
  } else {
    return false
  }
}

// ---------------------
func even(num int) int {
  var result int

  //---sucodoed : Dear my 2yrs son :  an even num is : is num when / 2 the remainder is 0
  var remainder int
  remainder = num % 2

  if remainder == 0 {
    //this number is even
    result = 1
  }

  return result
}

// ------------------------------

//-------------------------