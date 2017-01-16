package main

import (
	"net/http"
	"fmt"
	"http-md5/core"
)

func connect(w http.ResponseWriter, r *http.Request){
	//fmt.Println(r)
	w.Header().Set("Access-Control-Allow-Origin", "*")

	request := r.PostFormValue("request")
	message := core.ReadJSON(request)
	hash := core.GetHash(message)
	fmt.Fprint(w, hash)
}

func main()  {
	fmt.Println("Server started")
	http.HandleFunc("/", connect)
	http.ListenAndServe(":8888", nil)
}