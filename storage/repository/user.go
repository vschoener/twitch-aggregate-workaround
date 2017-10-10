package repository

import (
	"github.com/wonderstream/twitch/storage"
	"github.com/wonderstream/twitch/storage/model"
)

// UserRepository handles user database query
type UserRepository struct {
	*Repository
}

// StoreUsers inserts User info in the storage and keep history
func (r UserRepository) StoreUsers(user model.User) bool {
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
