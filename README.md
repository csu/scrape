# scrape
Quickly scrape web pages from the command line. Outputs results to stdout. Essentially a cli wrapper for goquery.

## Installation
```bash
go install github.com/csu/scrape
```

## Usage
```bash
$ scrape http://christopher.su ".home-nav span"
resume
linkedin
v1
latest
archive
email
twitter
pgp
github
list
grid
```

```bash
$ scrape help
NAME:
   scrape - scrapes a web page for a given selector

USAGE:
   scrape [url] [selector]
```