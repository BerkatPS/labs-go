package helper

import (
	"encoding/json"
	"net/http"
)

// parameter pertama responsewriter
// parameter kedua http code
// parameter ketiga payload interface , artinya bisa
func Writejson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
