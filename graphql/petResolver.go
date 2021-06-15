package graph

import (
	"fmt"

	"github.com/graph-gophers/graphql-go"
	"github.com/graphql-basics/gqldb"
)

type PetResolver struct {
	db  gqldb.BasicDb
	pet *gqldb.Pet
}

// Gets the user object for UserResolvers
func (r RootResolver) Pet(args struct{ ID int32 }) (*PetResolver, error) {
	pet, err := r.db.GetPetByID(uint(args.ID))
	if err != nil {
		
		return nil, fmt.Errorf("error [%s]: Could not find pet with id: %d", err, args.ID)
	}

	return &PetResolver{db: r.db, pet: pet}, nil
}

// Gets the ID for users
func (p *PetResolver) ID() *graphql.ID {
	petID := graphql.ID(fmt.Sprint(p.pet.ID))

	return &petID
}

// Gets name for user
func (p *PetResolver) Name() *string {
	return &p.pet.Name
}

// Gets name from user
func (p *PetResolver) Age() *int32 {
	return &p.pet.Age
}

// Gets name from user
func (p *PetResolver) Species() *string {
	return &p.pet.Species
}
  