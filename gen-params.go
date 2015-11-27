// +build ignore

package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"golang.org/x/net/html"
)

func main() {
	flag.Parse()
	fname := flag.Arg(0)
	log.Printf("fname=%q\n", fname)

	f, err := os.Open(fname)
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}
	defer f.Close()

	z := html.NewTokenizer(f)

	depth := 0
loop:
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			if z.Err() == io.EOF {
				break loop
			}
			log.Fatalf("error: %v\n", z.Err())
		case html.TextToken:
			if depth > -1 {
				// emitBytes should copy the []byte it receives,
				// if it doesn't process it immediately.
				emitBytes(depth, z.Text())
			}
		case html.StartTagToken, html.EndTagToken:
			tn, _ := z.TagName()
			if len(tn) == 1 && tn[0] == 'a' {
				if tt == html.StartTagToken {
					depth++
				} else {
					depth--
				}
			}

		case html.CommentToken:
			log.Printf("comment: %q\n", z.Text())

		case html.DoctypeToken:
			log.Printf("doctype: %q\n", z.Text())

		case html.SelfClosingTagToken:
			// no-op

		default:
			log.Printf("token: %#v\n", tt)
		}
	}
}

func emitBytes(depth int, data []byte) {
	fmt.Printf("%3d >>> %q\n", depth, string(data))
}

func init() {
	log.SetPrefix("[gen-params] ")
	log.SetFlags(0)
}
