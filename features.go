package main

import (
	"fmt"
	"strings"
)

type EdgeFeatures struct {
	AFollowingCount      int
	BFollowingCount      int
	commonFollowingCount int
	commonFollowingRate  float64
	isFollowing          int
}

func (el EdgeFeatures) StringSlices() []string {
	return strings.Split(fmt.Sprint(el.AFollowingCount, el.BFollowingCount, el.commonFollowingCount, el.commonFollowingRate), " ")
}

func (el EdgeFeatures) CSVString() string {
	//return fmt.Sprintf("%d,%d,%d,%f,%d", el.AFollowingCount, el.BFollowingCount, el.commonFollowingCount, el.commonFollowingRate, el.isFollowing)
	return fmt.Sprintf("%d,%d,%d", el.AFollowingCount, el.BFollowingCount, el.isFollowing)
}
