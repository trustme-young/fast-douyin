package service

type AuthorProto struct {
	Id            int64
	Name          string
	FollowCount   int64
	FollowerCount int64
	IsFollow      bool
}
type VideoProto struct {
	Id             int64
	Author         AuthorProto
	PlayUrl        string
	CoverUrl       string
	CommentCount   int64
	FavouriteCount int64
	Title          string
	IsFavourite    bool
}
type CommentProto struct {
	Id         int64       `json:"id,omitempty"`
	User       AuthorProto `json:"user"`
	Content    string      `json:"content,omitempty"`
	CreateDate string      `json:"create_date,omitempty"`
}
