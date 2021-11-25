package kemejakeranjangs_test

import (
	"context"
	"errors"
	"kemejaku/business/kemejakeranjangs"
	"kemejaku/business/kemejakeranjangs/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

//buat mock yang seolah-olah interface dari database
var kemejaKeranjangRepoInterfaceMock mocks.KemejaKeranjangRepoInterface
var kemejaKeranjangUseCaseInterface kemejakeranjangs.KemejaKeranjangUseCaseInterface
var kemejaKeranjangDataDummy, kemejaKeranjangDataDummyEdit kemejakeranjangs.KemejaKeranjang
var kkDummyGetAll []kemejakeranjangs.KemejaKeranjang

func setup() {
	kemejaKeranjangUseCaseInterface = kemejakeranjangs.NewKemejaKeranjangUsecase(&kemejaKeranjangRepoInterfaceMock, time.Hour*1)

	//data mock hasil login
	kemejaKeranjangDataDummy = kemejakeranjangs.KemejaKeranjang{
		Id:          1,
		IdKemeja:    1,
		IdKeranjang: 3,
		Jumlah:      2,
		Size:        "M",
	}

	kkDummyGetAll = []kemejakeranjangs.KemejaKeranjang{
		{
			Id:          1,
			IdKemeja:    1,
			IdKeranjang: 3,
			Jumlah:      2,
			Size:        "M",
		},
		{
			Id:          2,
			IdKemeja:    1,
			IdKeranjang: 3,
			Jumlah:      2,
			Size:        "M",
		},
	}
}

func TestInsertKemejaKeranjang(t *testing.T) {
	setup()
	t.Run("Success insert", func(t *testing.T) {
		kemejaKeranjangRepoInterfaceMock.On("InsertKemejaKeranjang", mock.AnythingOfType("kemejakeranjangs.KemejaKeranjang"), mock.Anything).Return(kemejaKeranjangDataDummy, nil).Once()

		var requestInsertkemejaKeranjang = kemejakeranjangs.KemejaKeranjang{
			IdKemeja:    1,
			IdKeranjang: 3,
			Jumlah:      1,
			Size:        "S",
		}

		kemejaKeranjang, err := kemejaKeranjangUseCaseInterface.InsertKemejaKeranjang(requestInsertkemejaKeranjang, context.Background())

		assert.Equal(t, nil, err)
		assert.Equal(t, kemejaKeranjangDataDummy, kemejaKeranjang)
	})

	t.Run("Insert failed", func(t *testing.T) {
		kemejaKeranjangRepoInterfaceMock.On("InsertKemejaKeranjang", mock.AnythingOfType("kemejakeranjangs.KemejaKeranjang"), mock.Anything).Return(kemejakeranjangs.KemejaKeranjang{}, errors.New("Insert failed")).Once()

		var requestInsertkemejaKeranjang = kemejakeranjangs.KemejaKeranjang{
			IdKemeja:    1,
			IdKeranjang: 3,
			Jumlah:      1,
			Size:        "S",
		}

		kemejaKeranjang, err := kemejaKeranjangUseCaseInterface.InsertKemejaKeranjang(requestInsertkemejaKeranjang, context.Background())

		assert.Equal(t, errors.New("Insert failed"), err)
		assert.Equal(t, kemejakeranjangs.KemejaKeranjang{}, kemejaKeranjang)
	})
}

func TestGetAllkemejaKeranjang(t *testing.T) {
	setup()
	t.Run("Success Get All Kemeja Keranjangs", func(t *testing.T) {
		kemejaKeranjangRepoInterfaceMock.On("GetAllKemejaKeranjang", mock.Anything, mock.Anything).Return(kkDummyGetAll, nil).Once()

		kemejaKeranjang, err := kemejaKeranjangUseCaseInterface.GetAllKemejaKeranjang(context.Background())

		assert.Equal(t, nil, err)
		assert.Equal(t, kkDummyGetAll, kemejaKeranjang)
	})

	t.Run("KemejaKeranjang not found in database", func(t *testing.T) {
		kemejaKeranjangRepoInterfaceMock.On("GetAllKemejaKeranjang", mock.Anything, mock.Anything).Return([]kemejakeranjangs.KemejaKeranjang{}, errors.New("There is no kemejaKeranjang column"))

		kemejaKeranjang, err := kemejaKeranjangUseCaseInterface.GetAllKemejaKeranjang(context.Background())

		assert.Equal(t, errors.New("There is no kemejaKeranjang column"), err)
		assert.Equal(t, []kemejakeranjangs.KemejaKeranjang{}, kemejaKeranjang)
	})
}

func TestGetKemejaKeranjangDetail(t *testing.T) {
	setup()
	t.Run("Success Get Kemeja Keranjang Detail", func(t *testing.T) {
		kemejaKeranjangRepoInterfaceMock.On("GetKemejaKeranjangDetail", mock.Anything, mock.Anything).Return(kemejaKeranjangDataDummy, nil).Once()

		kemejaKeranjang, err := kemejaKeranjangUseCaseInterface.GetKemejaKeranjangDetail(1, context.Background())

		assert.Equal(t, nil, err)
		assert.Equal(t, kemejaKeranjangDataDummy, kemejaKeranjang)
	})

	t.Run("KemejaKeranjangs not found in database", func(t *testing.T) {
		kemejaKeranjangRepoInterfaceMock.On("GetKemejaKeranjangDetail", mock.Anything, mock.Anything).Return(kemejakeranjangs.KemejaKeranjang{}, errors.New("KemejaKeranjang not found")).Once()

		kemejaKeranjang, err := kemejaKeranjangUseCaseInterface.GetKemejaKeranjangDetail(-1, context.Background())

		assert.Equal(t, errors.New("KemejaKeranjang not found"), err)
		assert.Equal(t, kemejakeranjangs.KemejaKeranjang{}, kemejaKeranjang)
	})
}

func TestEditKemejaKeranjang(t *testing.T) {
	setup()
	t.Run("Success Edit", func(t *testing.T) {
		kemejaKeranjangRepoInterfaceMock.On("EditKemejaKeranjang", mock.AnythingOfType("kemejakeranjangs.KemejaKeranjang"), mock.Anything, mock.Anything).Return(kemejaKeranjangDataDummyEdit, nil).Once()

		var requestEditKemejaKeranjang = kemejakeranjangs.KemejaKeranjang{
			Jumlah: 1,
			Size:   "L",
		}

		kemejaKeranjang, err := kemejaKeranjangUseCaseInterface.EditKemejaKeranjang(requestEditKemejaKeranjang, 1, context.Background())

		assert.Equal(t, nil, err)
		assert.Equal(t, kemejaKeranjangDataDummyEdit, kemejaKeranjang)
	})

	t.Run("Kemeja Keranjang not found", func(t *testing.T) {
		kemejaKeranjangRepoInterfaceMock.On("EditKemejaKeranjang", mock.AnythingOfType("kemejakeranjangs.KemejaKeranjang"), mock.Anything, mock.Anything).Return(kemejakeranjangs.KemejaKeranjang{}, errors.New("Kemeja Keranjang not found")).Once()

		var requestEditKemejaKeranjang = kemejakeranjangs.KemejaKeranjang{
			Jumlah: 1,
			Size:   "L",
		}
		kemejaKeranjang, err := kemejaKeranjangUseCaseInterface.EditKemejaKeranjang(requestEditKemejaKeranjang, 1, context.Background())

		assert.Equal(t, errors.New("Kemeja Keranjang not found"), err)
		assert.Equal(t, kemejakeranjangs.KemejaKeranjang{}, kemejaKeranjang)
	})
}

func TestDeleteKemejaKeranjang(t *testing.T) {
	setup()
	t.Run("Success delete", func(t *testing.T) {
		kemejaKeranjangRepoInterfaceMock.On("DeleteKemejaKeranjang", mock.Anything, mock.Anything).Return(kemejaKeranjangDataDummy, nil).Once()

		kemejaKeranjang, err := kemejaKeranjangUseCaseInterface.DeleteKemejaKeranjang(1, context.Background())

		assert.Equal(t, nil, err)
		assert.Equal(t, kemejaKeranjangDataDummy, kemejaKeranjang)
	})

	t.Run("kemejaKeranjangs not found", func(t *testing.T) {
		kemejaKeranjangRepoInterfaceMock.On("DeleteKemejaKeranjang", mock.Anything, mock.Anything).Return(kemejakeranjangs.KemejaKeranjang{}, errors.New("kemejaKeranjang not found")).Once()

		kemejaKeranjang, err := kemejaKeranjangUseCaseInterface.DeleteKemejaKeranjang(-1, context.Background())

		assert.Equal(t, errors.New("kemejaKeranjang not found"), err)
		assert.Equal(t, kemejakeranjangs.KemejaKeranjang{}, kemejaKeranjang)
	})
}
