package store

import (
	"context"
	"errors"
	"github.com/Mr-LvGJ/gobase/pkg/common/fields"
	"github.com/Mr-LvGJ/gobase/pkg/common/util/gormutil"
	metav1 "github.com/Mr-LvGJ/gobase/pkg/gobase/meta/v1"

	"gorm.io/gorm"

	v1 "github.com/Mr-LvGJ/gobase/pkg/gobase/model/v1"
)

type users struct {
	db *gorm.DB
}

func newUsers(ds *datastore) *users {
	return &users{ds.db}

}

func (u *users) Create(ctx context.Context, user *v1.User) error {
	return u.db.Create(&user).Error
}

func (u *users) Update(ctx context.Context, user *v1.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *users) Delete(ctx context.Context, username string) error {
	//TODO implement me
	panic("implement me")
}

func (u *users) Get(ctx context.Context, username string) (*v1.User, error) {
	user := &v1.User{}
	err := u.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}
	return user, nil
}

func (u *users) List(ctx context.Context, options metav1.ListOptions) (*v1.UserList, error) {
	ret := &v1.UserList{}
	ol := gormutil.Unpointer(options.Limit, options.Offset)
	selector, _ := fields.ParseSelector(options.FieldSelector)
	username, _ := selector.RequiresExactMatch("username")
	d := u.db.Where("username like ?", "%"+username+"%").
		Offset(ol.Offset).
		Limit(ol.Limit).
		Order("id desc").
		Find(&ret.Items)
	return ret, d.Error
}
