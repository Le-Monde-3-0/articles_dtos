package articles_dtos

import "github.com/lib/pq"

type ArticleInput struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type ArticleResponse struct {
	Id      int32
	UserId  int32
	Title   string
	Content string
	Likes   pq.Int32Array `gorm:"type:integer[]"`
}
