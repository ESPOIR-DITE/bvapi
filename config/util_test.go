package config

import (
	"fmt"
	"testing"
)

func TestGenerateBase64ID(t *testing.T) {
	res := GenerateID("user-")
	fmt.Println(res)
}
