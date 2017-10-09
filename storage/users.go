package storage

import (
	"fmt"
	"time"

	"github.com/wonderstream/twitch/core/api/users"
)

// Users is the storage manager
type Users struct {
	ID      int64
	DateAdd time.Time
	users.User
}

const (
	usersTable = "users"
)

// StoreUsers inserts User info in the storage and keep history
func (d *Database) StoreUsers(user users.User) bool {
	queryLogger := QueryLogger{
		Query: `
            INSERT INTO ` + usersTable + `
            (
                users_id,
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

	d.Logger.Log(fmt.Sprintf("StoreUsers on %#v", queryLogger))
	stmt, err := d.DB.Prepare(queryLogger.Query)

	if err != nil {
		d.Logger.LogInterface(err)
		return false
	}

	defer stmt.Close()
	_, err = stmt.Exec(
		user.ID,
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

	if err != nil {
		d.Logger.LogInterface(err)
		return false
	}

	return true
}
