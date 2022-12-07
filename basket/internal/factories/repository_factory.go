package factories

import (
	"github.com/ysfglmzz/e-shop-microservices/basket/config"
	"github.com/ysfglmzz/e-shop-microservices/basket/internal/app/repository"
	mongoRepository "github.com/ysfglmzz/e-shop-microservices/basket/internal/infrastructure/repository/mongo-repository"
)

type RepositoryFactory struct {
	cfg               config.AppConfig
	connectionFactory ConnectionFactory
}

func NewRepositoryFactory(cfg config.AppConfig, connectionFactory ConnectionFactory) *RepositoryFactory {
	return &RepositoryFactory{cfg: cfg, connectionFactory: connectionFactory}
}

func (r *RepositoryFactory) GetBasketRepository() repository.IBasketRepository {
	switch r.cfg.System.DbDriver {
	case "mongo":
		return mongoRepository.NewMongoBasketRepository(r.connectionFactory.GetMongoDb().Collection("baskets"))
	}
	return nil
}
