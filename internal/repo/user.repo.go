package repo

import (
	"github.com/hoangnguyen/demo-sqlc/global"
	model "github.com/hoangnguyen/demo-sqlc/internal/models"
)

// type UserRepo struct {
// }

// func NewUserRepo() *UserRepo {
// 	return &UserRepo{}
// }

// func(ur *UserRepo) GetUserDetail(id string) string {
// 	return fmt.Sprintf("Tips %s", id)
// }

// INTERFACE_VERSION
type IUserRepository interface {
	GetUserByGmail(email string) bool
}

type userRepository struct{
	sqlc *model.Queries
}

func NewUSerRepo() IUserRepository {
	return &userRepository{
		sqlc: model.New(global.Mdbc),
	}
}

// GetUserByGmail implements IUserRepository.
func (up *userRepository) GetUserByGmail(email string) bool {
	// SELECT * FROM user WHERE email = ? LIMIT 1
	// row := global.Mdb.Table(TableNameGoCrmUser).Where("usr_email = ?", email).First(&model.PreGoCrmUserC)
	user, err := up.sqlc.GetUserByEmailSQLC(ctx, email)
	if err != nil {
		return false
	}
	return user.UsrID != 0
}