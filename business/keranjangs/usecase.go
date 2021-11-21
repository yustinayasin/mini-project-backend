package keranjangs

import (
	"context"
	"kemejaku/business/kemejakeranjangs"
	"time"
)

type KeranjangUsecase struct {
	repo                KeranjangRepoInterface
	repoKemejaKeranjang kemejakeranjangs.KemejaKeranjangRepoInterface
	ctx                 time.Duration
}

func NewKeranjangUcecase(kRepo KeranjangRepoInterface, kkRepo kemejakeranjangs.KemejaKeranjangRepoInterface, contextTimeout time.Duration) KeranjangUseCaseInterface {
	return &KeranjangUsecase{
		repo:                kRepo,
		repoKemejaKeranjang: kkRepo,
		ctx:                 contextTimeout,
	}
}

//fungsi harus menempel pada struct
func (kUseCase *KeranjangUsecase) InsertKeranjang(kk Keranjang, ctx context.Context) (Keranjang, error) {
	//menghubungkan ke repo
	kkRepo, err := kUseCase.repo.InsertKeranjang(kk, ctx)

	if err != nil {
		return Keranjang{}, err
	}

	return kkRepo, nil
}

func (kUseCase *KeranjangUsecase) GetAllKeranjang(ctx context.Context) ([]Keranjang, error) {
	//menghubungkan ke repo
	kkRepo, err := kUseCase.repo.GetAllKeranjang(ctx)

	if err != nil {
		return []Keranjang{}, err
	}

	return kkRepo, nil
}

func (kUseCase *KeranjangUsecase) GetKeranjangDetail(id int, ctx context.Context) (Keranjang, error) {
	//menghubungkan ke repo
	kkRepo, err := kUseCase.repo.GetKeranjangDetail(id, ctx)

	if err != nil {
		return Keranjang{}, err
	}

	return kkRepo, nil
}

func (kUseCase *KeranjangUsecase) EditKeranjang(kk Keranjang, id int, ctx context.Context) (Keranjang, error) {
	//menghubungkan ke repo
	kkRepo, err := kUseCase.repo.EditKeranjang(kk, id, ctx)

	if err != nil {
		return Keranjang{}, err
	}

	return kkRepo, nil
}

func (kUseCase *KeranjangUsecase) DeleteKeranjang(id int, ctx context.Context) (Keranjang, error) {
	//menghubungkan ke repo
	kkRepo, err := kUseCase.repo.DeleteKeranjang(id, ctx)

	if err != nil {
		return Keranjang{}, err
	}

	return kkRepo, nil
}
