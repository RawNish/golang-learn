package main
import "fmt"

func main(){
	type Stud struct{
		Name string	
		Age int
	}
	var Nish Stud = Stud("Nishchcaya",20);
	fmt.Println(" hi my name is : ", Nish.Name); 
}