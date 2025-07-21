package main

import  (
	//"net/http"
	"html/template"
	"os"
)
// my limits are the ammount of data i can store on the server (10GB)
// the badwidth of the server (1GB/s)
// the ammount of memeory I can use (1GB)

const (
	PORT = ":8080"
)

var ARA_ALPHA = []string{ "ا", "ب", "ت", "ث", "ج", "ح", "خ", "د", "ذ", "ر", "ز", "س", "ش", "ص", "ض", "ط", "ظ", "ع", "غ", "ف", "ق", "ك", "ل", "م", "ن", "ه", "و", "ء", "ي", };

func blowup_if_present(err error) {
	if err != nil { panic(err); }
}

type Section struct {
	Words 	[]string
	Letters []string
	Vowel 	string
	Name 		string
}

func main() {
	var err error;
	// var mux *http.ServeMux;
	var t *template.Template;
	var sections []Section;

	t = template.Must(template.New("index.html").Funcs(template.FuncMap{
		"Add": func (a, b int) int { return a + b }, 
	}).ParseFiles("index.html"));

	sections = append(sections, Section{Letters: ARA_ALPHA, Name: "1", Vowel: ""});

	err = t.Execute(os.Stdout, sections);
	blowup_if_present(err);

	// mux = http.NewServeMux();

	// mux.Handle("/audio/",
	// 	http.StripPrefix("/audio/",
	// 		http.FileServer(http.Dir("audio"))));

	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	if r.Method != http.MethodGet {
	// 		http.Error(w,
	// 		"Method not allowed", http.StatusMethodNotAllowed);
	// 	}

	// 	t.Execute(w, nil);
	// 	if err != nil {
	// 		http.Error(w,
	// 		"Internal Server Error", http.StatusInternalServerError);
	// 	}
	// });

	// print("server running on port", PORT);
	// err = http.ListenAndServe(PORT, mux);
	// blowup_if_present(err);
}
