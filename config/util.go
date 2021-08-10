package config

import (
	"fmt"
	uuid "github.com/nu7hatch/gouuid"
)

func GenerateID(prefix string) string {
	id, err := uuid.NewV4()
	if err != nil {
		fmt.Println(err)
	}
	return prefix + id.String()
}
