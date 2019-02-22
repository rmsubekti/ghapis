package models

import "github.com/jinzhu/gorm"

// RoleType custom type
type RoleType string

// RoleUser enum
const (
	RoleUser  RoleType = "user"
	RoleAdmin RoleType = "admin"
	RolePM    RoleType = "pm"
)

// Role Account
type Role struct {
	gorm.Model
	RoleName RoleType
}
