package models

import "time"

type Group struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	Name       string    `json:"name" gorm:"not null"`
	Currency   string    `json:"currency" gorm:"default:TRY"`
	InviteCode string    `json:"invite_code" gorm:"uniqueIndex;not null"`
	CreatedBy  uint      `json:"created_by" gorm:"not null"`
	CreatedAt  time.Time `json:"created_at"`
	Members    []User    `json:"members" gorm:"many2many:group_members;"`
}

type GroupMember struct {
	GroupID  uint      `json:"group_id" gorm:"primaryKey"`
	UserID   uint      `json:"user_id" gorm:"primaryKey"`
	JoinedAt time.Time `json:"joined_at"`
}
