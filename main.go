package main

import (
	"fmt"
	entity "suggestApp/enity"
	"suggestApp/repository/mysql"
)

func main() {

}

func testUserMySqlRepo() {
	mysqlRepo := mysql.New()
	createUser, err := mysqlRepo.Register(entity.User{ID: 0, PhoneNumber: "0962", Name: "alierza"})
	if err != nil {

		fmt.Printf("error %v\n", err)
	}

	fmt.Printf("User created with ID: %v\n", createUser)

	isUnique, UErr := mysqlRepo.IsUniquePhoneNumber(createUser.PhoneNumber + "22")

	if UErr != nil {
		fmt.Printf("error %v\n", UErr)
	}

	fmt.Println("is unique", isUnique)
}
