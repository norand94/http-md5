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

	fmt.Println(r)
	fmt.Println(r.PostFormValue("request"))
	request := r.PostFormValue("request")
	if request == "" {
		http.Error(w, "Empty Request", http.StatusBadRequest)
		return
	}

	message, err := core.ReadJSON(request)
	if err != nil {
		http.Error(w, "Not valid JSON", http.StatusBadRequest)
		return
	}

	err = core.ValidateMessage(message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hash := core.GetHash(message)
	w.WriteHeader(200)
	fmt.Fprintln(w, hash)

}

func main()  {
	fmt.Println("Server started")
	http.HandleFunc("/md5", connect)
	http.ListenAndServe(":8888", nil)
}