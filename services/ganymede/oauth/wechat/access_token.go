package wechat

import (
	"amusingx.fit/amusingx/apistruct/wechat"
	"amusingx.fit/amusingx/services/ganymede/oauth/oauthstruct"
	"github.com/ItsWewin/superfactory/httputil"
	"github.com/ItsWewin/superfactory/httputil/rest"
	"github.com/ItsWewin/superfactory/xerror"
	"net/url"
)

type OAuth struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
}

func New(clientID, clientSecret, authorizationCode string) *OAuth {
	return &OAuth{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		GrantType:    authorizationCode,
	}
}

func (o *OAuth) GetAccessToken(accessTokenUrl, code string) (*oauthstruct.AccessToken, *xerror.Error) {
	var params = url.Values{}
	params.Set("appid", o.ClientID)
	params.Set("secret", o.ClientSecret)
	params.Set("code", code)
	params.Set("grant_type", code)

	u, err := url.Parse(accessTokenUrl)
	if err != nil {
		return nil, xerror.NewErrorf(nil, xerror.Code.BUnexpectedData, "accessTokenUrl format err")
	}
	u.RawQuery = params.Encode()

	opts := func(opts *rest.Options) {
		opts.Header = map[string]string{
			httputil.HeaderContent: httputil.JsonHeaderContent,
			httputil.HeaderAccept:  httputil.JsonHeaderAccept,
		}
	}

	var dest wechat.AccessTokenResponse
	xErr := rest.Get(u.String(), &dest, opts)
	if xErr != nil {
		return nil, xErr
	}

	if dest.IsError() {
		return nil, xerror.NewErrorf(nil, xerror.Code.OtherNetworkError, dest.Errmsg)
	}

	return &oauthstruct.AccessToken{
		AccessToken:  dest.AccessToken,
		Scope:        dest.Scope,
		ExpiresIn:    dest.ExpiresIn,
		RefreshToken: dest.RefreshToken,
		OpenID:       dest.Openid,
		UnionID:      dest.Unionid,
	}, nil
}