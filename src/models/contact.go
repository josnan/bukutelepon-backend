package models

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
