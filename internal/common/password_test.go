package common

import (
	"fmt"
	"testing"
)

func TestSignPassword(t *testing.T) {
	password := "piao1234"
	fmt.Println(SignPassword(password))
}
