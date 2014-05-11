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
    res, err := bb.Search(kw)
    if err == nil {
        b, _ := json.MarshalIndent(res, "", "  ")
        fmt.Fprintf(w, "%s", b)
    }
}

func grHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Access-Control-Allow-Origin", "*")
    kw := r.URL.Path[len("/GrSearch/"):]
    res, err := gr.Search(kw)
    if err == nil {
        b, _ := json.MarshalIndent(res, "", "  ")
        fmt.Fprintf(w, "%s", b)
    }
}

func main() {
    http.HandleFunc("/BbSearch/", bbHandler)
    http.HandleFunc("/GrSearch/", grHandler)
    http.ListenAndServe(":8080", nil)
}
