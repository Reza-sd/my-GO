package main
import (
    "fmt"
    "os"
)
func main() {
 //--------------------------------------------------------
   fmt.Println ("sum = int 1 + int 2")
    num1 := 0
    num2 := 0
    fmt.Print("Enter number 1: ")
    fmt.Scan(&num1)
    fmt.Print("\nEnter number 2: ")
    fmt.Scan(&num2)

    sum := num1+num2

    fmt.Println(sum)
    fmt.Printf("%03d   <---\n", sum)
    //------------------------------------------------
    fmt.Println ("\n--------------------\nsum = float 1 + float 2")
    ft1 := 0.00
    ft2 := 0.00
    fmt.Print("Enter float 1: ")
    fmt.Scan(&ft1)
    fmt.Print("\nEnter float 2: ")
    fmt.Scan(&ft2)

    sumf := (ft1+ft2) / 100;

    fmt.Println(sumf)
    //%10s specifies a minimum width of 10 characters for a string
    //%.2f specifies two decimal places for a floating-point number.
    // %% Prints a literal percent sign (%)

    fmt.Printf("%.3f %% %10s \n", sumf, "Percent")
    //-------------------------------------------------

    os.Exit(0) // Indicate successful completion
  // This code will not be executed
  fmt.Println("After os.Exit()")

}