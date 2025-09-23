package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"suggestApp/repository/mysql"
	userservice "suggestApp/service/userserice"
)

func main() {

	// Mysqlrep := mysql.New()
	// userscsv := userservice.NewService(Mysqlrep)

	// q, BErr := userscsv.Login(userservice.LoginRequest{PhoneNumber: "09384850116", Password: "1312312222mm"})

	// if BErr != nil {
	// 	fmt.Println("Login failed", BErr.Error())
	// } else {

	// 	fmt.Printf("Login successful %v", q)
	// }

	mux := http.NewServeMux()
	mux.HandleFunc("/health-check", healthCheckHandler)
	mux.HandleFunc("/users/register", userRegisterHandler)
	mux.HandleFunc("/users/login", userLoginHandler)

	port := "5961"
	log.Print("serve on " + port)
	server := http.Server{Addr: ":" + port, Handler: mux}
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

func userLoginHandler(writer http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(writer, `{"message":"Method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	data, err := io.ReadAll(req.Body)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error on read":"%s"}`, err.Error())))
		return
	}

	var UReq userservice.LoginRequest

	err = json.Unmarshal(data, &UReq)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error on json unmarshal":"%s"}`, err.Error())))
		return
	}

	mySqlRepo := mysql.New()

	userSvc := userservice.NewService(mySqlRepo)

	_, err = userSvc.Login(UReq)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error on login":"%s"}`, err.Error())))
		return
	}

	writer.Write([]byte(fmt.Sprintf(`{"message":"User login successfully"}`)))
	return
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `{"message":"everything is ok"}`)
}
