package main

import (
	"bufio"
	"bytes"
	"flag"
	"io/ioutil"
	"log"
	"os"

	"github.com/pkg/profile"
	m "github.com/veer66/mapkha"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var dixPath string

func init() {
	flag.StringVar(&dixPath, "dix", "", "Dictionary path")
}

func main() {
	flag.Parse()
	s := profile.Start(profile.CPUProfile, profile.ProfilePath("."))
	defer s.Stop()

	var dict m.PrefixTree
	var e error
	if dixPath == "" {
		dict, e = m.LoadDefaultDict()
	} else {
		dict, e = m.LoadDict(dixPath)
	}
	check(e)
	m.MakeEdgeBuilders(dict)
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal("could not read input:", err)
	}
	scanner := bufio.NewScanner(bytes.NewReader(b))
	outbuf := bufio.NewWriter(os.Stdout)
	for scanner.Scan() {
		m.Segment(outbuf, scanner.Text())
	}
	outbuf.Flush()
}
