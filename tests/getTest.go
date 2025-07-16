package main

import (
	"fmt"

	"github.com/estefspace/scout"
)

func main() {
	client := scout.NewClient("https://rickandmortyapi.com/api/")

	resp, err := client.PleaseGetThis("/character")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resp)
}
