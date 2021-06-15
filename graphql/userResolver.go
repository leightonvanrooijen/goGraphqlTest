package graph

import (
	"context"
	"fmt"
	"strconv"

	"github.com/graph-gophers/graphql-go"
	"github.com/graphql-basics/gqldb"
	"github.com/graphql-basics/logger"
)

type RootResolver struct {
	db gqldb.BasicDb
}

type UserResolver struct {
	db   gqldb.BasicDb
	user *gqldb.User
}

// Gets the user object for UserResolvers
func (r RootResolver) User(args struct{ ID int32 }) (*UserResolver, error) {
	user, err := r.db.GetUserByID(uint(args.ID))
	if err != nil {
		return nil, fmt.Errorf("error [%s]: Could not find user with Id: %d", err, args.ID)
	}

	return &UserResolver{db: r.db, user: user}, nil
}

// Gers the ID for users
func (u *UserResolver) ID() *graphql.ID {
	userID := graphql.ID(fmt.Sprint(u.user.ID))

	return &userID
}

// Gets name for user
func (u *UserResolver) Name() *string {
	return &u.user.Name
}

// Gets name from user
func (u *UserResolver) Age() *int32 {
	return &u.user.Age
}

// Creates a new user
func (r *RootResolver) CreateUser(ctx context.Context, args *struct {
	Name string
	Age  int32
}) *UserResolver {
	user := gqldb.User{
		Name: args.Name,
		Age:  args.Age,
	}

	dbUser := r.db.CreateUser(&user)

	return &UserResolver{db: r.db, user: dbUser}
}

// Creates a new user
func (u *RootResolver) DeleteUser(ctx context.Context, args *struct{ ID graphql.ID }) (*string, error) {
	id, _ := strconv.ParseInt(string(args.ID), 10, 32)

	if _, err := u.db.DeleteUser(uint(id)); err != nil {
		logger.Warn.Printf("user with id %d could not be deleted", id)
		return nil, fmt.Errorf("user with id %d could not be deleted", id)

		
	}

	message := fmt.Sprintf("User wit id %d has been deleted", id)

	return &message, nil
}

// Gets contact for a user
func (u *UserResolver) Contact() *ContactResolver {
	return &ContactResolver{db: u.db, contact: &u.user.Contact}
}

// Gets all pets for a user
func (u *UserResolver) Pets() []*PetResolver {
	pets := u.db.GetPetsByUserID(uint(u.user.ID))

	petRxs := make([]*PetResolver, len(pets))
	for i := range pets {
		petRxs[i] = &PetResolver{
			db:  u.db,
			pet: &pets[i],
		}
	}
	return petRxs
}
