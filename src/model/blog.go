package model

import (
	"gorm.io/gorm"
	"time"
)

type Article struct {
	gorm.Model  `gorm:"embedded" swaggerignore:"true"`
	Controller  UserInfoModel `gorm:"embedded" swaggerignore:"true"`
	Title       string        `gorm:"comment:title;Not Null;index:idx_article;"`
	Content     string        `gorm:"comment:content"`
	Description string        `gorm:"comment:description"`
	Author      string        `gorm:"comment:author;index:idx_author;"`
	MetaData    string        `gorm:"comment:metaData"`
	Status      bool          `gorm:"comment:status;"`
	Permission  string        `gorm:"comment:permission"` // 是否仅自己可见之类的权限
	Md5         string        `gorm:"comment:md5;index:idx_article;"`
	Category    []Category    `gorm:"many2many:category_articles;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" swaggerignore:"true"`
	Tag         []Tag         `gorm:"many2many:article_tags;constraint:OnUpdate:CASCADE;OnDelete:CASCADE;" swaggerignore:"true"`
	Background  string        `gorm:"comment:background"`
	ReadCount   int           `gorm:"comment:read_count;"`
}

type Category struct {
	gorm.Model    `gorm:"embedded"`
	TopCategoryID uint   `gorm:"foreignKey:TopCategoryID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Name          string `gorm:"comment:name;Not Null;"`
}

type Tag struct {
	gorm.Model `gorm:"embedded"`
	Name       string `gorm:"comment:name;Not Null;"`
}

type Comment struct {
	gorm.Model   `gorm:"embedded"`
	Controller   UserInfoModel `gorm:"embedded"`
	ArticleID    uint          `gorm:"comment:article_id;foreignKey:ArticleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	IsTop        bool          `gorm:"comment:is_top;"`
	Comment      string        `gorm:"comment:comment;"`
	TopCommentID uint          `gorm:"comment:top_comment;foreignKey:CommentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Todo struct {
	gorm.Model `gorm:"embedded"`
	Content    string `gorm:"comment:content;Not Null;"`
}

type TimeLine struct {
	IDModel IDModel `gorm:"embedded"`
	Content string  `gorm:"comment:content;Not Null;"`
}

// AfterFind 钩子函数，在查询文章后刷新阅读次数(每五秒允许增加一次)
func (ac *Article) AfterFind(tx *gorm.DB) (err error) {
	nowTime := time.Now()
	if nowTime.Sub(ac.UpdatedAt).Seconds() > 5 {
		ac.ReadCount += 1
		tx.Save(&ac)
	}
	return
}
