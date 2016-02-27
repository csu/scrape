package main

import (
  "fmt"
  "log"
  "os"
  "strings"

  "github.com/codegangsta/cli"
  "github.com/PuerkitoBio/goquery"
)

func main() {
  app := cli.NewApp()
  app.Name = "scrape"
  app.Usage = "scrapes a web page for a given selector"
  app.UsageText = "scrape [url] [selector]"
  app.Author = "Christopher Su"
  app.Email = "code@christopher.su"
  app.Action = func(c *cli.Context) {
    argCount := len(c.Args())
    if argCount < 2 {
      log.Fatal("Need at least 2 arguments.")
    }
    doc, err := goquery.NewDocument(c.Args()[0]) 
    if err != nil {
      log.Fatal(err)
    }
    if argCount == 2 {
      doc.Find(c.Args()[1]).Each(func(i int, s *goquery.Selection) {
        fmt.Println(s.Text())
      })
    } else {
      res := [][]string{}
      for i := 1; i < argCount; i++ {
        doc.Find(c.Args()[i]).Each(func(j int, s *goquery.Selection) {
          if j+1 > len(res) {
            res = append(res, []string{})
          }
          res[j] = append(res[j], s.Text())
        })
      }

      for _, elem := range res {
        for _, item := range elem {
          fmt.Println(strings.TrimSpace(item))
        }
        fmt.Println("==========")
      }
    }
  }
  app.Run(os.Args)
}

