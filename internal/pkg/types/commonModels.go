package types

import (
	"github.com/satori/go.uuid"
)

type DBData struct {
	DriverName       string `json: "DriverName"`
	DataBaseUser     string `json: "DataBaseUser"`
	DataBaseName     string `json: "DataBaseName"`
	DataBasePassword string `json: "DataBasePassword"`
	SSLMode          string `json: "SSLMode"`
}

type InfoForReg struct {
	Email           string       `json:"email"`
	Password        string       `json:"password"`
	ConfirmPassword string       `json:"confirmPassword"`
	Claims          []UserClaims `json:"claims"`
}

type UserClaims struct {
	Type  string `json:"type"`
	Value string `json:"Value"`
}
type User struct {
	FirstName       string `json: "firstname"`
	LastName        string `json: "lastname"`
	Email           string `json: "email"`
	Password        string `json: "password"`
	ConfirmPassword string `json: "confirmPassword"`
}

type UserInf struct {
	Id                 uuid.UUID `json: "id"`
	UserFirstname      string    `json: "firstname"`
	UserLastname       string    `json:"lastname"`
	NormalizedUserName string    `json: "normalizedUserNmae"`
	Email              string    `json: "email"`
	EmailConfirmed     bool      `json: "emailConfirmed"`
	NormalizedEmail    string    `json: "normalizedEmail"`
	PasswordHash       string    `json: "passwordHash"`
	SecurityStamp      string    `json: "securetyStamp"`
	ConcurrencyStamp   string    `json: "concurrencyStamp"`
	LockoutEnd         string    `json: "lockoutEnd"` // time.Time
	LockoutEnabled     bool      `json: "lockoutEnabled"`
	AccessFailedCount  int       `json: "accessFailedCount"`
	Birthdate          string    `json:"birthdate"`
}

type UserStruct struct {
	Email           string       `json:"Email" db:"email"`
	Password        string       `json:"Password" db:"hash_password"`
	ConfirmPassword string       `json:"ConfirmPassword" db:"confirmpassword"`
	Claims          []ClaimsType `json:"Claims" db:"claims"`
}

type ClaimsType struct {
	Type  string `json:"Type" db:"claimtype"`
	Value string `json:"Value" db "claimvalue"`
}
type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Answer struct {
	Claims map[string]string `json:"claims"`
}
