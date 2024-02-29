package repository

import (
	"architecture_go/services/contact/internal/domain"
	"database/sql"

	_ "github.com/lib/pq"
)

type ContactRepository interface {
	CreateContact(contact domain.Contact) error
	ReadContact(contactID int) (domain.Contact, error)
	UpdateContact(contactID int, newContact domain.Contact) error
	DeleteContact(contactID int) error

	CreateGroup(group domain.Group) error
	ReadGroup(groupID int) (domain.Group, error)

	AddContactToGroup(contactID, groupID int) error
}

type PostgresContactRepository struct {
	DB *sql.DB
}

func NewPostgresContactRepository(db *sql.DB) *PostgresContactRepository {
	return &PostgresContactRepository{DB: db}
}

func (r *PostgresContactRepository) CreateContact(contact *domain.Contact) error {
	query := `INSERT INTO contacts (first_name, last_name, middle_name, phone_number) VALUES ($1, $2, $3, $4)`
	_, err := r.DB.Exec(query, contact.FullName.FirstName, contact.FullName.LastName, contact.FullName.MiddleName, contact.PhoneNumber())
	return err
}

func (r *PostgresContactRepository) ReadContact(contactID int) (*domain.Contact, error) {
	query := `SELECT id, first_name, last_name, middle_name, phone_number FROM contacts WHERE id = $1`
	var firstName, lastName, middleName, phoneNumber string
	err := r.DB.QueryRow(query, contactID).Scan(&contactID, &firstName, &lastName, &middleName, &phoneNumber)
	if err != nil {
		return nil, err
	}
	return domain.NewContact(contactID, lastName, firstName, middleName, phoneNumber), nil
}

func (r *PostgresContactRepository) UpdateContact(contactID int, newContact *domain.Contact) error {
	query := `UPDATE contacts SET first_name = $2, last_name = $3, middle_name = $4, phone_number = $5 WHERE id = $1`
	_, err := r.DB.Exec(query, contactID, newContact.FullName.FirstName, newContact.FullName.LastName, newContact.FullName.MiddleName, newContact.PhoneNumber())
	return err
}

func (r *PostgresContactRepository) DeleteContact(contactID int) error {
	query := `DELETE FROM contacts WHERE id = $1`
	_, err := r.DB.Exec(query, contactID)
	return err
}
