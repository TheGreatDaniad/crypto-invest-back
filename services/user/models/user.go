package models

import (
	"errors"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	is "github.com/go-ozzo/ozzo-validation/v4/is"
	"gorm.io/gorm"
)

const (
	NotVerifiedUser = "not-verified"
	DisabledUser    = "disabled"
	StandardUser    = "standard"
	AdminUser       = "admin"
	RootUser        = "root"
)

type User struct {
	ID                  uint           `gorm:"primary_key" json:"id"`
	CreatedAt           time.Time      `json:"createdAt"`
	UpdatedAt           time.Time      `json:"-"`
	DeletedAt           gorm.DeletedAt `json:"-" sql:"index"`
	CompanyName         string         `json:"company_name"`
	FirstName           string         `json:"first_name"`
	LastName            string         `json:"last_name"`
	Email               string         `json:"email"`
	Password            string         `json:"-"`
	Country             string         `json:"country"`
	Phone               string         `json:"phone"`
	DateOfBirth         time.Time      `json:"date_of_birth"`
	Rank                string         `json:"rank"`
	Gender              string         `json:"gender"`
	Plan                string         `json:"plan"`
	ProfileImage        string         `json:"profile_image"`

}

type UserRegisterInfo struct {
	CompanyName string `json:"company_name"`
	Email       string `json:"email"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Phone       string `json:"phone"`
	Password    string `json:"password"`
}

type UserLoginInfo struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserQueryInfo struct {
	FirstName        string    `json:"first-name"`
	LastName         string    `json:"last-name"`
	CompanyName      string    `json:"company-name"`
	Email            string    `json:"email"`
	Country          string    `json:"country"`
	Gender           string    `json:"gender"`
	DateOfBirth      time.Time `json:"date-of-birth"`
	RegisteredBefore time.Time `json:"registered-before"`
	RegisteredAfter  time.Time `json:"registered-after"`
	Rank             string    `json:"rank"`
	Plan             string    `json:"plan"`
	Messages         bool      `json:"messages"`
	Files            bool      `json:"files"`
	Payments         bool      `json:"payments"`
	Limit            int       `json:"limit"`
}
type ResetPasswordRequestInfo struct {
	Email string `json:"email"`
}

type ResetPasswordInfo struct {
	Password string `json:"password"`
	Token    string `json:"token"`
}
type ChangePasswordInfo struct {
	Password string `json:"password"`
}

func (uli UserLoginInfo) Validate() error {

	return validation.ValidateStruct(&uli,
		validation.Field(&uli.Email, validation.Required, is.Email),

		//Password must contain minimum 8 characters, maximum 40 characters,
		//At least one uppercase letter, one lowercase letter and one number
		validation.Field(&uli.Password, validation.Required, validation.Length(8, 40)),
	)
}

func (uri UserRegisterInfo) Validate() error {
	return validation.ValidateStruct(&uri,
		validation.Field(&uri.Email, validation.Required, is.Email),
		//Password must contain minimum 8 characters, maximum 40 characters,
		//At least one uppercase letter, one lowercase letter and one number
		validation.Field(&uri.Password, validation.Required, validation.Length(8, 40)), //TODO check for other validation rules
		validation.Field(&uri.Phone, is.E164),                                          //check functionality

	)
}
func (u User) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Email, is.Email),
		//Password must contain minimum 8 characters, maximum 40 characters,
		//At least one uppercase letter, one lowercase letter and one number
		validation.Field(&u.Password, validation.Length(8, 40)), //TODO check for other validation rules
		validation.Field(&u.FirstName, validation.Length(2, 50), is.Alpha),
		validation.Field(&u.LastName, validation.Length(2, 50), is.Alpha),
		validation.Field(&u.Country, validation.Length(2, 50), is.Alpha),
		validation.Field(&u.Gender, validation.By(validateGender)),
		validation.Field(&u.Rank, validation.By(validateRank)),
		validation.Field(&u.Phone, is.E164), //check functionality
	)
}

func (uqi UserQueryInfo) Validate() error {
	return validation.ValidateStruct(&uqi,
		validation.Field(&uqi.Email, is.Email),
		validation.Field(&uqi.FirstName, validation.Length(2, 50), is.Alpha),
		validation.Field(&uqi.LastName, validation.Length(2, 50), is.Alpha),
		validation.Field(&uqi.Country, validation.Length(2, 50), is.Alpha),
		validation.Field(&uqi.Gender, validation.By(validateGender)),
	)
}

func (rpri ResetPasswordRequestInfo) Validate() error {
	return validation.ValidateStruct(&rpri,

		validation.Field(&rpri.Email, validation.Required, is.Email),
	)
}
func (rpi ResetPasswordInfo) Validate() error {
	return validation.ValidateStruct(&rpi,
		validation.Field(&rpi.Password, validation.Length(8, 40)), //TODO check for other validation rules
	)
}
func (cpi ChangePasswordInfo) Validate() error {
	return validation.ValidateStruct(&cpi,

		validation.Field(&cpi.Password, validation.Required, validation.Length(8, 40)), //TODO check for other validation rules
	)
}
func (u *User) HidePassword() {
	u.Password = "****"
}

func (uri UserRegisterInfo) GenerateUserModel() (User, error) {

	var birthDay time.Time

	return User{
		Email:       uri.Email,
		CompanyName: uri.CompanyName,
		Password:    uri.Password,
		FirstName:   uri.FirstName,
		LastName:    uri.LastName,
		DateOfBirth: birthDay,
		Rank:        NotVerifiedUser,
	}, nil
}
func validateGender(value interface{}) error {
	switch value {
	case "male":
		return nil
	case "female":
		return nil
	case "non-binary":
		return nil
	case "":
		return nil
	}
	return errors.New("gender must be one of: male, female, non-binary or prefer not to say")
}
func validateRank(value interface{}) error {
	switch value {
	case NotVerifiedUser:
		return nil
	case DisabledUser:
		return nil
	case StandardUser:
		return nil
	case AdminUser:
		return nil
	case RootUser:
		return nil
	case "":
		return nil
	}
	return errors.New("rank is not valid")
}

func RankNumber(rank string) int {
	rankPriority := 0
	switch rank {
	case NotVerifiedUser:
		rankPriority = 0

	case DisabledUser:
		rankPriority = 1

	case StandardUser:
		rankPriority = 2

	case AdminUser:
		rankPriority = 3

	case RootUser:
		rankPriority = 4

	}
	return rankPriority
}

func IsHigherRank(rank1 string, rank2 string) bool {
	return RankNumber(rank1) > RankNumber((rank2))
}
func IsHigherOrEqualRank(rank1 string, rank2 string) bool {
	return RankNumber(rank1) >= RankNumber((rank2))
}

func (u User) CanBeModifiedBy(user User) bool {
	if user.ID == u.ID {
		return true
	}
	if user.Rank == StandardUser {
		return false
	}
	return IsHigherOrEqualRank(user.Rank, u.Rank)
}
