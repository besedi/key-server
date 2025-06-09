package srv

import (
	"crypto/rand"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

const maxLen = 100_000

func Serve(size string, port string) {
	fmt.Println("Starting server..")
	http.HandleFunc("GET /key/{len}", handlerGenerate)
	http.HandleFunc("GET /key/", handlerDefault)
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))
}

func handlerDefault(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/key/1024", http.StatusFound)
}
func handlerGenerate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	len := r.PathValue("len")
	l, err := strconv.Atoi(len)

	if err != nil || l <= 0 || l >= maxLen {
		http.Error(w, "Invalid length", http.StatusBadRequest)
		return
	}

	b := make([]byte, l)
	rand.Read(b)
	for _, bin := range b {
		fmt.Fprintf(w, "%08b", bin)
	}
}
