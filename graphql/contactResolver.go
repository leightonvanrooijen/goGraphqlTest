package graph

import (
	"fmt"

	"github.com/graph-gophers/graphql-go"
	"github.com/graphql-basics/gqldb"
)

type ContactResolver struct {
	db      gqldb.BasicDb
	contact *gqldb.Contact
}

// Gets the contact object for ContactResolvers
func (r RootResolver) Contact(args struct{ ID int32 }) (*ContactResolver, error) {
	contact, err := r.db.GetContactByID(uint(args.ID))
	if err != nil {
		return nil, fmt.Errorf("error [%s]: Could not find contact with id: %d", err, args.ID)
	}

	return &ContactResolver{db: r.db, contact: contact}, nil
}

// Gets the ID for contact
func (c *ContactResolver) ID() *graphql.ID {
	contactID := graphql.ID(fmt.Sprint(c.contact.ID))

	return &contactID
}

// Gets email address for contact
func (c *ContactResolver) Email() *string {
	return &c.contact.Email
}

// Gets Phone number for contact
func (c *ContactResolver) Phone() *string {
	return &c.contact.Phone
}
