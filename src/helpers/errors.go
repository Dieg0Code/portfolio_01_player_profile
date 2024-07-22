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

// Achievement errors.
var ErrorAchievementNotFound = errors.New("achievement not found")
var ErrorUpdateAchievement = errors.New("error updating achievement")
var ErrorDeletingAchievement = errors.New("error deleting achievement")

// Services

// User errors.

var ErrInvalidUserID = errors.New("invalid user id")
var ErrUserDataValidation = errors.New("user data validation error")
