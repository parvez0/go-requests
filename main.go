package main

import (
	"encoding/json"
	"github.com/parvez0/go-requests/requests"
	"net/http"
)

var clog = requests.NewLogger()

func CreateServer() {
	http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		clog.Info("a request was made from - ", request.RemoteAddr, " - method - ", request.Method)
		writer.Header().Set("Content-Type", "application/json")
		resp := map[string]string{
			"message": "working",
			"method": request.Method,
		}
		js, _ := json.Marshal(resp)
		writer.Write(js)
		return
	})
	clog.Info("server is starting on : 5000")
	err := http.ListenAndServe(":5000", nil)
	if err != nil{
		clog.Panicf("failed start server - ", err)
	}
}

func main() {
	CreateServer()

}
