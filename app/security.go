package app

import (
	"net/http"

	"github.com/rmsubekti/ghapis/app/models"

	"github.com/gorilla/context"
)

// RoleAccess verification access
func RoleAccess(r *http.Request, expectedRoles []models.RoleType) bool {
	claim := context.Get(r, Xclaim).(*models.Token)
	roles := claim.Roles
	if len(roles) == 0 {
		return false
	}

	p := false
	for _, v := range expectedRoles {
		p = isAllowed(roles, v)
		if p {
			break
		}
	}
	return p
}

// GetUserID from Xclaim context
func GetUserID(r *http.Request) uint {
	claim := context.Get(r, Xclaim).(*models.Token)
	return claim.UserID
}

// isAllowed logged user roles
func isAllowed(roles []models.Role, expected models.RoleType) bool {
	for _, v := range roles {
		if v.RoleName == expected {
			return true
		}
	}
	return false
}
