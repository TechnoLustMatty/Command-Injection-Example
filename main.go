package main

import (
	"fmt"
	"net/http"
	"os/exec"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		command := r.URL.Query().Get("command")
		output, err := exec.Command("ls", command).Output()
		if err != nil {
			http.Error(w, "Error executinh command", http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "Command output: %s", output)
	})
	http.ListenAndServe(":8080", nil)
}
