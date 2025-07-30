package main

import (
	"fmt"
	"simple_contact_book/model"
	"simple_contact_book/util"
)

func main() {
	db := []*model.Contact{}
	user1 := model.Contact{
		Name:  "Ahmet",
		Phone: "05358771279",
		Email: "ahmetcndn",
	}

	db = util.AddContact(user1, db)
	fmt.Printf("value: %v",db)
}
