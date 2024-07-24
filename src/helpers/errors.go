package helpers

import "errors"

// Models

// User errors.
var ErrorUserNotFound = errors.New("user not found")
var ErrorUpdateUser = errors.New("error updating user")
var ErrorDeleteUser = errors.New("error deleting user")
var ErrorGetAllUsers = errors.New("error getting all users")

// Player Profile errors.
var ErrorPlayerProfileNotFound = errors.New("player profile not found")
var ErrorUpdatePlayer = errors.New("error updating player")
var ErrorDeletingUser = errors.New("error deleting user")
var ErrorGetAllPlayerProfiles = errors.New("error getting all player profiles")

// Achievement errors.
var ErrorAchievementNotFound = errors.New("achievement not found")
var ErrorUpdateAchievement = errors.New("error updating achievement")
var ErrorDeletingAchievement = errors.New("error deleting achievement")

// Services

// User errors.

var ErrInvalidUserID = errors.New("invalid user id")
var ErrUserDataValidation = errors.New("user data validation error")

// Player Profile errors.
var ErrInvalidPlayerProfileID = errors.New("invalid player profile id")
var ErrPlayerProfileDataValidation = errors.New("player profile data validation error")

var ErrRepository = errors.New("error in repository")

// Achievement errors.
var ErrAchievementDataValidation = errors.New("achievement data validation error")
var ErrInvalidAchievementID = errors.New("invalid achievement id")
var ErrAchievementNotFound = errors.New("achievement not found")
var ErrAchievementRepository = errors.New("error in achievement repository")
