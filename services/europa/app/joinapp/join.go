package joinapp

import (
	"amusingx.fit/amusingx/apistruct/europa"
	"amusingx.fit/amusingx/services/europa/model"
	"amusingx.fit/amusingx/services/europa/xredis/lock"
	"context"
	"fmt"
	"github.com/ItsWewin/superfactory/logger"
	"github.com/ItsWewin/superfactory/xerror"
)

func CreateUser(ctx context.Context, u *europa.JoinRequest) (*model.User, *xerror.Error) {
	var (
		err  *xerror.Error
		user = &model.User{
			Nickname: u.Nickname,
			Phone:    fmt.Sprintf("%s-%s", u.AreaCode, u.Phone),
			Password: u.Password,
		}
	)

	defer clearPassword(u, user)

	// 获取锁
	key := fmt.Sprintf("%s-%s", user.Nickname, user.Phone)
	locker := lock.NewLocker(key, user.Nickname, 30)
	if err := locker.Lock(ctx); err != nil {
		logger.Errorf("\nlock failed: %s\n", err)

		return nil, xerror.NewError(err, xerror.Code.SGetLockeErr, "Try again. ")
	}
	// 释放锁
	defer locker.Unlock(ctx)

	existed, err := user.ExistedWithNicknameOrPhone(ctx)
	if err != nil {
		return nil, err
	}

	if existed {
		return nil, xerror.NewError(err, xerror.Code.BDataIsNotAllow, "Phone or nickname is taken. ")
	}

	user, err = model.Create(ctx, user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func clearPassword(r *europa.JoinRequest, u *model.User) {
	r.Password = ""
	u.Password = ""
	u.PasswordDigest = ""
}
