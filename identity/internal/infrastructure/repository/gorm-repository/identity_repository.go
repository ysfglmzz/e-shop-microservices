package gormrepository

import (
	"github.com/google/uuid"
	"github.com/ysfglmzz/e-shop-microservices/identity/internal/app/model"
	"gorm.io/gorm"
)

type GormIdentityRepository struct {
	db *gorm.DB
}

func NewGormIdentityRepository(db *gorm.DB) *GormIdentityRepository {
	return &GormIdentityRepository{db: db}
}

func (g *GormIdentityRepository) AddUser(user *model.User) error {
	return g.db.Create(&user).Error
}

func (g *GormIdentityRepository) UpdateUser(user *model.User) error {
	return g.db.Save(&user).Error
}

func (g *GormIdentityRepository) DeleteUserById(id int) error {
	return g.db.Delete(&model.User{}, "id = ?", id).Error
}

func (g *GormIdentityRepository) GetUserById(id int) (*model.User, error) {
	var user model.User
	if err := g.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (g *GormIdentityRepository) IsTokenExist(uuid uuid.UUID) bool {
	if err := g.db.First(&model.TokenDetail{}, "uuid = ?", uuid).Error; err != nil {
		return false
	}
	return true
}

func (g *GormIdentityRepository) AddToken(tokenDetail *model.TokenDetail) error {
	return g.db.Create(&tokenDetail).Error
}

func (g *GormIdentityRepository) DeleteTokenByUUID(uuid uuid.UUID) error {
	return g.db.Delete(&model.TokenDetail{}, "uuid = ?", uuid).Error
}

func (g *GormIdentityRepository) GetUserRolesByUserId(id int) ([]string, error) {
	var roles []string

	if err := g.db.Model(&model.User{}).
		Joins("INNER JOIN user_roles AS ur ON ur.user_id = users.id").
		Joins("INNER JOIN roles AS r ON r.id = ur.role_id").
		Where("users.id = ?", id).
		Select("r.name").
		Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func (g *GormIdentityRepository) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	if err := g.db.First(&user, "email=?", email).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
