package handlers

import (
	"diplom/internal/storages"
	"encoding/json"
	"log"
	"net/http"
)

func ConnectionHandler(w http.ResponseWriter, _ *http.Request) {
	res := storages.GetResultData()
	resBytes, _ := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(resBytes)
	log.Println(res)

}
