package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	artifacturl := "https://example.com/artifact.tar"
	artifactname := "elasticsearch.tar.gz"

	err := downloadartifact(artifacturl, artifactname)
	if err != nil {
		fmt.Println("doesnt work buddy")
	}

}

func downloadartifact(artifacturl string, artifactname string) error {

	res, err := http.Get(artifacturl)
	if err != nil {
		fmt.Println("There is and error")
	}
	defer res.Body.Close()

	out, err := os.Create(artifactname)
	if err != nil {
		fmt.Println("Doesnt work buddy")
	}

	defer out.Close()

	cop, err := io.Copy(out, res.Body)
	if err != nil {
		fmt.Println("Doesnt work buddy")
	}

	fmt.Println(cop)
	return nil
}
