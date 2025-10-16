package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"suggestApp/repository/mysql"
	"suggestApp/service/userservice"
)

const (
	JWT_SECRET = "your_jwt_secret"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/health-check", healthCheckHandler)
	mux.HandleFunc("/users/register", userRegisterHandler)
	mux.HandleFunc("/users/login", userLoginHandler)
	mux.HandleFunc("/users/profile", userProfileHandler)

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

	userSvc := userservice.NewService(mySqlRepo, JWT_SECRET)

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

	userSvc := userservice.NewService(mySqlRepo, JWT_SECRET)

	Resp, err := userSvc.Login(UReq)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error on login":"%s"}`, err.Error())))
		return
	}

	writer.Write([]byte(fmt.Sprintf(`{"message":"User logged in successfully", "token":"%s"}`, Resp.AccessToken)))

	return
}

func userProfileHandler(writer http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(writer, `{"message":"Method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	Authorization := req.Header.Get("Authorization")

	fmt.Println(Authorization)
	Req := userservice.ProfileRequest{UserID: 10}

	mySqlRepo := mysql.New()

	userSvc := userservice.NewService(mySqlRepo, JWT_SECRET)

	userSvc.Profile(Req)
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `{"message":"everything is ok"}`)
}
