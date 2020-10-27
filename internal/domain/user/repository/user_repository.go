package repository

import (
	"context"

	"github.com/donech/nirvana/internal/domain/user/entity"

	"github.com/donech/nirvana/internal/conn"

	"github.com/donech/tool/xdb"
)

func NewUserRepository(db conn.NirvanaDB) *UserRepository {
	return &UserRepository{Repository: xdb.Repository{DB: db}}
}

type UserRepository struct {
	xdb.Repository
}

func (r UserRepository) ItemByID(ctx context.Context, id int64) (e entity.User, err error) {
	err = xdb.Trace(ctx, r.DB).Where("id = ?", id).First(&e).Error
	return e, err
}

func (r UserRepository) ItemByEmail(ctx context.Context, email string) (e entity.User, err error) {
	err = xdb.Trace(ctx, r.DB).Where("email = ?", email).First(&e).Error
	return e, err
}

func (r UserRepository) ItemsByCursor(ctx context.Context, cursor, size int64) (e []entity.User, err error) {
	err = xdb.Trace(ctx, r.DB).Where("id >= ?", cursor).Limit(size).Find(&e).Error
	return e, err
}

func (r UserRepository) ItemsByCursorReverse(ctx context.Context, cursor, size int64) (e []entity.User, err error) {
	db := xdb.Trace(ctx, r.DB).Order("id desc").Limit(size)
	if cursor > 0 {
		db = db.Where("id <= ?", cursor)
	}
	err = db.Find(&e).Error
	return e, err
}

func (r UserRepository) Create(ctx context.Context, data map[string]interface{}) (user entity.User, err error) {
	user = entity.User{
		Name:  data["name"].(string),
		Phone: data["phone"].(string),
	}
	return user, xdb.Trace(ctx, r.DB).Create(&user).Error
}

func (r UserRepository) Update(ctx context.Context, id int64, data map[string]interface{}) error {
	return xdb.Trace(ctx, r.DB).Model(&entity.User{}).Where("id = ?", id).Updates(data).Error
}

func (r UserRepository) Delete(ctx context.Context, id int64) error {
	return xdb.Trace(ctx, r.DB).Where("id = ?", id).Delete(entity.User{}).Error
}

func (r UserRepository) Migration() {
	r.DB.AutoMigrate(entity.User{})
}
