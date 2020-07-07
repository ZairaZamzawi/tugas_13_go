package main

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
)

func main() {
	var nama = "Zaira Maulidin Zamzawi"

	// encode
	var encodeString = base64.StdEncoding.EncodeToString([]byte(nama))
	fmt.Println("Encode Nama :", encodeString)

	// encrypt with sha
	var sha = sha1.New()
	sha.Write([]byte(encodeString))
	var encrypted = sha.Sum(nil)

	var shaString = fmt.Sprintf("%x", encrypted)
	fmt.Println(shaString)

}
