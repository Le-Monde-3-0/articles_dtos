package articles_dtos

import (
	"gorm.io/gorm"
	"time"
)

type ArticleInput struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type ArticleResponse struct {
	Id                    uint `gorm:"primarykey"`
	CreatedAt             time.Time
	UpdatedAt             time.Time
	DeletedAt             gorm.DeletedAt `gorm:"index"`
	UserId                int32
	AuthorName            string
	Title                 string
	Subtitle              string
	Content               string
	Topic                 string
	Draft                 bool
	HasConnectedUserLiked bool
	Likes                 []RecordLike `gorm:"foreignKey:ArticleID"`
	Views                 []RecordView `gorm:"foreignKey:ArticleID"`
}

type RecordView struct {
	ID        uint `gorm:"primaryKey"`
	ArticleID uint // Foreign key
	UserId    int32
	LikeTime  time.Time
}

type RecordLike struct {
	ID        uint `gorm:"primaryKey"`
	ArticleID uint // Foreign key
	UserId    int32
	LikeTime  time.Time
}