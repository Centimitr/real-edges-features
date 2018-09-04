package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Edges struct {
	train     map[int][]int
	LinkCount int
	allIds    []int
}

func (e *Edges) ReadTrainSet(filename string) {
	txt, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Read csv", err)
	}
	lines := strings.Split(string(txt), "\r\n")
	e.train = make(map[int][]int, len(lines))
	linkCount := 0
	for _, l := range lines {
		if l == "" {
			continue
		}
		values := strings.Split(l, "\t")
		cur, err := strconv.Atoi(values[0])
		if err != nil {
			fmt.Println("Read csv", err)
		}
		lc := len(values) - 1
		linkCount += lc
		followings := make([]int, lc)
		for i, v := range values[1:] {
			following, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println("Read csv", err)
			}
			followings[i] = following
		}
		e.train[cur] = followings
	}
	e.LinkCount = linkCount
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

func (e *Edges) isFollowing(a, b int) bool {
	for _, f := range e.followings(a) {
		if f == b {
			return true
		}
	}
	return false
}

func (e *Edges) commonFollowingCounts(a, b int) (cnt int, total int) {
	aflen, bflen := len(e.followings(a)), len(e.followings(b))
	if aflen < bflen {
		a, b = b, a
	}
	for _, va := range e.followings(a) {
		for _, vb := range e.followings(b) {
			if va == vb {
				cnt ++
			}
		}
	}
	return cnt, aflen + bflen
}

func (e *Edges) commonFollowersCount(a, b int) int {
	return 0
}
