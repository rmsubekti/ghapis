package models

import (
	"os"
	"strings"

	"github.com/rmsubekti/ghapis/app/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// Token JWT claim struct
type Token struct {
	UserID   uint
	UserName string
	Roles    []Role
	jwt.StandardClaims
}

// Account user struct
type Account struct {
	gorm.Model
	UserName string  `json:"username"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Roles    []Role  `gorm:"many2many:user_roles;"`
	Orders   []Order `gorm:"foreignkey:AccountID;association_foreignkey:Refer"`
}

// Validate on account creation
func (account *Account) Validate() (map[string]interface{}, bool) {

	if account.UserName == "" {
		return utils.Message(false, "Username is required"), false
	}
	// validate email
	if !strings.Contains(account.Email, "@") {
		return utils.Message(false, "Email address is required"), false
	}

	// validate password
	if len(account.Password) < 6 {
		return utils.Message(false, "Password should be 6 or more character"), false
	}

	// check email if it is used
	temp := &Account{}
	err := GetDB().Table("accounts").Where("email=?", account.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return utils.Message(false, "Connection error"), false
	}
	// email is in use
	if temp.Email != "" {
		return utils.Message(false, "Email already in use by another user"), false
	}

	// check username if it is used
	e := GetDB().Table("accounts").Where("user_name=?", account.UserName).First(temp).Error
	if e != nil && e != gorm.ErrRecordNotFound {
		return utils.Message(false, "Connection error"), false
	}
	// email is in use
	if temp.UserName != "" {
		return utils.Message(false, "Username already in use by another user"), false
	}

	// email can be use to register
	return utils.Message(false, "Requirement passed"), true

}

// Create new Account
func (account *Account) Create() map[string]interface{} {

	if resp, ok := account.Validate(); !ok {
		return resp
	}

	// set default user role
	role := &Role{}
	err := GetDB().Table("roles").Where("role_name=?", string(RoleUser)).First(role).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return utils.Message(false, "Connection error")
	}
	account.Roles = []Role{*role}
	// encrypt password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hashedPassword)

	//write account in database
	GetDB().Create(account)
	if account.ID <= 0 {
		return utils.Message(false, "Failed t create account, connection failed")
	}

	//generate token for newly registered account
	tk := &Token{
		UserID:   account.ID,
		UserName: account.UserName,
		Roles:    account.Roles,
	}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("TOKEN_PASWORD")))

	// send response with token
	response := utils.Message(true, "Account has been created")
	response["token"] = tokenString
	return response

}

// Login account
func Login(userNameOrEmail, pass string) map[string]interface{} {
	account := &Account{}
	err := GetDB().Where("email = ?", userNameOrEmail).Or("user_name = ?", userNameOrEmail).Preload("Roles").Find(account).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.Message(false, "invalid login credential, Email not registered")
		}
		return utils.Message(false, "Connection error")
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(pass))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return utils.Message(false, "invalid login credential, password not match")
	}

	tk := &Token{
		UserID:   account.ID,
		UserName: account.UserName,
		Roles:    account.Roles,
	}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("TOKEN_PASWORD")))

	response := utils.Message(true, "logged in")
	response["token"] = tokenString

	return response
}

// GetUser by id
func GetUser(u uint) map[string]interface{} {
	account := &Account{}
	GetDB().Table("accounts").Where("id = ?", u).First(account)
	if account.Email == "" {
		return nil
	}
	account.Password = "*o*"
	response := utils.Message(true, "You")
	response["user"] = account
	return response
}
