package model

import (
	"amusingx.fit/amusingx/mysqlstruct/ganymede"
	"context"
	"database/sql"
	"github.com/ItsWewin/superfactory/xerror"
	"github.com/jmoiron/sqlx"
)

func InsertOAuthInfo(ctx context.Context, tx *sqlx.Tx, oauth *ganymede.OauthInfo) (*ganymede.OauthInfo, *xerror.Error) {
	query := `INSERT INTO oauth_info (user_id, provider, access_token, code, outer_id, login, avatar_url, email)
	VALUES (:user_id, :provider, :access_token, :code, :outer_id, :login, :avatar_url, :email)
	ON DUPLICATE KEY UPDATE access_token = :access_token,
							code = :code,
							avatar_url = :avatar_url,
							email = :email
	`
	result, err := tx.NamedExecContext(ctx, query, oauth)
	if err != nil {
		return nil, xerror.NewError(err, xerror.Code.SSqlExecuteErr, "Unexpected error. ")
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, xerror.NewError(err, xerror.Code.SSqlExecuteErr, "Unexpected error. ")
	}
	oauth.ID = id

	return oauth, nil
}

func QueryOAuthInfoByProviderAndLogin(ctx context.Context, tx *sqlx.Tx, provider, login string) (*ganymede.OauthInfo, *xerror.Error) {
	query := `SELECT id, user_id, provider, access_token, code, outer_id, login, avatar_url, email, create_time, update_time
	FROM oauth_info
	WHERE provider = ? AND login = ?
	FOR UPDATE
	`

	var oauthInfo ganymede.OauthInfo
	err := tx.GetContext(ctx, &oauthInfo, query, provider, login)
	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, xerror.NewError(err, xerror.Code.SSqlExecuteErr, xerror.Message.SqlExecErr)
	}

	return &oauthInfo, nil
}
