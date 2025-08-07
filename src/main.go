package main

import  (
  "net/http"
	"html/template"
	"os"
	"log"
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

type Keymap struct {
	Names       []string // most keymaps have mulitple keys
	Description string
}

func main() {
	var b []byte;
	var err error;
	var f *os.File;
	var fi os.FileInfo;
	var mux *http.ServeMux;
	var sections []Section;
	var t *template.Template;

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
	Words: []string{ "جُعِلَ", "خُلِقَ", "ذُكِرَ", "سُءِلَ", "ظُلِمَ",
	"مُنِعَ", "لُعِنَ", "هُدِيَ", "وُجِدَ", "سُقِطَ", "عُفِيَ", "حُشِرَ", "حُبُكِ",
	"صُحُفِ", "اُذُنِ", "رُبُعُ", "اَعِضُ", "تَزِرُ", "يَرِتُ", "يَعِدُ" }});
	sections = append(sections,
	Section{Name: "14", Words: []string{"هُوَ", "لَكَ", "هِيَ",
	"بِكَ", "تَرَ", "لِمَ", "مَعَ", "خَلَقَ", "وَقَبَ", "حَسَدَ", "كَسَبَ",
	"قَدَرَ", "وَهَبَ", "صَدَقَ", "شَرَحَ", "فَرَضَ", "وَلَدَ", "قُدِرَ", "هُدِيَ",
	"نَفِخَ", "خُلِقَ", "غُفِرَ", "ضُرِبَ", "وُعِدَ", "كُتِبَ", "يَضَعُ", "قُضِيَ",
	"رَضِيَ", "سَمِعَ", "نَسِيَ", "بَخِلَ", "غَشِيَ", "لَبِتَ", "وَلِيَ", "اَقِيَ",
	"يَدِيَ", "تَبِعَ", "اَهُوَ", "اُخَرَ", "عُقَدِ", "اُفُقِ", "رُسُلُ", "وَجَدَكَ",
	"عَدَلَكَ", "خَلَقَكَ", "بِيَدِكَ" }});
	sections = append(sections, Section{Name: "15", Vowel: "ََ", // all letter have alif added which is not correct 
	Letters: ARA_ALPHA, Words: []string{}});                    // e.g. "alif double fatha alif"

	t, err = template.New("index.tmpl").Funcs(template.FuncMap{
		"Add": func (a, b int) int { return a + b },
		"Sub": func (a, b int) int { return a - b },
		"Compare": func (a , b string) bool { 
			var l int;
			var ok bool;

			if l = len(a); l < len(b) {
				l = len(b);
			}

			ok = true;
			for i := range a { 
				if a[i] != b[i] { ok=false;break; }
			}
			return ok;
		},
	}).ParseFiles("src/tmpl/index.tmpl");
	if err != nil { log.Fatal(err); }

	if _, err = os.Stat("src/html/"); err != nil {
		err = os.MkdirAll("src/html/", 0777);
		blowup_if_present(err);
	}

	f, err = os.Create("src/html/index.html");
	blowup_if_present(err);
	defer f.Close();

	err = t.Execute(f,
	struct{
		Sections []Section;
		Keymaps []Keymap;
	}{
	Sections: sections,
	Keymaps:  []Keymap{
		{Names: []string{"ArrowDown", "j"},
		Description: "Move letter selection down"},
		{Names: []string{"ArrowUp", "k"},
		Description: "Move letter selection up"},
		{Names: []string{"ArrowLeft", "l"},
		Description: "Move letter selection left"},
		{Names: []string{"ArrowRight", "h"},
		Description: "Move letter selection right"},
		{Names: []string{"n"},
		Description: "Cycle forward a page"},
		{Names: []string{"b"},
		Description: "Cycle back a page"},
	}});
	blowup_if_present(err);

	_, err = f.Seek(0, 0);
	blowup_if_present(err);
	fi, err = f.Stat();
	blowup_if_present(err);

	b = make([]byte, fi.Size());
	_, err = f.Read(b);
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
