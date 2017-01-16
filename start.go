package main

import (
	"net/http"
	"fmt"
)

func connect(w http.ResponseWriter, r *http.Request){
	//fmt.Println(r)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprint(w, "OK")

}

func main()  {
	fmt.Println("Server started")
	http.HandleFunc("/", connect)
	http.ListenAndServe(":8888", nil)
}