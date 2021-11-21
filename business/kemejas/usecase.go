package kemejas

import (
	"context"
	"time"
)

//buat struct supaya bisa pake interface biar bisa dipasangkan dengan yang lainnya
type KemejaUsecase struct {
	// interface repo
	repo KemejaRepoInterface
	ctx  time.Duration
}

//generate usecase baru
func NewKemejaUsecase(kemejaRepo KemejaRepoInterface, contextTimeout time.Duration) KemejaUseCaseInterface {
	return &KemejaUsecase{
		repo: kemejaRepo,
		ctx:  contextTimeout,
	}
}

func (kemejaUseCase *KemejaUsecase) InsertKemeja(kk Kemeja, ctx context.Context) (Kemeja, error) {
	//menghubungkan ke repo
	kemejaRepo, err := kemejaUseCase.repo.InsertKemeja(kk, ctx)

	if err != nil {
		return Kemeja{}, err
	}

	return kemejaRepo, nil
}

func (kemejaUseCase *KemejaUsecase) GetAllKemeja(ctx context.Context) ([]Kemeja, error) {
	//menghubungkan ke repo
	kemejaRepo, err := kemejaUseCase.repo.GetAllKemeja(ctx)

	if err != nil {
		return []Kemeja{}, err
	}

	return kemejaRepo, nil
}

func (kemejaUseCase *KemejaUsecase) GetKemejaDetail(id int, ctx context.Context) (Kemeja, error) {
	//menghubungkan ke repo
	kemejaRepo, err := kemejaUseCase.repo.GetKemejaDetail(id, ctx)

	if err != nil {
		return Kemeja{}, err
	}

	return kemejaRepo, nil
}

func (kemejaUseCase *KemejaUsecase) EditKemeja(kk Kemeja, id int, ctx context.Context) (Kemeja, error) {
	//menghubungkan ke repo
	kemejaRepo, err := kemejaUseCase.repo.EditKemeja(kk, id, ctx)

	if err != nil {
		return Kemeja{}, err
	}

	return kemejaRepo, nil
}

func (kemejaUseCase *KemejaUsecase) DeleteKemeja(id int, ctx context.Context) (Kemeja, error) {
	//menghubungkan ke repo
	kemejaRepo, err := kemejaUseCase.repo.DeleteKemeja(id, ctx)

	if err != nil {
		return Kemeja{}, err
	}

	return kemejaRepo, nil
}
