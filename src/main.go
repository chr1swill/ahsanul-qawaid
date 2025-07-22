package main

import  (
  "net/http"
	"html/template"
)

const (
	PORT = ":8080"
)

var ARA_ALPHA = []string{ "ا", "ب", "ت", "ث", "ج", "ح", "خ", "د", "ذ",
	"ر", "ز", "س", "ش", "ص", "ض", "ط", "ظ", "ع", "غ", "ف", "ق", "ك", "ل",
	"م", "ن", "و", "ه", "ء", "ي", };

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
	var mux *http.ServeMux;
	var t *template.Template;
	var sections []Section;

	sections = append(sections, 
		Section{Letters: ARA_ALPHA, Name: "1", Vowel: ""});
	sections = append(sections, 
		Section{Letters: ARA_ALPHA, Name: "11", Vowel: "َ",
    Words: []string{"اَمَرَ", "بَلَغَ", "ثَمَرَ", "جَمَعَ", "حَسَدَ", "ذَكَرَ",
    "رَفَعَ", "زَعَمَ", "سَرَقَ", "صَدَقَ", "ضَرَبَ", "ظَلَمَ", "عَدَلَ", "قَمَرَ",
    "كَسَبَ", "وَجَدَ"}});
  sections = append(sections, 
		Section{Letters: ARA_ALPHA, Name: "12", Vowel: "ِ",
    Words: []string{"اَذِنَ", "بَقِيَ", "حَمِدَ", "خَشِيَ", "سَخِرَ", "شَرِبَ",
    "عَجِبَ", "غَضِبَ", "كَذِبَ", "بَخِلَ", "فَلِمَ", "يَءِسَ", "اَبَتِ",
    "بَلَدِ", "مَلِكِ", "حَطَبِ", "اِبِلِ", "كِبَرِ", "عَلِمَ", "اِرَمَ"}});
  sections = append(sections,
    Section{Letters: ARA_ALPHA, Name: "13", Vowel: "ُ",
    Words: []string{ "جُعِلَ", "خُلِقَ", "ذُكِرَ", "سُءِلَ", "ظُلِمَ", "مُنِعَ", "لُعِنَ", "هُدِيَ", "وُجِدَ", "سُقِطَ", "عُفِيَ", "حُشِرَ", "حُبُكِ", "صُحُفِ", "اُذُنِ", "رُبُعُ", "اَعِضُ", "تَزِرُ", "يَرِتُ", "يَعِدُ" }});

	t = template.Must(template.New("index.html").Funcs(template.FuncMap{
		"Add": func (a, b int) int { return a + b }, 
	}).ParseFiles("src/html/index.html"));
  blowup_if_present(err);

	 mux = http.NewServeMux();

	 mux.Handle("/audio/",
	 	http.StripPrefix("/audio/",
	 		http.FileServer(http.Dir("audio"))));

	 mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	 	if r.Method != http.MethodGet {
	 		http.Error(w,
	 		"Method not allowed", http.StatusMethodNotAllowed);
	 	}

	 	err = t.Execute(w, sections);
	 	if err != nil {
	 		http.Error(w,
	 		"Internal Server Error", http.StatusInternalServerError);
	 	}
	 });

	 print("server running on port", PORT);
	 err = http.ListenAndServe(PORT, mux);
	 blowup_if_present(err);
}
