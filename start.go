package main

import (
	"net/http"
	"fmt"
	"http-md5/core"
)

func connect(w http.ResponseWriter, r *http.Request){
	defer func() {
		if r := recover(); r != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}()

	request := r.PostFormValue("request")
	if request == "" {
		http.Error(w, "Empty Request", http.StatusBadRequest)
		return
	}
	message := core.ReadJSON(request)

	err := core.ValidateMessage(message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hash := core.GetHash(message)
	w.WriteHeader(200)
	fmt.Fprint(w, hash)

}

func main()  {
	fmt.Println("Server started")
	http.HandleFunc("/", connect)
	http.ListenAndServe(":8888", nil)
}