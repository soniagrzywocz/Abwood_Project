package models

import (
	"go_server/db"
	"go_server/log"
)

type Contact struct {
	Name    string `db:"name" json:"name"`
	Email   string `db:"email" json:"email"`
	Message string `db:"message" json:"message"`
}

func (c Contact) SelectAllContacts() ([]*Contact, error) {
	query := "SELECT c.name, c.email, c.message FROM contact c"
	rows, err := db.Db().DB.Queryx(query)
	if err != nil {
		log.Errorf("Failed to Retrieve Contacts From Database: %v", err)
		return nil, err
	}
	defer rows.Close()
	var selectedContacts []*Contact

	for rows.Next() {
		contact := new(Contact)
		err := rows.StructScan(contact)
		if err != nil {
			log.Errorf("Failed to scan rows into contact struct: %v", err)
			return nil, err
		}
		selectedContacts = append(selectedContacts, contact)
	}

	return selectedContacts, nil
}
