package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "github.com/LordCHTsai/repoFinder/finder"
)

var bb finder.Bb
var gr finder.Gr

func bbHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Access-Control-Allow-Origin", "*")
    kw := r.URL.Path[len("/BbSearch/"):]
    count, res, err := bb.Search(kw)
    if err == nil {
        b, _ := json.MarshalIndent(res, "", "    ")
        fmt.Fprintf(w, `{"total_count": %d, "items": %s}`, count, b)
    }
}

func grHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Access-Control-Allow-Origin", "*")
    kw := r.URL.Path[len("/GrSearch/"):]
    count, res, err := gr.Search(kw)
    if err == nil {
        b, _ := json.MarshalIndent(res, "", "  ")
        fmt.Fprintf(w, `{"total_count": %d, "items": %s}`, count, b)
    }
}

func main() {
    http.Handle("/", http.FileServer(http.Dir("./www/")))
    http.HandleFunc("/BbSearch/", bbHandler)
    http.HandleFunc("/GrSearch/", grHandler)
    http.ListenAndServe(":8080", nil)
}
