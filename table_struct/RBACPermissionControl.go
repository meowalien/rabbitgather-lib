package table_struct

import (
	"gorm.io/gorm"
	"time"
)

type Role struct {
	gorm.Model
	Users       []User
	Permissions []Permission `gorm:"many2many:role_permission;"`
	Title       string       `gorm:"type:VARCHAR(75) NOT NULL;"`
	Slug        string       `gorm:"type:VARCHAR(100) NOT NULL;,uniqueIndex,sort:asc"`
	Description string       `gorm:"type:TINYTEXT NULL"`
	Active      string       `gorm:"type:TINYINT(1) NOT NULL DEFAULT 0"`
	Content     string       `gorm:"type:TEXT NULL DEFAULT NULL"`
}

type Permission struct {
	gorm.Model
	Title       string `gorm:"type:VARCHAR(75) NOT NULL;"`
	Slug        string `gorm:"type:VARCHAR(100) NOT NULL;,uniqueIndex,sort:asc"`
	Description string `gorm:"type:TINYTEXT NULL;"`
	Active      string `gorm:"type:TINYINT(1) NOT NULL DEFAULT 0;"`
	Content     string `gorm:"type:TEXT NULL DEFAULT NULL;"`
}

type User struct {
	gorm.Model
	RoleID uint

	FirstName  string `gorm:"type:VARCHAR(50) DEFAULT NULL;"`
	MiddleName string `gorm:"type:VARCHAR(50) DEFAULT NULL;"`
	LastName   string `gorm:"type:VARCHAR(50) DEFAULT NULL;"`

	Mobile string `gorm:"type:VARCHAR(15) NULL;,uniqueIndex,sort:asc"`
	Email  string `gorm:"type:VARCHAR(50) NULL;,uniqueIndex,sort:asc"`

	PasswordHash string `gorm:"type:char(60) NOT NULL;"`
	PasswordSalt string `gorm:"type:char(24) NOT NULL;"`

	RegisteredAt time.Time
	LastLogin    time.Time
	Intro        string `gorm:"type:TINYTEXT DEFAULT NULL;"`
	Profile      string `gorm:"type:TEXT DEFAULT NULL;"`
}
