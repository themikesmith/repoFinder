package finder

import (
    "code.google.com/p/go.net/html"
    _ "fmt"
    "net/http"
    "strconv"
    "strings"
)

const BbUrl         = "https://bitbucket.org/"
const BbSearchUrl   = "repo/all/relevance/"
const maxPages      = 10
const useRealAvatar = false

type BbSearchRes struct {
    Title  string
    Url    string
    Date   string
    Lang   []string
    Avatar string
}

type Bb struct{}

/*
   It takes too much time to extract a real avatar url, so use it on your own risk
*/
func BbAvatar(username string) (string) {
    r, err := http.Get(BbUrl + username)
    if err != nil {
        return ""
    }
    defer r.Body.Close()
    z := html.NewTokenizer(r.Body)
    for {
        if z.Next() == html.ErrorToken {
            break
        }
        token := z.Token()
        if token.Data == "img" && token.Attr[1].Val == username {
            r, err := http.Get(token.Attr[0].Val)
            if err != nil || r.StatusCode != 200 {
                return ""
            }
            return strings.Split(token.Attr[0].Val, "\u0026s=96")[0]
        }
    }
    return ""
}

func (bb Bb) Search(kw string) ([]BbSearchRes, error) {
    var res []BbSearchRes
    pages := 1
    for i := 1; i <= pages; i++ {
        r, err := http.Get(BbUrl + BbSearchUrl + strconv.Itoa(i) + "?name=" + kw)
        if err != nil {
            return res, err
        }
        defer r.Body.Close()
        z := html.NewTokenizer(r.Body)
        for {
            if z.Next() == html.ErrorToken {
                break
            }
            token := z.Token()
            switch token.Type {
            case html.StartTagToken:
                switch token.Data {
                case "article":
                    if len(token.Attr) > 0 && token.Attr[0].Val == "repo-summary" {
                        var repo BbSearchRes
                        for {
                            z.Next()
                            child := z.Token()
                            if child.Type == html.EndTagToken && child.Data == "article" {
                                break
                            }
                            if child.Type == html.StartTagToken {
                                switch child.Data {
                                case "a":
                                    if len(child.Attr) == 2 && child.Attr[0].Val == "repo-link" {
                                        repo.Url = BbUrl + child.Attr[1].Val[1:len(child.Attr[1].Val)-1]
                                        z.Next()
                                        content := z.Token()
                                        repo.Title = content.Data
                                    }
                                case "img":
                                    username := strings.Split(child.Attr[0].Val, "/")[0]
                                    if useRealAvatar {
                                        repo.Avatar = BbAvatar(username)
                                    }
                                    if len(repo.Avatar) == 0 {
                                        repo.Avatar = child.Attr[1].Val
                                    }
                                    urlStr := strings.Split(child.Attr[1].Val, "/")
                                    lang := strings.Split(urlStr[len(urlStr)-1], "_")[0]
                                    repo.Lang = append(repo.Lang, lang)
                                case "time":
                                    repo.Date = child.Attr[0].Val
                                }
                            }
                        }
                        res = append(res, repo)
                    }
                    break
                case "a":
                    if len(token.Attr) > 0 && token.Attr[0].Key == "href" && strings.Contains(token.Attr[0].Val, BbSearchUrl) {
                        z.Next()
                        content := z.Token()
                        num, err := strconv.Atoi(content.Data)
                        if err == nil && num > pages && num <= maxPages {
                            pages = num
                        }
                    }
                    break
                }
            }
        }
    }
    return res, nil
}
