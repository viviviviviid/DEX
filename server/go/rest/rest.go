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
	"github.com/m/go/sub"
	"github.com/m/go/utils"
)

var (
	authReq  = model.AuthReq{}
	userInfo = model.User{}
	// orderReq = model.OrderReq{}
	orders = []model.Order{}
)

func documentation(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		json.NewEncoder(rw).Encode("Welcome to the API.")
	}
}

func authenticate(res http.ResponseWriter, req *http.Request) {
	fmt.Println("/api/auth")
	body, err := io.ReadAll(req.Body)
	utils.HandleErr(err)
	err = json.Unmarshal(body, &authReq)
	utils.HandleErr(err)

	userInfo, err := db.GetUserInfoDB(authReq.WalletAddress)
	if err != nil {
		if err == sql.ErrNoRows {
			// DB에 정보가 없을 때
			authRes := model.AuthRes{
				WalletAddress: "",
				IsNewUser:     true,
				IsVerified:    false,
				IsSignature:   false,
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
		IsSignature:   true,
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(authRes)
}

func register(res http.ResponseWriter, req *http.Request) {
	fmt.Println("/api/register")
	body, err := io.ReadAll(req.Body)
	utils.HandleErr(err)
	err = json.Unmarshal(body, &userInfo)
	utils.HandleErr(err)
	db.SaveUserDB(userInfo)

	if err != nil {
		utils.HandleErr(err)
		res.WriteHeader(http.StatusInternalServerError)
	}
	res.WriteHeader(http.StatusCreated) // 201 Created

}

func sign(res http.ResponseWriter, req *http.Request) {
	fmt.Println("/api/sign")
	body, err := io.ReadAll(req.Body)
	utils.HandleErr(err)
	var orderReq model.OrderReq
	err = json.Unmarshal(body, &orderReq)
	utils.HandleErr(err)

	// Signature 비교
	if !db.CheckSignatureDB(orderReq.WalletAddress, orderReq.Signature) {
		sub.SigInvalidRes(res)
		return
	}

	sub.SigValidRes(res)
}

func order(res http.ResponseWriter, req *http.Request) {
	fmt.Println("/api/order")
	body, err := io.ReadAll(req.Body)
	utils.HandleErr(err)
	var orderReq model.OrderReq
	err = json.Unmarshal(body, &orderReq)
	utils.HandleErr(err)

	newOrder, err := db.SaveOrderDB(orderReq)
	utils.HandleErr(err)

	// 매수 주문인 경우 매도 주문과 매칭
	if orderReq.IsBuy {
		orders, err := db.GetSellOrder4MatchDB(orderReq)
		utils.HandleErr(err)

		fmt.Println(orders)

		// 매칭된 오더가 있으면 실행
		if len(orders) > 0 {
			sub.ExecuteBuyOrders(newOrder, orders)
		}
	} else {
		// 매도 주문인 경우 매수 주문과 매칭 (이 부분은 다른 함수에서 구현 예정)
		// orders, err = db.fetchMatchingBuyOrders(orderReq)
		// utils.HandleErr(err)
	}

	if len(orders) != 0 { // 매칭할 오더북이 있다면
		fmt.Println("Order processed")
		// 주문 처리 로직을 여기에 추가
	}

	// 성공적으로 주문 처리 완료 응답
	fmt.Fprintf(res, "Order processed successfully")
}

func Start() {
	fmt.Println("Starting REST API server on :5555")

	router := mux.NewRouter()
	router.HandleFunc("/", documentation).Methods("GET")
	router.HandleFunc("/api/auth", authenticate).Methods("POST")
	router.HandleFunc("/api/register", register).Methods("POST")
	router.HandleFunc("/api/sign", sign).Methods("POST")
	router.HandleFunc("/api/order", order).Methods("POST")

	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With", "Authorization"}),
	)

	err := http.ListenAndServe(":5555", corsHandler(router))
	if err != nil {
		fmt.Println("Failed to start server:", err)
		os.Exit(1)
	}
}
