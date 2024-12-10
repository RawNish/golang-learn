//howe to capture arguments that i passed while running the go file
package main
import (
	"fmt"
	"os"
)

func main(){
 args:= os.Args
 fmt.Println("Arguemnts for thisrun are %v\n" , args)
}