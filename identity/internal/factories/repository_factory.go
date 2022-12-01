package factories

import (
	"github.com/ysfglmzz/e-shop-microservices/identity/config"
	"github.com/ysfglmzz/e-shop-microservices/identity/internal/app/repository"
	gormrepository "github.com/ysfglmzz/e-shop-microservices/identity/internal/infrastructure/repository/gorm-repository"
)

type RepositoryFactory struct {
	cfg               config.AppConfig
	connectionFactory ConnectionFactory
}

func NewRepositoryFactory(cfg config.AppConfig, connectionFactory ConnectionFactory) *RepositoryFactory {
	return &RepositoryFactory{cfg: cfg, connectionFactory: connectionFactory}
}

func (r *RepositoryFactory) GetIdentityRepository() repository.IIdentityRepository {
	switch r.cfg.System.DbManager {
	case "gorm":
		return gormrepository.NewGormIdentityRepository(r.connectionFactory.GetGormDb())
	}
	return nil
}
