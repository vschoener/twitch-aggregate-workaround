package repository

import (
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
	query := storage.Query{
		Query: `
			SELECT
				user_id,
				bio,
				display_name,
				logo,
				name,
				type,
				created_at,
				updated_at,
				date_add
			FROM ` + model.UserTable + `
			ORDER BY id DESC
		`,
	}

	rows := r.Database.Query(query)
	if rows == nil {
		return nil
	}

	defer rows.Close()

	users := []model.User{}

	for rows.Next() {
		user := model.User{}
		state := r.Database.ScanRows(rows,
			&user.UserID,
			&user.Bio,
			&user.DisplayName,
			&user.Logo,
			&user.Name,
			&user.Type,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.DateAdd,
		)
		if state {
			users = append(users, user)
		}
	}
	if err := rows.Err(); err != nil {
		r.Logger.LogInterface(err)
		return nil
	}

	return users
}

// StoreUser inserts User info in the storage and keep history
func (r UserRepository) StoreUser(user model.User) bool {
	query := storage.Query{
		Query: `
            INSERT INTO ` + model.UserTable + `
            (
                user_id,
                bio,
                display_name,
                logo,
                name,
                type,
                created_at,
                updated_at
            )
            VALUES(?, ?, ?, ?, ?, ?, ?, ?)
            ON DUPLICATE KEY UPDATE
                bio = ?,
                display_name = ?,
                logo = ?,
                name = ?,
                type = ?
        `,
		Parameters: map[string]interface{}{
			"StoreUsers": user,
		},
	}

	state := r.Database.Run(query,
		user.UserID,
		user.Bio,
		user.DisplayName,
		user.Logo,
		user.Name,
		user.Type,
		user.CreatedAt,
		user.UpdatedAt,
		user.Bio,
		user.DisplayName,
		user.Logo,
		user.Name,
		user.Type,
	)

	return state
}
