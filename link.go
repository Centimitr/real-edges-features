package main

type EdgeLink struct {
	A int
	B int
}

func (el *EdgeLink) features(e *Edges) *EdgeFeatures {
	a, b := el.A, el.B
	cFollowing, total := e.commonFollowingCounts(a, b)
	efs := &EdgeFeatures{
		AFollowingCount:      e.followingCount(a),
		BFollowingCount:      e.followingCount(b),
		commonFollowingCount: cFollowing,
		commonFollowingRate:  float64(cFollowing) / float64(total),
		isFollowing:          BoolToInt(e.isFollowing(b, a), 1, 0),
	}
	return efs
}
