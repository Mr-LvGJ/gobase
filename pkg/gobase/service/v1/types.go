package v1

import (
	"context"
	v1 "github.com/Mr-LvGJ/gobase/pkg/gobase/model/v1"
)

type UserSrv interface {
	Create(ctx context.Context, user *v1.User) error
	Update(ctx context.Context, user *v1.User) error
	Delete(ctx context.Context, username string) error
	Get(ctx context.Context, username string) (*v1.User, error)
	List(ctx context.Context) (*v1.UserList, error)
}

type Service interface {
	Users() UserSrv
}
