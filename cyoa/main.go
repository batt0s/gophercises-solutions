package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

type Arc struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

var Arcs map[string]Arc

func main() {

	Arcs = loadArcs()
	log.Println("Arcs loaded.")

	mux := http.NewServeMux()
	mux.HandleFunc("/", ArcHanlder)
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8000", mux))
}

// GET /:arc
func ArcHanlder(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		path := req.URL.Path
		if path == "/" {
			http.Redirect(w, req, "/intro", http.StatusSeeOther)
		}
		ps := strings.Split(path, "/")
		arc := ps[len(ps)-1]
		arcObj := Arcs[arc]
		tmpl, err := template.ParseFiles("arc.html")
		if err != nil {
			fmt.Fprintf(w, "Failed to parse template\nError : %s", err.Error())
			return
		}
		tmpl.Execute(w, map[string]interface{}{
			"Arc":    arc,
			"ArcObj": arcObj,
		})
	default:
		fmt.Fprint(w, "Not found.")
	}
}

func loadArcs() map[string]Arc {
	data, err := os.ReadFile("gopher.json")
	if err != nil {
		log.Fatalf("Can't read gopher.json\nError : %s", err.Error())
	}
	AllArcs := make(map[string]Arc)
	if err = json.Unmarshal(data, &AllArcs); err != nil {
		log.Fatalf("Can't unmarshal data\nError : %s", err.Error())
	}
	return AllArcs
}
