package publication_management_author

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

func (model *Author) BeforeCreate(scope *gorm.Scope) error {
	var uid uuid.UUID
	var err error
	if uid, err = uuid.NewV4(); err != nil {
		return err
	}
	return scope.SetColumn("Id", uid.String())
}
