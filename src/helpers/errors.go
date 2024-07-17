package helpers

import "errors"

var ErrorUserNotFound = errors.New("user not found")
var ErrorUpdateUser = errors.New("error updating user")
var ErrorDeleteUser = errors.New("error deleting user")
