package main

import  (
	"net/http"
	"os"
)
// my limits are the ammount of data i can store on the server (10GB)
// the badwidth of the server (1GB/s)
// the ammount of memeory I can use (1GB)

const (
	PORT = ":8080"
)

func blowup_if_present(err error) {
	if err != nil { panic(err); }
}

func main() {
	var mux *http.ServeMux;
	var err error;
	var b []byte;

	b, err = os.ReadFile("index.html");
	blowup_if_present(err);

	mux = http.NewServeMux();

	mux.Handle("/audio/",
	http.StripPrefix("/audio/", http.FileServer(http.Dir("audio"))));

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w,
			"Method not allowed", http.StatusMethodNotAllowed);
		}

		_, err = w.Write(b);
		if err != nil {
			http.Error(w,
			"Internal Server Error", http.StatusInternalServerError);
		}
	});

	print("server running on port", PORT);
	err = http.ListenAndServe(PORT, mux);
	blowup_if_present(err);
}
