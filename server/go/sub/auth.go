package sub

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func SigInvalidRes(res http.ResponseWriter) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusUnauthorized)
	response := map[string]string{"error": "Signature is different"}
	jsonResp, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	res.Write(jsonResp)
}

func SigValidRes(res http.ResponseWriter) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"message":       "Signature is valid. ",
		"okayToProceed": true,
	}
	jsonResp, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	res.Write(jsonResp)
}
