package gqldb

import (
	"fmt"

	"github.com/graphql-basics/logger"
)

// find User by ID
func (ctx BasicDb) GetUserByID(id uint) (*User, error) {
	var user User
	if err := ctx.db.Joins("Contact").First(&user, id); err.Error != nil {
		return nil, err.Error
	}
	return &user, nil
}

// find Contact by ID
func (ctx BasicDb) GetContactByID(id uint) (*Contact, error) {
	var contact Contact
	if err := ctx.db.First(&contact, "ID = ?", id); err.Error != nil {
		return nil, err.Error
	}
	return &contact, nil
}

// find Pet by ID
func (ctx BasicDb) GetPetByID(id uint) (*Pet, error) {
	var pet Pet
	if err := ctx.db.First(&pet, "ID = ?", id); err.Error != nil {
		return nil, fmt.Errorf("error [%v]: Could not find pet with the ID of %d", err, id)
	}
	return &pet, nil
}

// find contacts belonging to a user
func (ctx BasicDb) GetContactByUserID(id uint) (*Contact, error) {
	var contact Contact
	if err := ctx.db.Where("user_id = ?", id).Find(&contact); err.Error != nil {
		return nil, fmt.Errorf("error [%v]: Could not find conatcts with the user ID of %d", err, id)
	}
	return &contact, nil
}

// find pets belonging to a user
func (ctx BasicDb) GetPetsByUserID(id uint) []Pet {
	var pets []Pet
	if err := ctx.db.Where("user_id = ?", id).Find(&pets); err.Error != nil {
		logger.Warn.Printf("Could not find any pets with the User ID of %d. Error: %v", id, err)
	}
	return pets
}

// Creates a user based on input
func (ctx BasicDb) CreateUser(User *User) *User {
	if err := ctx.db.Create(&User); err.Error != nil {
		logger.Warn.Printf("Could not create user: %s.", User.Name)
	}
	return User
}

// delete pets belonging to a user
func (ctx BasicDb) DeletePetsByUserID(id uint) (bool, error) {
	var pet Pet
	if err := ctx.db.Where("user_id = ?", id).Delete(&pet); err.Error != nil {
		logger.Warn.Printf("Could not delete pets belonging to user id: %d", id)
		return false, err.Error
	}
	return true, nil
}

// delete contact belonging to a user
func (ctx BasicDb) DeleteContactByUserID(id uint) (bool, error) {
	var contact Contact
	if err := ctx.db.Where("user_id = ?", id).Delete(&contact); err.Error != nil {
		logger.Warn.Printf("Could not delete contact belonging to user id: %d", id)
		return false, err.Error
	}
	return true, nil
}

// deletes a user and all related information based on ID
func (ctx BasicDb) DeleteUser(id uint) (bool, error) {
	user, err := ctx.GetUserByID(id)
	if err != nil {
		logger.Warn.Printf("Could not delete contact belonging to user id: %d", id)
	}

	if _, err := ctx.DeleteContactByUserID(id); err != nil {
		return false, err
	}

	if _, err := ctx.DeletePetsByUserID(id); err != nil {
		return false, err
	}

	if err := ctx.db.Delete(&user); err.Error != nil {
		logger.Warn.Printf("Could not delete user id: %d", id)
		return false, err.Error
	}
	return true, nil
}
