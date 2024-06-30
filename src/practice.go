package main

import (
	"fmt"
)

//creating  custom function
type Department struct{
	Name string
	Location  string
	Employee
}
type Employee struct{
	Name string
	Age int
	Titel string
}
func myDepratment(){
	dept := Department{Name: "It Department",Location: "Bangladesh",
    Employee: Employee{Name: "Rahat",Age: 24,Titel: "programmer"},
}
	fmt.Println(dept.Employee,dept)
}
//Now play with loop and array
func arrasAndLoop(){
	var a[5]int
	a[1] = 200
	fmt.Println("The array is",a,"The length is", len(a))
	 b:= []int{2,3,4 ,5}
	// b = append(b, 10)
	fmt.Println("The array is",b,"The length is", len(b))
	//LOOPING IN THE ARRAY GOLANG
	for i := 0; i < len(b); i++ {
		fmt.Println("The value of the array index is ", b[i])
	}
}


/*This is for creating function  and returning
multiple value*/
// func mul()(sum, mul int){
// 	z:= 24
// 	p := 30
// 	return z+p, p*z
// }

// func  sum(){
// 	var x int
// 	x = 5
// 	y:= 20
// 	fmt.Println("The value is ", x,"Y value is",y);
// }
// func main() {
// 	fmt.Println("hello there")
// 	sum()
// 	a , b:=mul()
// 	fmt.Println("Return value is ", a, "Other value",b)
// 	arrasAndLoop()
// 	myDepratment()
}
