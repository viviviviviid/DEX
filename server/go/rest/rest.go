package rest

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/m/go/db"
	"github.com/m/go/model"
	"github.com/m/go/utils"
)

var (
	authReq = model.AuthReq{}
)

// documentation provides basic API documentation
func documentation(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		json.NewEncoder(rw).Encode("Welcome to the API. Use /api/register to sign up and /api/login to sign in.")
	}
}

func authenticate(res http.ResponseWriter, req *http.Request) {
	fmt.Println("/api/auth")
	body, err := io.ReadAll(req.Body)
	if err != nil {
		utils.HandleErr(err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(body, &authReq)
	if err != nil {
		utils.HandleErr(err)
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	userInfo, err := db.GetUserInfo(authReq.WalletAddress)
	if err != nil {
		if err == sql.ErrNoRows {
			// DB에 정보가 없을 때
			authRes := model.AuthRes{
				WalletAddress: "",
				IsNewUser:     true,
				IsVerified:    false,
			}
			res.WriteHeader(http.StatusOK)
			json.NewEncoder(res).Encode(authRes)
			return
		}
		// 다른 에러가 있을 경우
		utils.HandleErr(err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	authRes := model.AuthRes{
		WalletAddress: userInfo.WalletAddress,
		IsNewUser:     false,
		IsVerified:    userInfo.HumanVerified,
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(authRes)
}

// // signUp handles user registration
// func signUp(res http.ResponseWriter, req *http.Request) {
// 	fmt.Println("Accessing /api/register")
// 	body, err := io.ReadAll(req.Body)
// 	utils.HandleErr(err)
// 	var user model.User
// 	err = json.Unmarshal(body, &user)
// 	utils.HandleErr(err)

// 	// Add user to the database
// 	err = db.SetUser(user.WalletAddress)
// 	utils.HandleErr(err)

// 	res.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(res).Encode(user)
// }

// // signIn handles user login
// func signIn(res http.ResponseWriter, req *http.Request) {
// 	fmt.Println("Accessing /api/login")
// 	body, err := io.ReadAll(req.Body)
// 	utils.HandleErr(err)
// 	var user model.User
// 	err = json.Unmarshal(body, &user)
// 	utils.HandleErr(err)

// 	// Check if user exists in the database
// 	userInfo, err := db.GetUserInfo(user.WalletAddress)
// 	if err == sql.ErrNoRows {
// 		// If user does not exist, redirect to signUp
// 		signUp(res, req)
// 	} else {
// 		utils.HandleErr(err)
// 		// Update HumanVerified status
// 		err = db.SetVerified(user.WalletAddress)
// 		utils.HandleErr(err)

// 		// Respond with successful login
// 		res.WriteHeader(http.StatusOK)
// 		json.NewEncoder(res).Encode(userInfo)
// 	}
// }

// Start initializes the HTTP server and listens on port 5555
func Start() {
	fmt.Println("Starting REST API server on :5555")

	router := mux.NewRouter()
	router.HandleFunc("/", documentation).Methods("GET")
	router.HandleFunc("/api/auth", authenticate).Methods("POST")

	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With", "Authorization"}),
	)

	err := http.ListenAndServe(":5555", corsHandler(router))
	if err != nil {
		fmt.Println("Failed to start server:", err)
		os.Exit(1) // Optionally exit if server fails to start
	}
}
