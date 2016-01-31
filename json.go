package main
import (
	"net/http"
	"encoding/json"
	"encoding/xml"
	"path"
	"html/template"
)

type ProfileJson struct {
	Name    string
	Hobbies []string
}
type ProfileXML struct {
	Name    string
	Hobbies []string `xml:"Hobbies>Hobby"`
}

func main() {
	http.HandleFunc("/json", jsonHandler)
	http.HandleFunc("/xml", xmlHandler)
	http.HandleFunc("/png", pngHandler)
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}

func xmlHandler(w http.ResponseWriter, r *http.Request) {
	profile := &ProfileXML{Name:"Alex", Hobbies:[]string{"snowboarding", "programming"}}
	xml, err := xml.MarshalIndent(profile, "", "	")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/xml")
	w.Write(xml)
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	profile := &ProfileJson{Name:"Alex", Hobbies:[]string{"snowboarding", "programming"}}
	js, err := json.Marshal(profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func pngHandler(w http.ResponseWriter, r *http.Request) {
	fp := path.Join("", "mario.png")
	http.ServeFile(w, r, fp)
}


func indexHandler(w http.ResponseWriter, r *http.Request) {
	profile := &ProfileJson{Name:"Alex", Hobbies:[]string{"snowboarding", "programming"}}

	fp := path.Join("assert","index.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, profile); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}


}
