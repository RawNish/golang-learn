package main
import "fmt"

func main(){
	var a [5]int=[5]int{1,2,3,4,5};
	fmt.Println("Array= ", a)
	var names [4]string= [4]string{"Nish","Rish","Rob","Ami"};
	fmt.Println("THe name array is :", names)
	var booarr [5]bool=[5]bool{true, true , false,false, true};
	fmt.Println("Boolean array is : ", booarr)



	//slices -> arrays that dont need a size before usage
	c := []int{1,2,3,4,5,9};
    fmt.Println("slice:", c)
	
	



}