package main

import (
	"suggestApp/config"
	"suggestApp/delivery/httpserver"
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

	cfg := config.Config{
		HttpServer: config.HttpServer{Port: 8088},
		Auth: authservice.Config{
			SignKey:               JWT_SECRET,
			AccessExpiretionTime:  Access_tokenTime,
			RefreshExpiretionTime: Refresh_tokenTime,
			AccessSubject:         Access_token,
			RefreshSubject:        Refresh_token,
		},
		Mysql: mysql.Config{
			Username: "gameapp",
			Password: "gameappt0lk2o20",
			Port:     3308,
			Host:     "localhost",
			DbName:   "gameapp_db",
		},
	}

	authSVC, userSVC := setupServices(cfg)

	server := httpserver.New(cfg, authSVC, userSVC)

	server.Serve()

	// mux.HandleFunc("/health-check", healthCheckHandler)
	// mux.HandleFunc("/users/register", userRegisterHandler)
	// mux.HandleFunc("/users/login", userLoginHandler)
	// mux.HandleFunc("/users/profile", userProfileHandler)

}

// func userLoginHandler(writer http.ResponseWriter, req *http.Request) {
// 	if req.Method != http.MethodPost {
// 		http.Error(writer, `{"message":"Method not allowed"}`, http.StatusMethodNotAllowed)
// 		return
// 	}
// 	authSVC := authservice.New(JWT_SECRET, Access_token, Refresh_token, Access_tokenTime, Refresh_tokenTime)

// 	data, err := io.ReadAll(req.Body)
// 	if err != nil {
// 		writer.Write([]byte(fmt.Sprintf(`{"error on read":"%s"}`, err.Error())))
// 		return
// 	}

// 	var UReq userservice.LoginRequest

// 	err = json.Unmarshal(data, &UReq)
// 	if err != nil {
// 		writer.Write([]byte(fmt.Sprintf(`{"error on json unmarshal":"%s"}`, err.Error())))
// 		return
// 	}

// 	mySqlRepo := mysql.New()

// 	userSvc := userservice.NewService(mySqlRepo, authSVC)

// 	Resp, err := userSvc.Login(UReq)
// 	if err != nil {
// 		writer.Write([]byte(fmt.Sprintf(`{"error on login":"%s"}`, err.Error())))
// 		return
// 	}
// 	data, err = json.Marshal(Resp)
// 	if err != nil {
// 		writer.Write([]byte(
// 			fmt.Sprintf(`{"error": "%s"}`, err.Error()),
// 		))

// 		return
// 	}

// 	writer.Write(data)

// 	return
// }

// func userProfileHandler(writer http.ResponseWriter, req *http.Request) {
// 	if req.Method != http.MethodGet {
// 		http.Error(writer, `{"message":"Method not allowed"}`, http.StatusMethodNotAllowed)
// 		return
// 	}

// 	authSVC := authservice.New(JWT_SECRET, Access_token, Refresh_token, Access_tokenTime, Refresh_tokenTime)

// 	Authorization := req.Header.Get("Authorization")

// 	Authorization = strings.Replace(Authorization, "Bearer ", "", 1)
// 	jwtParsed, err := authSVC.ParseToken(Authorization)
// 	if err != nil {
// 		writer.Write([]byte(fmt.Sprintf(`{"error on parseJWT":"%s"}`, err.Error())))
// 		return
// 	}

// 	fmt.Println("user", jwtParsed.UserID)
// 	// fmt.Println(Authorization, jwtParsed)
// 	Req := userservice.ProfileRequest{UserID: jwtParsed.UserID}

// 	mySqlRepo := mysql.New()

// 	userSvc := userservice.NewService(mySqlRepo, authSVC)

// 	res, errP := userSvc.Profile(Req)

// 	if errP != nil {
// 		writer.Write([]byte(fmt.Sprintf(`{"error on profile":"%s"}`, err.Error())))
// 		return
// 	}

// 	data, errJSON := json.Marshal(res)

// 	if errJSON != nil {
// 		writer.Write([]byte(fmt.Sprintf(`{"error on json marshal":"%s"}`, err.Error())))
// 		return
// 	}

// 	writer.Write(data)
// }

func setupServices(cfg config.Config) (authservice.Service, userservice.Service) {

	authSVC := authservice.New(cfg.Auth)

	MysqlRepo := mysql.New(cfg.Mysql)

	userSVC := userservice.NewService(MysqlRepo, authSVC)

	return authSVC, userSVC

}
