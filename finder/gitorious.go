package finder

import (
    "code.google.com/p/go.net/html"
    _ "fmt"
    "net/http"
    "strconv"
    "strings"
    "encoding/xml"
    "io/ioutil"
)

const GrUrl         = "https://gitorious.org/"
const GrSearchUrl   = "search"
const GrMaxPages      = 10

type GrSearchRes struct {
  Title         string
  Description   string
  Url           string
  Date          string
}

type GrXml struct {
  Date          string  `xml:"last-pushed-at"`
  CreateDate    string  `xml:"created-at"`
}

type Gr struct{}

/*
*/
func GrUpdateTime(reponame string) (string) {
  r, err := http.Get(GrUrl + reponame + ".xml")
  if err != nil {
    return ""
  }
  defer r.Body.Close()
  dxml := GrXml{}
  b, _ := ioutil.ReadAll(r.Body)
  xml.Unmarshal(b, &dxml)
  if len(dxml.Date) > 0 {
    return dxml.Date
  }
  return dxml.CreateDate
}


func (gr Gr) Search(kw string) (int, []GrSearchRes, error) {
  var res []GrSearchRes
  pages := 1
  total_count := 0
  for i := 1; i <= pages; i++ {
    r, err := http.Get(GrUrl + GrSearchUrl + "?page=" + strconv.Itoa(i) + "&q=" + kw)
    if err != nil {
      return total_count, res, err
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
          case "dt":
            var repo GrSearchRes
            for {
              z.Next()
              child := z.Token()
              if child.Type == html.EndTagToken && child.Data == "dd" {
                if len(repo.Url) > 0 {
                  res = append(res, repo)
                }
                break
              }
              if child.Type == html.StartTagToken {
                switch child.Data {
                  case "a":
                    slashes := strings.Split( child.Attr[0].Val, "/" )
                    if len(slashes) == 3 {
                      repo.Url   = GrUrl + child.Attr[0].Val[1:]
                      repo.Date  = GrUpdateTime(child.Attr[0].Val[1:])
                      z.Next()
                      title := z.Token()
                      repo.Title = title.Data
                    }
                    break
                  case "div":
                    if len(child.Attr) == 1 && child.Attr[0].Val == "muted" {
                      z.Next()
                      content := z.Token()
                      repo.Description = strings.Trim(strings.Replace(content.Data, "\n", " ", -1), " ")
                    }
                    break
                }
              }
            }
          case "a":
            if len(token.Attr) == 3 && token.Attr[0].Val == "next_page" {
              num := strings.Split( strings.Split( token.Attr[2].Val, "=" )[1], "&amp;q" )[0]
              realNum, err := strconv.Atoi(num)
              if err == nil && realNum <= GrMaxPages {
                pages = realNum
              }
            }
            break
          case "p":
            if len(token.Attr) == 1 && token.Attr[0].Val == "hint search_time" {
              for {
                z.Next()
                token := z.Token()
                if token.Type == html.StartTagToken && token.Data == "small" {
                  z.Next()
                  content := z.Token()
                  num, err := strconv.Atoi(strings.Split(strings.Trim(content.Data, " \t\n"), " ")[1])
                  if err == nil {
                    total_count = num
                  }
                  break
                }
              }
            }
            break
        }
      }
    }
  }
  return total_count, res, nil
}
