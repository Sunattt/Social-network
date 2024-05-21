package models

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Server struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type Db struct {
	DbName   string `json:"db_name"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
}

type ConfigsModels struct {
	Database  Db     `json:"database"`
	Server    Server `json:"server"`
	JWTSecret string `json:"secret_key"`
}

type Users struct {
	NikName      string    `gorm:"type:text;not null" json:"nik_name"`
	Name         string    `gorm:"type:text;not null"`
	Birthday     string    `gorm:"type:timestamp;not null"`
	Phone        string    `gorm:"type:text;not null"`
	Country      string    `gorm:"type:text;not null"`
	Email        string    `gorm:"type:text;not null"`
	Password     string    `gorm:"type:text;not null"`
	PhotoProfile string    `gorm:"type:text;not null;default:'no data'"`
	Bio          string    `gorm:"type:text;not null;default:'no data'"`
	Id           uint      `gorm:"primary_key;auto_increment;not null"`
	Active       bool      `gorm:"not null;default:true"`
	CreatedAt    time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt    time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
}

type Posts struct {
	Id              uint      `gorm:"primary_key;auto_increment;not null"`
	Active          bool      `gorm:"not null;default:true"`
	CreatedAt       time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt       time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt       time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	ContentNameFile string    `gorm:"type:text;not null"`
	ViewCount       int       `gorm:"not null"`
	DataPublished   time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UserID          int       `gorm:"not null"`
	User            Users     `gorm:"not null;foreignkey:UserID"`
}

type Tokens struct {
	Id        uint      `gorm:"primary_key;auto_increment;not null"`
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	Token     string    `gorm:"type:text;not null"`
	UserID    uint      `gorm:"not null"`
	User      Users     `gorm:"not null;foreignkey:UserID"`
	Expire    time.Time
}

type Following struct {
	Id          uint      `gorm:"primary_key;auto_increment;not null"`
	Active      bool      `gorm:"not null;default:true"`
	CreatedAt   time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt   time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UserID      int       `gorm:"not null"`
	User        Users     `gorm:"not null;foreignkey:UserID"`
	FollowingID int       `gorm:"not null"`
	Following   Users     `gorm:"not null;foreignkey:FollowingID"`
}

type Followers struct {
	Id         uint      `gorm:"primary_key;auto_increment;not null"`
	Active     bool      `gorm:"not null;default:true"`
	CreatedAt  time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt  time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UserID     int       `gorm:"not null"`
	User       Users     `gorm:"not null;foreignkey:UserID"`
	FollowerID int       `gorm:"not null"`
	Follower   Users     `gorm:"not null;foreignkey:FollowerID"`
}

type Likes struct {
	Id        uint      `gorm:"primary_key;auto_increment;not null"`
	Active    bool      `gorm:"not null;default:true"`
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UserID    int       `gorm:"not null"`
	User      Users     `gorm:"not null;foreignkey:UserID"`
	PostID    int       `gorm:"not null"`
	Post      Posts     `gorm:"not null;foreignkey:PostID"`
}

type Comments struct {
	Id        uint      `gorm:"primary_key;auto_increment;not null"`
	Active    bool      `gorm:"not null;default:true"`
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UserID    int       `gorm:"not null"`
	User      Users     `gorm:"not null;foreignkey:UserID"`
	PostID    int       `gorm:"not null"`
	Post      Posts     `gorm:"not null;foreignkey:PostID"`
	Comment   string    `gorm:"type:text;not null"`
	Date      time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
}

type Answer struct {
	Date   time.Time
	Answer string
}

type SendToken struct {
	Date   time.Time
	Answer string
	Token  string
}
type TokenClaims struct {
	NikName string    `json:"nik_name"`
	Expire  time.Time `json:"expire"`
	jwt.StandardClaims
}
