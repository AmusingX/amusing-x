package oauth

import (
	"amusingx.fit/amusingx/apistruct/europa"
	ganymedeservice "amusingx.fit/amusingx/protos/ganymede/service/ganymede/proto"
	"amusingx.fit/amusingx/services/europa/rpcclient/ganymede"
	"context"
	"github.com/ItsWewin/superfactory/aerror"
	"net/url"
)

func GetOAuthInfo(ctx context.Context, provider, service string) (*europa.OAuthInfo, aerror.Error) {
	in := &ganymedeservice.OAuthInfoRequest{Provider: provider, Service: service}

	resp, err := ganymede.RPCClient.Client.OAuthInfo(ctx, in)
	if err != nil {
		return nil, aerror.NewErrorf(nil, aerror.Code.SUnexpectedErr, err.Error())
	}

	u, err := url.Parse(resp.OauthUrl)
	if err != nil {
		return nil, aerror.NewError(nil, aerror.Code.BUnexpectedData, err.Error())
	}

	_, err = url.Parse(resp.RedirectUrl)
	if err != nil {
		return nil, aerror.NewError(nil, aerror.Code.BUnexpectedData, err.Error())
	}

	q := u.Query()
	q.Set("client_id", resp.ClientId)
	q.Set("redirect_uri", resp.RedirectUrl)
	q.Set("scope", resp.Scope)
	q.Set("state", resp.State)
	q.Set("grant_type", resp.GrantType)
	u.RawQuery = q.Encode()

	return &europa.OAuthInfo{
		Provider:     resp.Provider,
		OauthUrl:     resp.OauthUrl,
		ClientID:     resp.ClientId,
		Scope:        resp.Scope,
		State:        resp.State,
		GrantType:    resp.GrantType,
		RedirectUrl:  resp.RedirectUrl,
		CompletePath: u.String(),
	}, nil
}
