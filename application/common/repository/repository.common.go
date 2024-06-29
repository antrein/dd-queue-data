package repository

import (
	"antrein/dd-queue-data/application/common/resource"
	"antrein/dd-queue-data/internal/repository/config"
	"antrein/dd-queue-data/internal/repository/room"
	cfg "antrein/dd-queue-data/model/config"
)

type CommonRepository struct {
	ConfigRepo *config.Repository
	RoomRepo   *room.Repository
}

func NewCommonRepository(cfg *cfg.Config, rsc *resource.CommonResource) (*CommonRepository, error) {
	configRepo := config.New(cfg, rsc.Redis, rsc.GRPC)
	roomRepo := room.New(cfg, rsc.Redis)

	commonRepo := CommonRepository{
		ConfigRepo: configRepo,
		RoomRepo:   roomRepo,
	}
	return &commonRepo, nil
}
