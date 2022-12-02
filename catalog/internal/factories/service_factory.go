package factories

import (
	service "github.com/ysfglmzz/e-shop-microservices/catalog/internal/app/service"
	serviceImp "github.com/ysfglmzz/e-shop-microservices/catalog/internal/infrastructure/service"
)

type ServiceFactory struct {
	repositoryFactory RepositoryFactory
}

func NewServiceFactory(repositoryFactory RepositoryFactory) *ServiceFactory {
	return &ServiceFactory{repositoryFactory: repositoryFactory}
}

func (s *ServiceFactory) GetProductService() service.IProductService {
	return serviceImp.NewProductService(s.repositoryFactory.GetProductRepository())
}
