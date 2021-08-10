package mysql

import (
	"bvapi/config"
	"bvapi/domain"
	"bvapi/repository"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func CreateUser(user domain.User) (domain.User, error) {
	db, err := repository.Repository()
	if err != nil {
		return domain.User{}, err
	}
	stmt, err := db.Prepare("INSERT user SET id=?,contact=?,names=?,surname=?,date_of_birth=?,bios=?")
	if err != nil {
		return domain.User{}, err
	}
	res, err := stmt.Exec(user.Id, user.Contact, user.Name, user.Surname, user.DateOfBirth, user.Bios)
	config.CheckErr(domain.User{}, err, "repository")
	id, err := res.LastInsertId()
	config.CheckErr(domain.User{}, err, "repository")
	fmt.Println(id)
	return user, nil
}

func readUsers(user domain.User) (domain.User, error) {
	var entity domain.User
	db, err := repository.Repository()
	if err == nil {
		return entity, err
	}
	res, err := db.Query("SELECT * FROM cities")

	defer res.Close()

	if err != nil {
		log.Fatal(err)
	}

	for res.Next() {
		var entity domain.User
		var entities []domain.User
		err := res.Scan(&entity.Id, &entity.Name, &entity.Contact, &entity.Bios, &entity.DateOfBirth, &entity.Surname)
		if err != nil {
			log.Fatal(err)
		} else {
			entities = append(entities, entity)
		}
		fmt.Printf("%v\n", entity)
	}
	return user, nil
}

func readUser(user domain.User) (domain.User, error) {
	var entity domain.User
	db, err := repository.Repository()
	if err == nil {
		return entity, err
	}
	res, err := db.Query("SELECT * FROM cities")

	defer res.Close()

	if err != nil {
		log.Fatal(err)
	}

	for res.Next() {
		err := res.Scan(&entity.Id, &entity.Contact, &entity.Name, &entity.Surname, &entity.Bios, &entity.DateOfBirth)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v\n", entity)
	}
	return entity, nil
}
