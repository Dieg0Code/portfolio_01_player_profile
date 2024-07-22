package impl

import (
	"github.com/dieg0code/player-profile/src/data/request"
	"github.com/dieg0code/player-profile/src/data/response"
	"github.com/dieg0code/player-profile/src/helpers"
	"github.com/dieg0code/player-profile/src/models"
	"github.com/dieg0code/player-profile/src/repository"
	"github.com/dieg0code/player-profile/src/services"
	"github.com/go-playground/validator/v10"
)

type PlayerProfileServiceImpl struct {
	PlayerProfileRepository repository.PlayerProfileRepository
	Validate                *validator.Validate
	PasswordHasher          services.PasswordHasher
}

// Create implements services.PlayerProfileService.
func (p *PlayerProfileServiceImpl) Create(playerProfile request.CreatePlayerProfileRequest) error {
	err := p.Validate.Struct(playerProfile)
	if err != nil {
		return helpers.ErrPlayerProfileDataValidation
	}

	playerProfileModel := models.PlayerProfile{
		Nickname:   playerProfile.Nickname,
		Avatar:     playerProfile.Avatar,
		Level:      playerProfile.Level,
		Experience: playerProfile.Experience,
		Points:     playerProfile.Points,
		UserID:     playerProfile.UserID,
	}

	err = p.PlayerProfileRepository.CreatePlayerProfile(&playerProfileModel)
	if err != nil {
		return helpers.ErrRepository
	}

	return nil
}

// Delete implements services.PlayerProfileService.
func (p *PlayerProfileServiceImpl) Delete(playerProfileID uint) error {
	if playerProfileID == 0 {
		return helpers.ErrInvalidPlayerProfileID
	}

	err := p.PlayerProfileRepository.DeletePlayerProfile(playerProfileID)
	if err != nil {
		return err
	}

	return nil
}

// GetAll implements services.PlayerProfileService.
func (p *PlayerProfileServiceImpl) GetAll() ([]response.UserResponse, error) {
	panic("unimplemented")
}

// GetByID implements services.PlayerProfileService.
func (p *PlayerProfileServiceImpl) GetByID(playerProfileID uint) (*response.PlayerProfileResponse, error) {
	if playerProfileID == 0 {
		return nil, helpers.ErrInvalidPlayerProfileID
	}

	playerProfile, err := p.PlayerProfileRepository.GetPlayerProfile(playerProfileID)
	if err != nil {
		return nil, err
	}

	playerProfileResponse := response.PlayerProfileResponse{
		ID:         playerProfile.ID,
		Nickname:   playerProfile.Nickname,
		Avatar:     playerProfile.Avatar,
		Level:      playerProfile.Level,
		Experience: playerProfile.Experience,
		Points:     playerProfile.Points,
		UserID:     playerProfile.UserID,
	}

	return &playerProfileResponse, nil

}

// Update implements services.PlayerProfileService.
func (p *PlayerProfileServiceImpl) Update(playerProfileID uint, playerProfile request.UpdatePlayerProfileRequest) error {

	playerData, err := p.PlayerProfileRepository.GetPlayerProfile(playerProfileID)
	if playerData == nil {
		return helpers.ErrorPlayerProfileNotFound
	}

	if err != nil {
		return err
	}

	err = p.Validate.Struct(playerProfile)

	if err != nil {
		return helpers.ErrPlayerProfileDataValidation
	}

	playerData.Nickname = playerProfile.Nickname
	playerData.Avatar = playerProfile.Avatar
	playerData.Level = playerProfile.Level
	playerData.Experience = playerProfile.Experience
	playerData.Points = playerProfile.Points

	err = p.PlayerProfileRepository.UpdatePlayerProfile(playerProfileID, playerData)
	if err != nil {
		return err
	}

	return nil
}

func NewPlayerProfileServiceImpl(playerProfileRepository repository.PlayerProfileRepository, validate *validator.Validate) services.PlayerProfileService {
	return &PlayerProfileServiceImpl{
		PlayerProfileRepository: playerProfileRepository,
		Validate:                validate,
	}
}
