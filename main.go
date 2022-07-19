package main

import (
	"encoding/base64"
	"fmt"
)

func main() {

	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"
	sd, _ := base64.StdEncoding.DecodeString(secret)

	result := Reverse(string(sd))

	fmt.Println("ciphertext => ", secret)
	fmt.Println("original plaintext => " + result)
}

func Reverse(source string) (result string) {
	for _, value := range source {
		result = string(value) + result
	}
	return
}
