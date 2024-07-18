package helpers

import "errors"

// User errors.
var ErrorUserNotFound = errors.New("user not found")
var ErrorUpdateUser = errors.New("error updating user")
var ErrorDeleteUser = errors.New("error deleting user")

// Player Profile errors.
var ErrorPlayerProfileNotFound = errors.New("player profile not found")
var ErrorUpdatePlayer = errors.New("error updating player")
var ErrorDeletingUser = errors.New("error deleting user")
