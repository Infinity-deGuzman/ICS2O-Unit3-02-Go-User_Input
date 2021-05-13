// Created by: Infinity de Guzman
// Created on: May 2021
//
// This program calculates the area of a trapezoid

package main

import (
  "fmt"

	"github.com/leekchan/accounting"
)

func main() {
	// This function prints the address
	var length float64
	var width float64
	var height float64

	// input
  accountingFormater := accounting.Accounting{Precision: 2}

	fmt.Println("This program calculates the volume of a pyramid.")
	fmt.Println()
	fmt.Print("Enter the length: ")
	fmt.Scanln(&length)
	fmt.Print("Enter the width: ")
	fmt.Scanln(&width)
	fmt.Print("Enter the height: ")
	fmt.Scanln(&height)

  //process
  var volume = (length * width * height) / 3

	// output
	fmt.Println("The volume is: ", accountingFormater.FormatMoney(volume) , "cm³")
	fmt.Println("\nDone.")
}
