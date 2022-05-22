// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Comment struct {
	SessionID string `json:"sessionId" bson:"sessionId"`
	Author    string `json:"author" bson:"author"`
	Data      string `json:"data" bson:"data"`
}

type NewComment struct {
	SessionID string `json:"sessionId" bson:"sessionId"`
	Author    string `json:"author" bson:"author"`
	Data      string `json:"data" bson:"data"`
}

type NewPlayer struct {
	SessionID string `json:"sessionId" bson:"sessionId"`
	UserID    string `json:"userId" bson:"userId"`
}

type NewSession struct {
	Name string `json:"name" bson:"name"`
	Host string `json:"host" bson:"host"`
}

type Session struct {
	ID       string     `json:"_id" bson:"_id"`
	Name     string     `json:"name" bson:"name"`
	Ongoing  bool       `json:"ongoing" bson:"ongoing"`
	Players  []string   `json:"players" bson:"players"`
	Comments []*Comment `json:"comments" bson:"comments"`
}
