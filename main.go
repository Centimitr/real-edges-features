package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Edges struct {
	train map[int][]int
}

func (e *Edges) readTrainSet(filename string) {
	txt, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(txt), "\r\n")
	e.train = make(map[int][]int, len(lines))
	for _, l := range lines {
		if l == "" {
			continue
		}
		values := strings.Split(l, "\t")
		cur, err := strconv.Atoi(values[0])
		if err != nil {
			fmt.Println("Read csv", err)
		}
		followings := make([]int, len(values)-1)
		for i, v := range values[1:] {
			following, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println("Read csv", err)
			}
			followings[i] = following
		}
		e.train[cur] = followings
	}
}

func (e *Edges) followings(id int) []int {
	if v, ok := e.train[id]; ok {
		return v
	}
	return []int{}
}

func (e *Edges) followingCount(id int) int {
	return len(e.followings(id))
}

func main() {
	e := Edges{}
	e.readTrainSet("data/train.txt")
}
