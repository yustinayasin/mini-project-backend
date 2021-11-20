package users

import (
	"context"
	"time"
)

//buat struct supaya bisa pake interface biar bisa dipasangkan dengan yang lainnya
type UserUseCase struct {
	// interface repo
	repo UserRepoInterface
	ctx  time.Duration //context untuk time duration
	//misal usecase interaksi dengan yang lainnya nanti bisa ditambahkan
	//usecase lain
	//repo lain dipasangkan di main lewat interface
}

//generate usecase baru
func NewUseCase(userRepo UserRepoInterface, contextTimeout time.Duration) UserUseCaseInterface {
	return &UserUseCase{
		repo: userRepo,
		ctx:  contextTimeout,
	}
}

//fungsi harus menempel pada struct
func (userUseCase *UserUseCase) LoginController(user User, ctx context.Context) (User, error) {
	// if user.Email == "" {
	// 	return User{}, errors.New("Email Empty")
	// }

	// if user.Password == "" {
	// 	return User{}, errors.New("Password Empty")
	// }

	//menghubungkan ke repo
	userRepo, err := userUseCase.repo.LoginController(user, ctx)

	if err != nil {
		return User{}, err
	}

	return userRepo, nil
}

func (userUseCase *UserUseCase) GetAllUsersController(ctx context.Context) ([]User, error) {
	userRepo, err := userUseCase.repo.GetAllUsersController(ctx)

	if err != nil {
		return []User{}, err
	}

	return userRepo, nil
}

func (userUseCase *UserUseCase) SignUpController(user User, ctx context.Context) (User, error) {
	userRepo, err := userUseCase.repo.SignUpController(user, ctx)

	if err != nil {
		return User{}, err
	}

	return userRepo, nil
}

func (userUseCase *UserUseCase) GetDetailUserController(id int, ctx context.Context) (User, error) {
	userRepo, err := userUseCase.repo.GetDetailUserController(id, ctx)

	if err != nil {
		return User{}, err
	}

	return userRepo, nil
}

func (userUseCase *UserUseCase) EditUserController(user User, id int, ctx context.Context) (User, error) {
	userRepo, err := userUseCase.repo.EditUserController(user, id, ctx)

	if err != nil {
		return User{}, err
	}

	return userRepo, nil
}

func (userUseCase *UserUseCase) DeleteUserController(id int, ctx context.Context) (User, error) {
	userRepo, err := userUseCase.repo.DeleteUserController(id, ctx)

	if err != nil {
		return User{}, err
	}

	return userRepo, nil
}
