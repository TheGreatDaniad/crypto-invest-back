package impl

import (
	"errors"
	"time"

	"github.com/thegreatdaniad/crypto-invest/constants/customErrors"
	gatewayModels "github.com/thegreatdaniad/crypto-invest/gateway/models"
	"github.com/thegreatdaniad/crypto-invest/services/user/models"
	"gorm.io/gorm"
)

type Database struct {
	Postgres *gorm.DB
}

func (db Database) CreateUser(c *gatewayModels.Carrier, user models.UserRegisterInfo) (models.User, error) {

	select {
	case <-c.Context.Done():
		err := errors.New("inserting user to the database has been canceled from a parent process")
		return models.User{}, err
	default:
		// creates user in the database and return the ID of the user
		newUser, err := user.GenerateUserModel()
		if err != nil {
			return models.User{}, err
		}
		res := db.Postgres.Create(&newUser)

		return newUser, res.Error
	}

}

func (db Database) UserDoesNotExist(c *gatewayModels.Carrier, user models.UserRegisterInfo) (bool, error) {

	select {
	case <-c.Context.Done():
		err := errors.New(customErrors.ProccessCanceledByParent)
		return false, err
	default:
		userQuery := models.User{}
		res := db.Postgres.First(&userQuery, "email = ?", user.Email)

		if res.Error != nil {
			return !errors.Is(res.Error, gorm.ErrRecordNotFound), nil
		}
		return true, nil

	}
}

func (db Database) GetUserByEmail(c *gatewayModels.Carrier, email string) (models.User, error) {
	select {
	case <-c.Context.Done():
		err := errors.New(customErrors.ProccessCanceledByParent)
		return models.User{}, err
	default:
		userQuery := models.User{}
		res := db.Postgres.First(&userQuery, "email = ?", email)

		if res.Error != nil {

			return models.User{}, res.Error
		}
		return userQuery, nil

	}
}
func (db Database) GetUserById(c *gatewayModels.Carrier, id uint) (models.User, error) {
	select {
	case <-c.Context.Done():
		err := errors.New(customErrors.ProccessCanceledByParent)
		return models.User{}, err
	default:
		userQuery := models.User{}
		res := db.Postgres.First(&userQuery, "id = ?", id)

		if res.Error != nil {
			return models.User{}, res.Error
		}
		return userQuery, nil

	}
}

func (db Database) GetUsers(c *gatewayModels.Carrier, queryParams models.UserQueryInfo) ([]models.User, error) {
	select {
	case <-c.Context.Done():
		err := errors.New(customErrors.ProccessCanceledByParent)
		return []models.User{}, err
	default:
		return queryUsers(db, queryParams)
	}
}

func queryUsers(db Database, queryParams models.UserQueryInfo) ([]models.User, error) {
	var users []models.User
	queryMap := generateQueryMap(queryParams)
	if !queryParams.RegisteredBefore.IsZero() || !queryParams.RegisteredAfter.IsZero() {
		if queryParams.RegisteredBefore.IsZero() {
			queryParams.RegisteredBefore = time.Now()
		}

		res := db.Postgres.Where(queryMap).Where("created_at BETWEEN ? AND ?", queryParams.RegisteredAfter, queryParams.RegisteredBefore).Find(&users)
		if res.Error != nil {
			return []models.User{}, res.Error
		}
		return users, nil
	}
	res := db.Postgres.Where(queryMap).Find(&users)
	if res.Error != nil {
		return []models.User{}, res.Error
	}
	return users, nil

}
func (db Database) UpdateUser(c *gatewayModels.Carrier, id uint, u models.User) error {
	user := models.User{}
	updateMap := generateUpdateMap(u)
	res := db.Postgres.Model(&user).Where("id =?", id).Updates(updateMap)

	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (db Database) UpdateProfileImage(c *gatewayModels.Carrier, id uint, imagePath string) error {
	user := models.User{}
	res := db.Postgres.Model(&user).Where("id =?", id).Update("profile_image", imagePath)

	if res.Error != nil {
		return res.Error
	}
	return nil
}
func (db Database) UpdateUserPassword(c *gatewayModels.Carrier, id uint, u models.User) error {

	res := db.Postgres.Model(&u).Where("id =?", id).Update("password", u.Password)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
func (db Database) DeleteUser(c *gatewayModels.Carrier, ID uint) error {
	select {
	case <-c.Context.Done():
		err := errors.New(customErrors.ProccessCanceledByParent)
		return err
	default:
		res := db.Postgres.Delete(&models.User{}, ID)

		if res.Error != nil {
			return res.Error
		}
		return nil

	}
}

func generateQueryMap(queryParams models.UserQueryInfo) map[string]interface{} {
	queryMap := make(map[string]interface{})
	if !(queryParams.Email == "") {
		queryMap["email"] = queryParams.Email
	}
	if !(queryParams.FirstName == "") {
		queryMap["first_name"] = queryParams.FirstName
	}
	if !(queryParams.LastName == "") {
		queryMap["last_name"] = queryParams.LastName
	}
	if !(queryParams.Country == "") {
		queryMap["country"] = queryParams.Country
	}
	if !(queryParams.Gender == "") {
		queryMap["gender"] = queryParams.Gender
	}
	if !(queryParams.Rank == "") {
		queryMap["rank"] = queryParams.Rank
	}

	if !(queryParams.DateOfBirth == time.Time{}) {
		queryMap["date_of_birth"] = queryParams.DateOfBirth
	}

	return queryMap
}

func generateUpdateMap(user models.User) map[string]interface{} {
	updateMap := make(map[string]interface{})
	if !(user.Email == "") {
		updateMap["email"] = user.Email
	}
	if !(user.FirstName == "") {
		updateMap["first_name"] = user.FirstName
	}
	if !(user.LastName == "") {
		updateMap["last_name"] = user.LastName
	}
	if !(user.Country == "") {
		updateMap["country"] = user.Country
	}
	if !(user.Gender == "") {
		updateMap["gender"] = user.Gender
	}
	if !(user.Rank == "") {
		updateMap["rank"] = user.Rank
	}
	if !(user.DateOfBirth == time.Time{}) {
		updateMap["date_of_birth"] = user.DateOfBirth
	}
	if !(user.Plan == "") {
		updateMap["plan"] = user.Plan
	}

	return updateMap
}
