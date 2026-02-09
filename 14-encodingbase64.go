package main

import "encoding/base64"

func main() {
	var encoded = base64.StdEncoding.EncodeToString([]byte("Hello, World!"))
	println(encoded)

	var decoded, err = base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		println("Error decoding:", err.Error())
	} else {
		println(string(decoded))
	}
}
