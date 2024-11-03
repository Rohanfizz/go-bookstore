package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseBody(r *http.Request,x interface{}){
	if body, err := io.ReadAll(r.Body); err == nil{
		if err := json.Unmarshal([]byte(body),x); err != nil{
			return
		}
	}
}

func SendOkResponse(w http.ResponseWriter, res []byte){
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}