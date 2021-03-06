package code

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sharizzle/my-snippets/server/db"
	customHTTP "github.com/sharizzle/my-snippets/server/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	var codes []Code
	db.DB.Find(&codes)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Total-Count", "10")
	json.NewEncoder(w).Encode(codes)
}

func OptionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w)
}


func ShowHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var code Code
	db.DB.First(&code, params["codeId"])
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(code)
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	var code Code

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&code)	

	err := db.DB.Create(&code).Error
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: " + err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&code)
}


func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var code Code
	var codes []Code

	db.DB.First(&code, params["codeId"])
	db.DB.Delete(&code)

	db.DB.Find(&codes)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(codes)
}

func UpdateHandler(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	var code Code
	
	db.DB.First(&code, params["codeId"])
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&code)

	db.DB.Model(&code).Save(&code)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&code)


}