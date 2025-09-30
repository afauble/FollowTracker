package models

type Following struct {
	RelationshipsFollowing []Follower `json:"relationships_following"`
}
