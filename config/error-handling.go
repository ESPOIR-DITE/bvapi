package config

import (
	"fmt"
	"log"
)

func CheckErr(value interface{}, err error, location string) (interface{}, error) {
	if err != nil {
		fmt.Println("ERROR\t", log.Ldate|log.Ltime)
		return value, err
	}
	return nil, nil
}
