package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	entity "suggestApp/enity"
	"suggestApp/repository/mysql"
	userservice "suggestApp/service/userserice"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/users/register", userRegisterHandler)
	mux.HandleFunc("/health-check", healthCheckHandler)
	log.Print("serve on 8080")
	server := http.Server{Addr: ":8080", Handler: mux}
	log.Fatal(server.ListenAndServe())

}

func userRegisterHandler(writer http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(writer, `{"message":"Method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	data, err := io.ReadAll(req.Body)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error on read":"%s"}`, err.Error())))
		return
	}

	var UReq userservice.RegisterRequest

	err = json.Unmarshal(data, &UReq)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error on json unmarshal":"%s"}`, err.Error())))
		return
	}

	mySqlRepo := mysql.New()

	userSvc := userservice.NewService(mySqlRepo)

	_, err = userSvc.Register(UReq)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error on register":"%s"}`, err.Error())))
		return
	}

	writer.Write([]byte(fmt.Sprintf(`{"message":"User registered successfully"}`)))
	return
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `{"message":"everything is ok"}`)
}

func testUserMySqlRepo() {
	mysqlRepo := mysql.New()
	createUser, err := mysqlRepo.Register(entity.User{ID: 0, PhoneNumber: "0962", Name: "alierza"})
	if err != nil {

		fmt.Printf("error %v", err.Error())
	}

	fmt.Printf("User created with ID: %v\n", createUser)

	isUnique, UErr := mysqlRepo.IsUniquePhoneNumber(createUser.PhoneNumber + "22")

	if UErr != nil {
		fmt.Printf("error %v", UErr.Error())
	}

	fmt.Println("is unique", isUnique)
}
