package main

import  (
	"net/http"
)

const (
	PORT = ":8080"
)

func blowup(err error) {
	if err != nil { panic(err); }
}

func main() {
	var mux *http.ServeMux;
	var err error;

	mux = http.NewServeMux();

	print("server running on port", PORT);
	err = http.ListenAndServe(PORT, mux);
	blowup(err);
}
