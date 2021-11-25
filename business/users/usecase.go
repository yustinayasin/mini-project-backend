package users

import (
	"context"
	_middleware "kemejaku/app/middleware"
	"time"
)

//buat struct supaya bisa pake interface biar bisa dipasangkan dengan yang lainnya
type UserUseCase struct {
	// interface repo
	repo UserRepoInterface
	ctx  time.Duration //context untuk time duration
	jwt  *_middleware.ConfigJWT
	//misal usecase interaksi dengan yang lainnya nanti bisa ditambahkan
	//usecase lain
	//repo lain dipasangkan di main lewat interface
}

//generate usecase baru configJwt *_middleware.ConfigJWT
func NewUseCase(userRepo UserRepoInterface, contextTimeout time.Duration) UserUseCaseInterface {
	return &UserUseCase{
		repo: userRepo,
		ctx:  contextTimeout,
		// jwt:  configJwt,
	}
}

//fungsi harus menempel pada struct
func (userUseCase *UserUseCase) Login(user User, ctx context.Context) (User, error) {
	//menghubungkan ke repo
	userRepo, err := userUseCase.repo.Login(user, ctx)

	if err != nil {
		return User{}, err
	}

	// userRepo.Token = userUseCase.jwt.GenerateToken(user.Id)

	return userRepo, nil
}

func (userUseCase *UserUseCase) GetAllUsers(ctx context.Context) ([]User, error) {
	userRepo, err := userUseCase.repo.GetAllUsers(ctx)

	if err != nil {
		return []User{}, err
	}

	return userRepo, nil
}

func (userUseCase *UserUseCase) SignUp(user User, ctx context.Context) (User, error) {
	userRepo, err := userUseCase.repo.SignUp(user, ctx)

	if err != nil {
		return User{}, err
	}

	return userRepo, nil
}

func (userUseCase *UserUseCase) GetUserDetail(id int, ctx context.Context) (User, error) {
	userRepo, err := userUseCase.repo.GetUserDetail(id, ctx)

	if err != nil {
		return User{}, err
	}

	return userRepo, nil
}

func (userUseCase *UserUseCase) EditUser(user User, id int, ctx context.Context) (User, error) {
	userRepo, err := userUseCase.repo.EditUser(user, id, ctx)

	if err != nil {
		return User{}, err
	}

	return userRepo, nil
}

func (userUseCase *UserUseCase) DeleteUser(id int, ctx context.Context) (User, error) {
	userRepo, err := userUseCase.repo.DeleteUser(id, ctx)

	if err != nil {
		return User{}, err
	}

	return userRepo, nil
}
