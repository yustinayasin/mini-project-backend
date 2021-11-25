package sales

import (
	"context"
	"time"
)

//buat struct supaya bisa pake interface biar bisa dipasangkan dengan yang lainnya
type SaleUsecase struct {
	// interface repo
	repo SaleRepoInterface
	ctx  time.Duration
}

//generate usecase baru
func NewSaleUsecase(saleRepo SaleRepoInterface, contextTimeout time.Duration) SaleUseCaseInterface {
	return &SaleUsecase{
		repo: saleRepo,
		ctx:  contextTimeout,
	}
}

func (saleUseCase *SaleUsecase) InsertSale(kk Sale, ctx context.Context) (Sale, error) {
	//menghubungkan ke repo
	saleRepo, err := saleUseCase.repo.InsertSale(kk, ctx)

	if err != nil {
		return Sale{}, err
	}

	return saleRepo, nil
}

func (saleUseCase *SaleUsecase) GetAllSale(ctx context.Context) ([]Sale, error) {
	//menghubungkan ke repo
	saleRepo, err := saleUseCase.repo.GetAllSale(ctx)

	if err != nil {
		return []Sale{}, err
	}

	return saleRepo, nil
}

func (saleUseCase *SaleUsecase) GetSaleDetail(id int, ctx context.Context) (Sale, error) {
	//menghubungkan ke repo
	saleRepo, err := saleUseCase.repo.GetSaleDetail(id, ctx)

	if err != nil {
		return Sale{}, err
	}

	return saleRepo, nil
}

func (saleUseCase *SaleUsecase) EditSale(kk Sale, id int, ctx context.Context) (Sale, error) {
	//menghubungkan ke repo
	saleRepo, err := saleUseCase.repo.EditSale(kk, id, ctx)

	if err != nil {
		return Sale{}, err
	}

	return saleRepo, nil
}

func (saleUseCase *SaleUsecase) DeleteSale(id int, ctx context.Context) (Sale, error) {
	//menghubungkan ke repo
	saleRepo, err := saleUseCase.repo.DeleteSale(id, ctx)

	if err != nil {
		return Sale{}, err
	}

	return saleRepo, nil
}
