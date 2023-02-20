package test

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestBase64(t *testing.T) {
	input := "=小信1n!"
	fmt.Println("input:", input)

	e := base64.StdEncoding.EncodeToString([]byte(input))
	fmt.Println("RawStdEncoding:", e)

	decoded, _ := base64.StdEncoding.DecodeString(e)
	fmt.Println("RawStdEncoding:", string(decoded))
}
