package util

import "simple_contact_book/model"

func AddContact(contact model.Contact, db []*model.Contact) []*model.Contact {
	db = append(db, &contact)
	return db
}
