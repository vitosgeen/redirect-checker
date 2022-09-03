package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"redirect-checker/internal/service"
	"strings"
)

func HandlerCheckRedirect(w http.ResponseWriter, req *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	myURL := strings.Join(req.URL.Query()["url"], "")
	result := service.MakeRequest(myURL)
	fmt.Println(result)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
