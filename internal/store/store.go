package store

import (
	"container/list"
)

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

type Session struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	SessionID string `json:"session_id"`
	UserID    uint   `json:"user_id"`
	User      User   `gorm:"foreignKey:UserID" json:"user"`
}

type Cart struct {
	ID     uint    `gorm: "primaryKey" json:"id"`
	Amount float32 `json:"amount"`
	User   User    `gorm:"foreignKey:UserID" json:"user"`
	Items  list.List
}

type Item struct {
	ID    uint    `gorm: "primaryKey" json:"id"`
	price float32 `json:"amount"`
	name  string  `json:name"`
}

type CartStore interface {
	CreateCart(amount float32, user User) error
	GetCart(user User) (*Cart, error)
	AddItemToCart(user User, item Item) error
}

type UserStore interface {
	CreateUser(email string, password string) error
	GetUser(email string) (*User, error)
}

type SessionStore interface {
	CreateSession(session *Session) (*Session, error)
	GetUserFromSession(sessionID string, userID string) (*User, error)
}
