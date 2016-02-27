package main

import (
  "fmt"
  "log"
  "os"

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
    if len(c.Args()) != 2 {
      log.Fatal("Need 2 arguments.")
    }
    doc, err := goquery.NewDocument(c.Args()[0]) 
    if err != nil {
      log.Fatal(err)
    }
    doc.Find(c.Args()[1]).Each(func(i int, s *goquery.Selection) {
      fmt.Println(s.Text())
    })
  }
  app.Run(os.Args)
}

