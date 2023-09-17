package userbiz

import (
	"G05-food-delivery/common"
	usermodel "G05-food-delivery/module/user/model"
	"context"
)

type RegisterStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
	CreateUser(ctx context.Context, data *usermodel.UserCreate) error
}

type Hasher interface {
	Hash(data string) string
}

type registerBusiness struct {
	registerStorage RegisterStorage
	hasher Hasher
}

func NewRegisterBusiness(registerStorage RegisterStorage, hasher Hasher) *registerBusiness  {
	return &registerBusiness{
		registerStorage: registerStorage,
		hasher: hasher,
	}
}

func (business *registerBusiness) Register(ctx context.Context, data *usermodel.UserCreate) error  {
	user,_ := business.registerStorage.FindUser(ctx, map[string]interface{}{"email":data.Email})

	if user != nil {
		//if user.Status == 0 {
		//	return errors.New("user has been disable")
		//}
		return usermodel.ErrEmailExisted
	}

	salt := common.GenSalt(50)

	//log.Println("data.Password: ",data.Password)

	//log.Println("salt: ",salt)

	data.Password = business.hasher.Hash(data.Password + salt)

	//log.Println("data.Password md5: ",data.Password)

	data.Salt = salt
	data.Role = "user" // hard code
	//data.Status = 1

	if err := business.registerStorage.CreateUser(ctx,data) ; err != nil {
		return common.ErrCanNotCreateEntity(usermodel.EntityName, err)
	}

	return nil

}