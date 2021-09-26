package table_struct

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	ArticleTags []Tag  `gorm:"many2many:article_tag;"`
	Title       string `gorm:"type:VARCHAR(75) NOT NULL;"`
	Content     string `gorm:"type:MEDIUMTEXT NOT NULL;"`
	Coords      string `gorm:"type:POINT NOT NULL;"`
}
type TagType struct {
	gorm.Model
	Tags []Tag
	Name string `gorm:"type:CHAR(24) NOT NULL;,uniqueIndex,sort:asc"`
}

type Tag struct {
	gorm.Model
	Name      string `gorm:"type:CHAR(24) NOT NULL;,uniqueIndex,sort:asc"`
	TagTypeID uint
}

