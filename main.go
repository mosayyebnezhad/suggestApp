package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"suggestApp/repository/mysql"
	"suggestApp/service/authservice"
	"suggestApp/service/userservice"
	"time"
)

const (
	JWT_SECRET        = "your_jwt_secret"
	Access_token      = "at"
	Refresh_token     = "rt"
	Access_tokenTime  = time.Hour * 24
	Refresh_tokenTime = time.Hour * 72
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

	authSVC := authservice.New(JWT_SECRET, Access_token, Refresh_token, Access_tokenTime, Refresh_tokenTime)

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

	userSvc := userservice.NewService(mySqlRepo, authSVC)

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
	authSVC := authservice.New(JWT_SECRET, Access_token, Refresh_token, Access_tokenTime, Refresh_tokenTime)

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

	userSvc := userservice.NewService(mySqlRepo, authSVC)

	Resp, err := userSvc.Login(UReq)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error on login":"%s"}`, err.Error())))
		return
	}
	data, err = json.Marshal(Resp)
	if err != nil {
		writer.Write([]byte(
			fmt.Sprintf(`{"error": "%s"}`, err.Error()),
		))

		return
	}

	writer.Write(data)

	return
}

func userProfileHandler(writer http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(writer, `{"message":"Method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	authSVC := authservice.New(JWT_SECRET, Access_token, Refresh_token, Access_tokenTime, Refresh_tokenTime)

	Authorization := req.Header.Get("Authorization")

	Authorization = strings.Replace(Authorization, "Bearer ", "", 1)
	jwtParsed, err := authSVC.ParseToken(Authorization)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error on parseJWT":"%s"}`, err.Error())))
		return
	}

	fmt.Println("user", jwtParsed.UserID)
	// fmt.Println(Authorization, jwtParsed)
	Req := userservice.ProfileRequest{UserID: jwtParsed.UserID}

	mySqlRepo := mysql.New()

	userSvc := userservice.NewService(mySqlRepo, authSVC)

	res, errP := userSvc.Profile(Req)

	if errP != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error on profile":"%s"}`, err.Error())))
		return
	}

	data, errJSON := json.Marshal(res)

	if errJSON != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error on json marshal":"%s"}`, err.Error())))
		return
	}

	writer.Write(data)
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `{"message":"everything is ok"}`)
}

// func parseJWT(tokenStr string) (*userservice.Claims, error) {

// 	fmt.Println("old", tokenStr)
// 	tokenStr = strings.Replace(tokenStr, "bearer ", "", 1)
// 	fmt.Println("new", tokenStr)

// 	token, err := jwt.ParseWithClaims(tokenStr, &userservice.Claims{}, func(token *jwt.Token) (interface{}, error) {
// 		fmt.Println("p1")
// 		return []byte(JWT_SECRET), nil
// 	})

// 	if token == nil {
// 		fmt.Println("p2")
// 		return nil, err
// 	}

// 	if claims, ok := token.Claims.(*userservice.Claims); ok && token.Valid {
// 		fmt.Printf("userID %v ExpiresAt %v\n", claims.UserID, claims.RegisteredClaims.ExpiresAt)

// 		return claims, nil
// 	} else {
// 		fmt.Println("p3")
// 		return nil, err
// 	}

// }
