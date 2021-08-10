package domain

import "time"

type User struct {
	Id          string    `json:"id"`
	Contact     string    `json:"contact"`
	Name        string    `json:"name"`
	Surname     string    `json:"surname"`
	DateOfBirth time.Time `json:"dateOfBirth"`
	Bios        []byte    `json:"bios"`
}
