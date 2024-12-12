package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	var choice int
	var por int = 1
	for por < 1000 {
		fmt.Println("This program is going to run docker prune \nInput the numbers based on their number \nPress 1 for image prune\nPress 2 for volume prune\nPress 3 for network prune\nPress 4 for container prune\nPress 5 for exiting")
		fmt.Scan(&choice)
		if choice == 5 {
			os.Exit(1)
		}
		parentCmd := "docker"

		resources := []string{"image", "volume", "network", "container"}
		cmnd := "prune"

		cmd := exec.Command(parentCmd, resources[choice-1], cmnd, "-f")
		out, err := cmd.Output()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Output: ", string(out))

	}
}
