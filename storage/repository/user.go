package repository

import (
	"time"

	"github.com/wonderstream/twitch/logger"
	"github.com/wonderstream/twitch/storage"
	"github.com/wonderstream/twitch/storage/model"
)

// UserRepository handles user database query
type UserRepository struct {
	*Repository
}

// NewUserRepository returns user repository
func NewUserRepository(db *storage.Database, l logger.Logger) UserRepository {
	commonRepository := NewRepository(db, l)
	r := UserRepository{
		Repository: commonRepository,
	}

	return r
}

// GetUsers returns a User list
func (r UserRepository) GetUsers() []model.User {
	users := []model.User{}
	r.Database.Gorm.Find(&users)

	return users
}

// StoreUser inserts User info in the storage and keep history
func (r UserRepository) StoreUser(user model.User) bool {
	newUser := model.User{}
	newUser.MetaDateAdd = time.Now()
	user.MetaDateUpdate = time.Now()
	err := r.Database.Gorm.
		Where(model.User{UserID: user.UserID}).
		Assign(user).
		FirstOrCreate(&newUser).
		Error

	return r.CheckErr(err)
}
