package dto

type (
	ResUserRegister struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Email     string `json:"email"`
		Username  string `json:"username"`
	}

	ReqUserLogin struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	ReqUserRefreshTokenOrLogout struct {
		RefreshToken string `json:"refreshToken" validate:"required"`
	}

	ResUserLogin struct {
		AccessToken  string `json:"accessToken"`
		RefreshToken string `json:"refreshToken"`
		ExpiresIn    int    `json:"expiresIn"`
	}

	ResUserLogout struct {
		Status  string `json:"status"`
		Massage string `json:"massage"`
	}
)
