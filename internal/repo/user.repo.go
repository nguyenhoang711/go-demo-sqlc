package repo

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

type userRepository struct{}

func NewUSerRepo() IUserRepository {
	return &userRepository{}
}

// GetUserByGmail implements IUserRepository.
func (ur *userRepository) GetUserByGmail(email string) bool {
	return false
}