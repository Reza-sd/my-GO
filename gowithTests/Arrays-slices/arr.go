package main

import (
	"fmt"
	//"reflect"
)

// --------------Sum---------------------
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}

	result := sum
	return result
}

// -------------SumAll----------------------
func SumAll(numbersToSum ...[]int) []int {

	/*
	fmt.Println("numbersToSum----->", numbersToSum)
	fmt.Println("++++++>type of numbersToSum= ", reflect.ValueOf(numbersToSum).Kind())

	fmt.Printf("----->type of =  %T\n", numbersToSum)
	lengthOfNumbers := len(numbersToSum)
	fmt.Println("lengthOfNumbers----->", lengthOfNumbers)
	item1 := numbersToSum[0]
	
	fmt.Println("item1----->", item1)
	fmt.Printf("----->type of item1=  %T \n ", item1)
	//reflect.TypeOf(tst)
	fmt.Println("*****>type of item1= ", reflect.TypeOf(item1))
	//reflect.ValueOf(variable1).Kind()
	fmt.Println("++++++>type of item1= ", reflect.ValueOf(item1).Kind())
	sum1 := Sum(item1)
	sums := item1
	fmt.Println("sums--->", sums)
	fmt.Println("sum1--->", sum1, reflect.ValueOf(sum1).Kind())
	//sums[0] =
	*/
	lengthOfNumbers := len(numbersToSum)
	//fmt.Println("lengthOfNumbers----->", lengthOfNumbers)
	sliceResult := []int{}
	itemInput := []int{}
	sumItemInput := 0
	
	for i:=0;i<lengthOfNumbers;i++ {
	itemInput = numbersToSum[i]
  sumItemInput = Sum(itemInput)
	fmt.Println("sumItemInput--->",sumItemInput)
	
	sliceResult = append(sliceResult, sumItemInput)
	//fmt.Println("sliceResult--->", sliceResult)
	}
	
	result := sliceResult
	return result
}

//-----------------------------------
