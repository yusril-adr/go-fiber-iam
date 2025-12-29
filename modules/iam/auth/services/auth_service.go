package services

import (
	"time"

	"github.com/google/uuid"

	"iam-service/infrastructure/config"
	"iam-service/infrastructure/errors"
	"iam-service/infrastructure/types"
	"iam-service/infrastructure/utils"
	"iam-service/models/maindb/iam"
	"iam-service/modules/iam/auth/dtos/params"
	"iam-service/modules/iam/auth/dtos/results"
	"iam-service/modules/iam/auth/messages"
	authRepository "iam-service/modules/iam/auth/repositories"
	userResults "iam-service/modules/iam/user/dtos/results"
	userRepository "iam-service/modules/iam/user/repositoires"
)

func SignIn(dto *params.AuthSignInParam) *results.AuthSignInResult {
	user := userRepository.FindByEmail(dto.Email)

	if user == nil {
		panic(errors.BadRequest(messages.ErrAuthFailed))
	}

	isPasswordValid := utils.CheckPassword(dto.Password, user.Password)
	if !isPasswordValid {
		panic(errors.BadRequest(messages.ErrAuthFailed))
	}

	accessToken, accessTokenExpiredAt := generateAccessToken(user)
	refreshToken, refreshTokenExpiredAt := generateRefreshToken(user)

	return &results.AuthSignInResult{
		AccessToken: results.AuthTokenResult{
			Token:     accessToken,
			ExpiresAt: accessTokenExpiredAt,
		},
		RefreshToken: results.AuthTokenResult{
			Token:     refreshToken,
			ExpiresAt: refreshTokenExpiredAt,
		},
	}
}

func generateUserTokenPayload(user *iam.User) types.TUserPayload {
	userRoleIds := utils.Map(user.Roles, func(role iam.Role) uuid.UUID {
		return role.Id
	})
	userPayload := types.TUserPayload{
		Id:      user.Id,
		RoleIds: userRoleIds,
	}
	return userPayload
}

func generateAccessToken(user *iam.User) (string, time.Time) {
	userPayload := generateUserTokenPayload(user)
	accessTokenExpiredIn, err := time.ParseDuration(config.JWT_ACCESS_TOKEN_EXPIRES_IN)
	if err != nil {
		panic(err)
	}
	accessTokenExpiredAt := time.Now().Add(accessTokenExpiredIn)
	accessToken := utils.CreateToken(userPayload, config.JWT_ACCESS_TOKEN_SECRET, accessTokenExpiredAt.Unix())
	return accessToken, accessTokenExpiredAt
}

func generateRefreshToken(user *iam.User) (string, time.Time) {
	userPayload := generateUserTokenPayload(user)
	refreshTokenExpiredIn, err := time.ParseDuration(config.JWT_REFRESH_TOKEN_EXPIRES_IN)
	if err != nil {
		panic(err)
	}
	refreshTokenExpiredAt := time.Now().Add(refreshTokenExpiredIn)
	refreshToken := utils.CreateToken(userPayload, config.JWT_REFRESH_TOKEN_SECRET, refreshTokenExpiredAt.Unix())

	refreshTokenModel := iam.UserToken{
		UserId:    user.Id,
		Token:     refreshToken,
		ExpiresAt: refreshTokenExpiredAt,
	}
	authRepository.Create(&refreshTokenModel, nil)
	return refreshToken, refreshTokenExpiredAt
}

func Profile(accessToken string) *userResults.UserResult {
	payload := parsingToken(accessToken, config.JWT_ACCESS_TOKEN_SECRET)

	user := userRepository.FindById(payload.Id)
	if user == nil {
		panic(errors.NotFound(messages.ErrUserTokenNotFoundOrDeleted))
	}

	result := &userResults.UserResult{}
	result.MapModel(*user)
	return result
}

func parsingToken(token string, secret string) *types.TUserPayload {
	defer func() {
		if recover() != nil {
			panic(errors.Unauthorized(messages.ErrTokenInvalidOrExpired))
		}
	}()

	return utils.ParseToken[*types.TUserPayload](token, secret)
}

func RenewToken(refreshToken string) *results.AuthTokenResult {
	payload := parsingToken(refreshToken, config.JWT_REFRESH_TOKEN_SECRET)

	userToken := authRepository.GetUserToken(payload.Id, refreshToken)
	if userToken == nil {
		panic(errors.Unauthorized(messages.ErrTokenInvalidOrExpired))
	}

	user := userRepository.FindById(payload.Id)
	if user == nil {
		panic(errors.NotFound(messages.ErrUserTokenNotFoundOrDeleted))
	}

	accessToken, accessTokenExpiredAt := generateAccessToken(user)

	return &results.AuthTokenResult{
		Token:     accessToken,
		ExpiresAt: accessTokenExpiredAt,
	}
}

func SignOut(refreshToken string) {
	payload := parsingToken(refreshToken, config.JWT_REFRESH_TOKEN_SECRET)

	userToken := authRepository.GetUserToken(payload.Id, refreshToken)
	if userToken == nil {
		panic(errors.Unauthorized(messages.ErrTokenInvalidOrExpired))
	}

	authRepository.DeleteTokenWithUserId(refreshToken, payload.Id, nil)
}

func ClearExpiredToken() {
	today := time.Now()
	authRepository.DeleteTokenWithExpiredAt(today, nil)
}
