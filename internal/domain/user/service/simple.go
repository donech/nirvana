package service

import (
	"context"

	"github.com/donech/nirvana/internal/domain/user/entity"

	"github.com/donech/nirvana/internal/conn"

	"github.com/donech/tool/xdb"
	"github.com/jinzhu/gorm"
)

func NewSimpleService(db conn.NirvanaDB) *SimpleService {
	return &SimpleService{db: db}
}

type SimpleService struct {
	db *gorm.DB
}

func (s SimpleService) ItemByID(ctx context.Context, id int64) (e entity.User, err error) {
	err = xdb.Trace(ctx, s.db).Where("id = ?", id).First(&e).Error
	return e, err
}

func (s SimpleService) ItemsByCursor(ctx context.Context, cursor, size int64) (e []entity.User, err error) {
	err = xdb.Trace(ctx, s.db).Where("id >= ?", cursor).Limit(size).Find(&e).Error
	return e, err
}

func (s SimpleService) Create(ctx context.Context, data map[string]interface{}) (user entity.User, err error) {
	user = entity.User{
		Name:  data["name"].(string),
		Phone: data["phone"].(string),
	}
	return user, xdb.Trace(ctx, s.db).Create(&user).Error
}

func (s SimpleService) Update(ctx context.Context, id int64, data map[string]interface{}) error {
	return xdb.Trace(ctx, s.db).Model(&entity.User{}).Where("id = ?", id).Updates(data).Error
}

func (s SimpleService) Delete(ctx context.Context, id int64) error {
	return xdb.Trace(ctx, s.db).Where("id = ?", id).Delete(entity.User{}).Error
}

func (s SimpleService) Migration() {
	s.db.AutoMigrate(entity.User{})
}
