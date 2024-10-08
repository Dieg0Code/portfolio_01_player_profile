package impl

import (
	"github.com/dieg0code/player-profile/src/data/request"
	"github.com/dieg0code/player-profile/src/data/response"
	"github.com/dieg0code/player-profile/src/helpers"
	"github.com/dieg0code/player-profile/src/models"
	"github.com/dieg0code/player-profile/src/repository"
	"github.com/dieg0code/player-profile/src/services"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type AchievementServiceImpl struct {
	AchievementRepository repository.AchievementRepository
	Validate              *validator.Validate
}

// GetAchievementWithPlayers implements services.AchievementService.
func (a *AchievementServiceImpl) GetAchievementWithPlayers(achievementID uint) (*response.AchievementWithPlayers, error) {

	achievement, err := a.AchievementRepository.GetAchievementWithPlayers(achievementID)

	if err != nil {
		logrus.WithError(err).Error("[AchievementServiceImpl.GetAchievementWithPlayers] Failed to get achievement with players")
		return nil, helpers.ErrAchievementRepository
	}

	achievementWithPlayers := response.AchievementWithPlayers{
		ID:   achievement.ID,
		Name: achievement.Name,
	}

	for _, player := range achievement.PlayerProfiles {
		achievementWithPlayers.Players = append(achievementWithPlayers.Players, response.PlayerSumary{
			ID:       player.ID,
			Nickname: player.Nickname,
		})
	}

	return &achievementWithPlayers, nil
}

// Create implements services.AchievementService.
func (a *AchievementServiceImpl) Create(achievement request.CreateAchievementRequest) error {
	err := a.Validate.Struct(achievement)
	if err != nil {
		logrus.WithError(err).Error("[AchievementServiceImpl.Create] Failed to validate achievement data")
		return helpers.ErrAchievementDataValidation
	}

	achievementModel := models.Achievement{
		Name:        achievement.Name,
		Description: achievement.Description,
	}

	err = a.AchievementRepository.CreateAchievement(&achievementModel)
	if err != nil {
		logrus.WithError(err).Error("[AchievementServiceImpl.Create] Failed to create achievement")
		return helpers.ErrAchievementRepository
	}

	return nil
}

// Delete implements services.AchievementService.
func (a *AchievementServiceImpl) Delete(achievementID uint) error {
	if achievementID == 0 {
		return helpers.ErrInvalidAchievementID

	}

	err := a.AchievementRepository.DeleteAchievement(achievementID)
	if err != nil {
		logrus.WithError(err).Error("[AchievementServiceImpl.Delete] Failed to delete achievement")
		return helpers.ErrAchievementRepository
	}

	return nil
}

// GetAll implements services.AchievementService.
func (a *AchievementServiceImpl) GetAll(page int, pageSize int) ([]response.AchievementResponse, error) {
	if page <= 0 || pageSize <= 0 {
		return nil, helpers.ErrInvalidPagination
	}

	offset := (page - 1) * pageSize

	achievements, err := a.AchievementRepository.GetAllAchievements(offset, pageSize)
	if err != nil {
		logrus.WithError(err).Error("[AchievementServiceImpl.GetAll] Failed to get all achievements")
		return nil, helpers.ErrAchievementRepository
	}

	var achievementResponses []response.AchievementResponse
	for _, achievement := range achievements {
		achievementResponse := response.AchievementResponse{
			ID:          achievement.ID,
			Name:        achievement.Name,
			Description: achievement.Description,
		}

		err = a.Validate.Struct(achievementResponse)
		if err != nil {
			logrus.WithError(err).Error("[AchievementServiceImpl.GetAll] Failed to validate achievement data")
			return nil, helpers.ErrAchievementDataValidation
		}

		achievementResponses = append(achievementResponses, achievementResponse)
	}

	return achievementResponses, nil
}

// GetByID implements services.AchievementService.
func (a *AchievementServiceImpl) GetByID(achievementID uint) (*response.AchievementResponse, error) {
	if achievementID == 0 {
		return nil, helpers.ErrInvalidAchievementID
	}

	achievement, err := a.AchievementRepository.GetAchievement(achievementID)
	if err != nil {
		logrus.WithError(err).Error("[AchievementServiceImpl.GetByID] Failed to get achievement")
		return nil, helpers.ErrAchievementRepository
	}

	if achievement == nil {
		return nil, helpers.ErrAchievementNotFound
	}

	achievementResponse := response.AchievementResponse{
		ID:          achievement.ID,
		Name:        achievement.Name,
		Description: achievement.Description,
	}

	err = a.Validate.Struct(achievementResponse)
	if err != nil {
		logrus.WithError(err).Error("[AchievementServiceImpl.GetByID] Failed to validate achievement data")
		return nil, helpers.ErrAchievementDataValidation
	}

	return &achievementResponse, nil
}

// Update implements services.AchievementService.
func (a *AchievementServiceImpl) Update(achievementID uint, achievement request.UpdateAchievementRequest) error {
	if achievementID == 0 {
		return helpers.ErrInvalidAchievementID
	}

	err := a.Validate.Struct(achievement)
	if err != nil {
		logrus.WithError(err).Error("[AchievementServiceImpl.Update] Failed to validate achievement data")
		return helpers.ErrAchievementDataValidation
	}

	achievementModel, err := a.AchievementRepository.GetAchievement(achievementID)
	if err != nil {
		logrus.WithError(err).Error("[AchievementServiceImpl.Update] Failed to get achievement")
		return helpers.ErrAchievementRepository
	}

	if achievementModel == nil {
		return helpers.ErrAchievementNotFound
	}

	achievementModel.Name = achievement.Name
	achievementModel.Description = achievement.Description

	err = a.AchievementRepository.UpdateAchievement(achievementID, achievementModel)
	if err != nil {
		logrus.WithError(err).Error("[AchievementServiceImpl.Update] Failed to update achievement")
		return helpers.ErrAchievementRepository
	}

	return nil
}

func NewAchievementServiceImpl(achievementRepository repository.AchievementRepository, validate *validator.Validate) services.AchievementService {
	return &AchievementServiceImpl{
		AchievementRepository: achievementRepository,
		Validate:              validate,
	}
}
