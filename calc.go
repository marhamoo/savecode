package main

import "fmt"

func main () {
	var angka1 float64
	var angka2 float64
	var operator string 

	fmt.Println("angka1: ")
	fmt.Scanln(&angka1)

	fmt.Println("angka2: ")
	fmt.Scanln(&angka2)

	fmt.Println("Masukan operator ( + - * / ) : ")
	fmt.Scanln(&operator)
	
	switch operator {
	case "+" : 
		fmt.Printf("%f %s %f = %f", angka1, operator, angka2, angka1 + angka2)
	
	case "-" : 
		fmt.Printf("%f %s %f = %f", angka1, operator, angka2, angka1 - angka2)

	case "*" : 
		fmt.Printf("%f %s %f = %f", angka1, operator, angka2, angka1 * angka2)

	case "/" : 
		if angka2 == 0.0 {
			fmt.Println("Pembagian dengan nol tidak terdefinisi")
		}else{
			fmt.Printf("%f %s %f = %f", angka1, operator, angka2, angka1 / angka2)
		}
	default:
		fmt.Println("invalid operator")
	} 
}