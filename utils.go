package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strconv"
)

func IntsToStrings(a []int) []string {
	r := make([]string, len(a))
	for i, v := range a {
		r[i] = strconv.Itoa(v)
	}
	return r
}

func StringsToInts(a []string) []int {
	r := make([]int, len(a))
	for i, s := range a {
		v, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println("StringsToInts Error:", err)
		}
		r[i] = v
	}
	return r
}

func BoolToInt(v bool, t int, f int) int {
	if v {
		return t
	}
	return f
}

func ProcessCSV(rc io.Reader) (ch chan []string) {
	ch = make(chan []string, 10)
	go func() {
		r := csv.NewReader(rc)
		r.Comma = '	'
		if _, err := r.Read(); err != nil {
			log.Fatal(err)
		}
		defer close(ch)
		for {
			rec, err := r.Read()
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Fatal(err)
			}
			ch <- rec
		}
	}()
	return
}
