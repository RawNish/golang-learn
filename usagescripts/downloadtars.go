package main

import (
	"fmt"
	"os/exec"
)

func main() {

	parentcmd := "wget"
	var url string
	fmt.Println("Enter the url of the tar file")
	fmt.Scan(&url)
	out := exec.Command(parentcmd, url)

	fmt.Println("Output:", out)
}
