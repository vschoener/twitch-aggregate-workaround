package migrate

import (
	"fmt"
	"log"

	"github.com/wonderstream/twitch/storage"
	"github.com/wonderstream/twitch/storage/model"
)

// Migrate manager
type Migrate struct {
	DB      *storage.Database
	Options Options
}

// Options for migration
type Options struct {
	DropIfInstall bool
}

// TableInfo contains table information for the install / migration
type TableInfo struct {
	Name    string
	Model   interface{}
	Options string
}

func (m Migrate) installTables(tables []TableInfo) {
	m.DB.Gorm.LogMode(true)
	for _, table := range tables {
		log.Println(fmt.Sprintf("Working on %s", table.Name))

		if m.Options.DropIfInstall && m.DB.Gorm.HasTable(table.Name) {
			log.Println("Drop table...")
			err := m.DB.Gorm.DropTable(table.Name).Error
			if err != nil {
				log.Println(err)
			}
		}

		if !m.DB.Gorm.HasTable(table.Name) {
			log.Println("Table doesn't exist, installing...")
			m.DB.Gorm.
				Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci").
				Table(table.Name).
				CreateTable(table.Model)
			errors := m.DB.Gorm.GetErrors()
			if len(errors) > 0 {
				log.Println(errors)
			} else {
				log.Println(fmt.Sprintf("Done"))
			}
		} else {
			log.Println("Table already exists", table.Name)
		}
	}

}

// Install will start the install database process
func (m Migrate) Install() {

	if nil == m.DB {
		log.Fatal("Did you forget to set a storage.Database instance ?")
	}

	mainOption := "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci"

	tables := []TableInfo{
		{
			Name:    model.Channel{}.TableName(),
			Model:   &model.Channel{},
			Options: mainOption,
		},
		{
			Name:    model.Credential{}.TableName(),
			Model:   &model.Credential{},
			Options: mainOption,
		},
		{
			Name:    model.User{}.TableName(),
			Model:   &model.User{},
			Options: mainOption,
		},
		{
			Name:    model.Video{}.TableName(),
			Model:   &model.Video{},
			Options: mainOption,
		},
		{
			Name:    model.PrecomputedChannel{}.TableName(),
			Model:   &model.PrecomputedChannel{},
			Options: mainOption,
		},
	}

	m.installTables(tables)
}
