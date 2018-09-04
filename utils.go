package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
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

func CSVToChan(rc io.Reader) (ch chan []string) {
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

func check(err error, s string) bool {
	if err != nil {
		fmt.Println(s+":", err)
		return true
	}
	return false
}

func SplitAndTrimSpace(s string, sep string) []string {
	vs := strings.Split(s, sep)
	for i, v := range vs {
		vs[i] = strings.TrimSpace(v)
	}
	return vs
}

func CombineMultipleCSV(dir string, out string) {
	var m map[string]map[string]string
	extractPairKey := func(vs []string) string {
		a, b := vs[1], vs[2]
		return a + "->" + b
	}
	getABFromPairKey := func(k string) (a, b string) {
		kvs := strings.Split(k, "->")
		a, b = kvs[0], kvs[1]
		return
	}

	// collect csv files and fill in map
	first := true
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(info.Name(), ".csv") {
			byts, err := ioutil.ReadFile(path)
			if check(err, "CombineMultipleCSV.ReadFile") {
				return err
			}

			lines := strings.Split(string(byts), "\r\n")
			if first {
				m = make(map[string]map[string]string, len(lines)-1)
				for _, line := range lines[1:] {
					vs := SplitAndTrimSpace(line, ",")
					if len(vs) < 3 {
						continue
					}
					m[extractPairKey(vs)] = make(map[string]string)
				}
				first = false
			}
			heading := SplitAndTrimSpace(lines[0], ",")
			featureNames := heading[3:]
			for _, line := range lines[1:] {
				vs := SplitAndTrimSpace(line, ",")
				if len(vs) < len(heading) {
					continue
				}
				k := extractPairKey(vs)
				vs = vs[3:]
				for i, name := range featureNames {
					m[k][name] = vs[i]
				}
			}
		}
		return err
	})

	// open output file
	f, err := os.OpenFile(out, os.O_RDWR|os.O_CREATE, 0755)
	if check(err, "CombineMultipleCSV.out") {
		return
	}
	defer f.Close()

	// write map to output file
	w := csv.NewWriter(f)
	first = true
	var allFeatureNames []string
	for k, features := range m {
		if first {
			allFeatureNames = make([]string, len(features))
			i := 0
			for name := range features {
				allFeatureNames[i] = name
				i++
			}
			sort.Strings(allFeatureNames)
			w.Write(append([]string{"source", "sink"}, allFeatureNames...))
			first = false
		}
		a, b := getABFromPairKey(k)
		values := make([]string, len(allFeatureNames))
		i := 0
		for _, name := range allFeatureNames {
			values[i] = features[name]
			i++
		}
		w.Write(append([]string{a, b}, values...))
	}
	w.Flush()
}
