package main

import (
	"fmt"
	"io"

	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("web reqquest")
	response, error := http.Get(url)

	if error != nil {
		panic(error)
	}
	fmt.Printf("Response of type %T\n", response)
	fmt.Println(response)

	dataBytes, error :=io.ReadAll(response.Body)

	if(error!=nil){
		panic(error)
	}
	content := string(dataBytes)
	fmt.Println(content)
	defer response.Body.Close() //callers responsibility to close the connection
}
