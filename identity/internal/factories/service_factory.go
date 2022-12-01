package factories

import (
	"github.com/ysfglmzz/e-shop-microservices/identity/config"
	usecase "github.com/ysfglmzz/e-shop-microservices/identity/internal/app/use-case"
	usecaseImp "github.com/ysfglmzz/e-shop-microservices/identity/internal/infrastructure/use-case"
)

type ServiceFactory struct {
	cfg               config.AppConfig
	connectionFactory ConnectionFactory
	repositoryFactory RepositoryFactory
}

func NewServiceFactory(
	connectionFactory ConnectionFactory,
	repositoryFactory RepositoryFactory,
	cfg config.AppConfig,
) *ServiceFactory {
	return &ServiceFactory{
		connectionFactory: connectionFactory,
		repositoryFactory: repositoryFactory,
		cfg:               cfg,
	}
}

func (s *ServiceFactory) GetIdentityService() usecase.IIdentityService {
	return usecaseImp.NewIdentityService(s.GetMessageService(), s.repositoryFactory.GetIdentityRepository())
}

func (s *ServiceFactory) GetMessageService() usecase.IMessageService {
	switch s.cfg.System.MessageBus {
	case "rabbitMq":
		return usecaseImp.NewRabbitMqMessageService(s.cfg.RabbitMq, s.connectionFactory.GetRabbitMqConnection())
	}
	return nil
}
