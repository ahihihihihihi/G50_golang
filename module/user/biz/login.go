package userbiz

import (
	"G05-food-delivery/common"
	"G05-food-delivery/component/tokenprovider"
	usermodel "G05-food-delivery/module/user/model"
	"context"
)

type LoginStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

type loginBusiness struct {
	//appCtx appctx.AppContext
	storeUser LoginStorage
	tokenProvider tokenprovider.Provider
	hasher Hasher
	expiry int
}

func NewLoginBusiness(storeuser LoginStorage, tokenProvider tokenprovider.Provider, hasher Hasher, expiry int) *loginBusiness {
	return &loginBusiness{
		storeUser:     storeuser,
		tokenProvider: tokenProvider,
		hasher:        hasher,
		expiry:        expiry,
	}
}

// 1. Find user, email
// 2. Hash pass from input and compare with pass in db
// 3. Provider: issue JWT token for client
// 3.1. Access token and refresh token
// 4. Return token(s)

func (business *loginBusiness) Login(ctx context.Context, data *usermodel.UserLogin) (*tokenprovider.Token, error) {
	user, err := business.storeUser.FindUser(ctx, map[string]interface{}{"email":data.Email})

	//log.Println("data.Email: ",data.Email)
	//log.Println("data.Password: ",data.Password)
	//log.Println("data-user: ",user)

	if err != nil {
		return nil, usermodel.ErrUserOrPasswordInvalid
	}

	//log.Println("Step 1: ---------------------")

	passHashed := business.hasher.Hash(data.Password + user.Salt)

	if user.Password != passHashed {
		return nil, usermodel.ErrUserOrPasswordInvalid
	}

	//log.Println("Step 2: ---------------------")

	payload := tokenprovider.TokenPayload{
		UserId: user.Id,
		Role:   user.Role,
	}

	accessToken, err := business.tokenProvider.Generate(payload, business.expiry)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	return accessToken, nil
}