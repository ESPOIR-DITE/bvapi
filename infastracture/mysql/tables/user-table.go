package tables

import (
	"bvapi/config"
	"bvapi/domain"
	"bvapi/repository"
)

func CreateUser(user domain.User) bool {
	db, err := repository.Repository()
	//insert
	stmt, err := db.Prepare("INSERT user SET id=?, contact=?, user_names=?, surname=?,date_of_birth =?,bio=?")
	config.CheckErr(domain.User{}, err, "table")

	id := config.GenerateID("Us-")
	_, err = stmt.Exec(id, user.Contact, user.Name, user.Surname, user.DateOfBirth, user.Bios)
	config.CheckErr(domain.User{}, err, "table")
	return true
}
