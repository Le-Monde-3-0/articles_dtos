package articles_dtos

import "github.com/lib/pq"

type ArticleRequest struct {
	Id      int32
	UserId  int32
	Title   string
	Content string
	Likes   pq.Int32Array `gorm:"type:integer[]"`
}
