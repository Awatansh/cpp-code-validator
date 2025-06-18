package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/judge/backend", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
			return
		}
		err := r.ParseMultipartForm(10 << 20)
		if err != nil {
			http.Error(w, "Error Parsing Form", http.StatusBadRequest)
			return
		}
		file, _, err := r.FormFile("code")
		if err != nil {
			http.Error(w, "Error retrieving file", http.StatusBadRequest)
			return
		}
		defer file.Close()
		content, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, "Error Reading file", http.StatusInternalServerError)
			return
		}
		fmt.Println("Recieved File")
		fmt.Println("Recieved code :")
		fmt.Println(string(content))
		verdict, argstring := spinner(string(content))
		if verdict {
			w.WriteHeader(http.StatusAccepted)
			w.Write([]byte("Testcases Passed Successfully"))
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(fmt.Sprintf("Testcases failed Error logs: %s", argstring)))
		}
	})
	fmt.Println("Server starting on 8000...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
