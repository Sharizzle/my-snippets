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
	json.NewEncoder(w).Encode(codes)
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
	code.Date = r.FormValue("date")
	code.Title = r.FormValue("title")
	code.Description = r.FormValue("description")
	code.Category = r.FormValue("category")
	code.Content = r.FormValue("content")

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
	db.DB.Model(&code).Update("date", r.FormValue("date"))
	db.DB.Model(&code).Update("title", r.FormValue("title"))
	db.DB.Model(&code).Update("description", r.FormValue("description"))
	db.DB.Model(&code).Update("category", r.FormValue("category"))
	db.DB.Model(&code).Update("content", r.FormValue("content"))

	json.NewEncoder(w).Encode(&code)
}