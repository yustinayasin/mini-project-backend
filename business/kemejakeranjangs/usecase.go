package kemejakeranjangs

import (
	"context"
	"time"
)

//buat struct supaya bisa pake interface biar bisa dipasangkan dengan yang lainnya
type KemejaKeranjangUsecase struct {
	// interface repo
	repo KemejaKeranjangRepoInterface
	ctx  time.Duration //context untuk time duration
	//misal usecase interaksi dengan yang lainnya nanti bisa ditambahkan
	//usecase lain
	//repo lain dipasangkan di main lewat interface
}

//generate usecase baru
func NewKemejaKeranjangUsecase(kkRepo KemejaKeranjangRepoInterface, contextTimeout time.Duration) KemejaKeranjangUseCaseInterface {
	return &KemejaKeranjangUsecase{
		repo: kkRepo,
		ctx:  contextTimeout,
	}
}

//fungsi harus menempel pada struct
func (kkUseCase *KemejaKeranjangUsecase) InsertKemejaKeranjang(kk KemejaKeranjang, ctx context.Context) (KemejaKeranjang, error) {
	//menghubungkan ke repo
	kkRepo, err := kkUseCase.repo.InsertKemejaKeranjang(kk, ctx)

	if err != nil {
		return KemejaKeranjang{}, err
	}

	return kkRepo, nil
}

func (kkUseCase *KemejaKeranjangUsecase) GetAllKemejaKeranjang(ctx context.Context) ([]KemejaKeranjang, error) {
	//menghubungkan ke repo
	kkRepo, err := kkUseCase.repo.GetAllKemejaKeranjang(ctx)

	if err != nil {
		return []KemejaKeranjang{}, err
	}

	return kkRepo, nil
}

func (kkUseCase *KemejaKeranjangUsecase) GetKemejaKeranjangDetail(id int, ctx context.Context) (KemejaKeranjang, error) {
	//menghubungkan ke repo
	kkRepo, err := kkUseCase.repo.GetKemejaKeranjangDetail(id, ctx)

	if err != nil {
		return KemejaKeranjang{}, err
	}

	return kkRepo, nil
}

func (kkUseCase *KemejaKeranjangUsecase) EditKemejaKeranjang(kk KemejaKeranjang, id int, ctx context.Context) (KemejaKeranjang, error) {
	//menghubungkan ke repo
	kkRepo, err := kkUseCase.repo.EditKemejaKeranjang(kk, id, ctx)

	if err != nil {
		return KemejaKeranjang{}, err
	}

	return kkRepo, nil
}

func (kkUseCase *KemejaKeranjangUsecase) DeleteKemejaKeranjang(id int, ctx context.Context) (KemejaKeranjang, error) {
	//menghubungkan ke repo
	kkRepo, err := kkUseCase.repo.DeleteKemejaKeranjang(id, ctx)

	if err != nil {
		return KemejaKeranjang{}, err
	}

	return kkRepo, nil
}
