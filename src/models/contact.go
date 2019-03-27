package models

import "fmt"

// Contact structure
type Contact struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// ContactModelAll to get all contacts
func ContactModelAll() ([]*Contact, error) {
	rows, err := db.Query("SELECT * FROM contacts")

	if err != nil {
		return nil, err
	}

	// defer rows.Close()

	contacts := make([]*Contact, 0)
	for rows.Next() {
		contact := new(Contact)
		err := rows.Scan(&contact.ID, &contact.Name, &contact.Phone, &contact.CreatedAt, &contact.UpdatedAt)
		if err != nil {
			return nil, err
		}
		contacts = append(contacts, contact)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return contacts, nil
}

// ContactModelByID to get a contact by id
func ContactModelByID(id string) (*Contact, error) {
	rows, err := db.Query(fmt.Sprintf("SELECT * FROM contacts WHERE id=%s", id))

	if err != nil {
		return nil, err
	}

	contact := new(Contact)
	for rows.Next() {
		err := rows.Scan(&contact.ID, &contact.Name, &contact.Phone, &contact.CreatedAt, &contact.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}

	return contact, nil
}

// ContactModelCreate to create a contact
func ContactModelCreate(contact *Contact) (bool, error) {
	_, err := db.Query(fmt.Sprintf("INSERT INTO contacts(name, phone) VALUES ('%s', '%s')", contact.Name, contact.Phone))

	if err != nil {
		return false, err
	}

	return true, nil
}

// ContactModelDelete to delete a contact by ID
func ContactModelDelete(id string) (bool, error) {

	_, err := db.Query(fmt.Sprintf("DELETE FROM contacts WHERE id=%s", id))

	if err != nil {
		return false, nil
	}

	return true, nil
}
