package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {
	p("containers: ", s.Contains("test", "es"))
	p("Count: ", s.Count("test", "t"))
	p("HasPerfix", s.HasPrefix("test", "te"))
	p("HasSuffix", s.HasSuffix("test", "st"))
	p("Index", s.Index("test", "e"))
	p("join:", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:", s.Repeat("a", 5))
	p("replace:", s.Replace("foo", "o", "0", -1))
	p("ReplaceAll:", s.ReplaceAll("foo", "o", "0"))
	p("split: ", s.Split("a,b,c", ","))
	p("tolower: ", s.ToLower("TEST"))
	p("toupper: ", s.ToUpper("TEST"))
	p()
	p("len:", len("hello"))
	p("char:", "hello"[1])
}
