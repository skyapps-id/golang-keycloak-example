package service

import (
	"context"
	"golang-keycloak/dto"
	"golang-keycloak/pkg"

	"github.com/Nerzal/gocloak/v11"
)

type (
	UserSerivce interface {
		LoginAdmin(ctx context.Context, req dto.ReqUserLogin) (*dto.ResUserLogin, error)
		Create(ctx context.Context, token string) (*dto.ResUserRegister, error)
		Login(ctx context.Context, req dto.ReqUserLogin) (*dto.ResUserLogin, error)
		RefreshToken(ctx context.Context, req dto.ReqUserRefreshTokenOrLogout) (*dto.ResUserLogin, error)
		Logout(ctx context.Context, req dto.ReqUserRefreshTokenOrLogout) (*dto.ResUserLogout, error)
		Info(ctx context.Context, token string) (*gocloak.UserInfo, error)
	}

	userImpl struct {
		keycloak pkg.Keycloak
	}
)

func NewUserService(keycloak pkg.Keycloak) UserSerivce {
	return userImpl{keycloak: keycloak}
}

func (s userImpl) LoginAdmin(ctx context.Context, req dto.ReqUserLogin) (*dto.ResUserLogin, error) {
	token, err := s.keycloak.Gocloak.LoginAdmin(ctx, req.Username, req.Password, s.keycloak.Realm)
	if err != nil {
		return nil, err
	}

	return &dto.ResUserLogin{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		ExpiresIn:    token.ExpiresIn,
	}, nil
}

func (s userImpl) Create(ctx context.Context, token string) (*dto.ResUserRegister, error) {
	user := gocloak.User{
		FirstName: gocloak.StringP("Bob"),
		LastName:  gocloak.StringP("Uncle"),
		Email:     gocloak.StringP("something@really.wrong"),
		Enabled:   gocloak.BoolP(true),
		Username:  gocloak.StringP("CoolGuy"),
	}
	_, err := s.keycloak.Gocloak.CreateUser(ctx, token, s.keycloak.Realm, user)
	if err != nil {
		return nil, err
	}

	return &dto.ResUserRegister{
		FirstName: *user.FirstName,
		LastName:  *user.LastName,
		Email:     *user.Email,
		Username:  *user.Username,
	}, nil
}

func (s userImpl) Login(ctx context.Context, req dto.ReqUserLogin) (*dto.ResUserLogin, error) {
	token, err := s.keycloak.Gocloak.Login(ctx, s.keycloak.ClientID, s.keycloak.ClientSecret, s.keycloak.Realm, req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	return &dto.ResUserLogin{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		ExpiresIn:    token.ExpiresIn,
	}, nil
}

func (s userImpl) RefreshToken(ctx context.Context, req dto.ReqUserRefreshTokenOrLogout) (*dto.ResUserLogin, error) {
	token, err := s.keycloak.Gocloak.RefreshToken(ctx, req.RefreshToken, s.keycloak.ClientID, s.keycloak.ClientSecret, s.keycloak.Realm)
	if err != nil {
		return nil, err
	}

	return &dto.ResUserLogin{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		ExpiresIn:    token.ExpiresIn,
	}, nil
}

func (s userImpl) Logout(ctx context.Context, req dto.ReqUserRefreshTokenOrLogout) (*dto.ResUserLogout, error) {
	err := s.keycloak.Gocloak.Logout(ctx, s.keycloak.ClientID, s.keycloak.ClientSecret, s.keycloak.Realm, req.RefreshToken)
	if err != nil {
		return nil, err
	}

	return &dto.ResUserLogout{
		Status:  "success",
		Massage: "user logout",
	}, nil
}

func (s userImpl) Info(ctx context.Context, token string) (*gocloak.UserInfo, error) {

	user, err := s.keycloak.Gocloak.GetUserInfo(ctx, token, s.keycloak.Realm)
	if err != nil {
		return nil, err
	}

	return user, nil
}
