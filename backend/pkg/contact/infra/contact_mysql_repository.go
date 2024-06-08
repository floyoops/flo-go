package infra

import (
	"fmt"
	"github.com/floyoops/flo-go/backend/pkg/contact/domain/model"
	"github.com/floyoops/flo-go/backend/pkg/contact/repository"
	"github.com/floyoops/flo-go/backend/pkg/database"
)

type ContactMysqlRepository struct {
	db *database.Database
}

func NewContactMysqlRepository(database *database.Database) *ContactMysqlRepository {
	r := &ContactMysqlRepository{db: database}
	r.InitSchema()
	return r
}

func (r *ContactMysqlRepository) InitSchema() {
	schema := `
		CREATE TABLE IF NOT EXISTS contact (
		    uuid VARCHAR(36) PRIMARY KEY,
		    name VARCHAR(255) NOT NULL,
		    email VARCHAR(255) NOT NULL,
		    message TEXT NULL,
		    created_at DATETIME NOT NULL,
		    updated_at DATETIME NOT NULL
		)
	`
	r.db.Connection.MustExec(schema)
}

func (c *ContactMysqlRepository) Create(contact *model.Contact) error {
	query := "INSERT INTO contact (uuid, name, email, message, created_at, updated_at) VALUES (:uuid, :name, :email, :message, :created_at, :updated_at)"
	tx := c.db.Connection.MustBegin()
	_, err := tx.NamedExec(
		query,
		map[string]interface{}{
			"uuid":       contact.Uuid.String(),
			"name":       contact.Name,
			"email":      contact.Email.String(),
			"message":    contact.Message,
			"created_at": contact.CreatedAt.ToMysqlDateTime(),
			"updated_at": contact.UpdatedAt.ToMysqlDateTime(),
		},
	)

	if err != nil {
		errRollback := tx.Rollback()
		isRollback := errRollback == nil

		return fmt.Errorf("query: %s, isRollback: %t, %w", query, isRollback, repository.ErrOnSaveContact)
	}
	return tx.Commit()
}
