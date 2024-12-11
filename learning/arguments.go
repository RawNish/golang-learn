//howe to capture arguments that i passed while running the go file
package main
import (
	"fmt"
	"os"
)

func main(){
 args:= os.Args
 
 if(len(args)<2){
	fmt.Println("NOt enough arguments\nExiting now")
	os.Exit(1)

 }
 //fmt.Println("Arguemnts for thisrun are \n" , args[1],args[2])
 fmt.Println("Arguemnts for thisrun are \n" , args[1:])

}