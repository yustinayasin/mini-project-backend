package keranjangs_test

import (
	"context"
	"errors"
	"kemejaku/business/keranjangs"
	"kemejaku/business/keranjangs/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

//buat mock yang seolah-olah interface dari database
var keranjangRepoInterfaceMock mocks.KeranjangRepoInterface
var keranjangUseCaseInterface keranjangs.KeranjangUseCaseInterface
var keranjangDataDummy, keranjangDataDummyEdit keranjangs.Keranjang
var keranjangDataDummyGetAllkeranjangs []keranjangs.Keranjang

func setup() {
	keranjangUseCaseInterface = keranjangs.NewKeranjangUcecase(&keranjangRepoInterfaceMock, time.Hour*1)

	//data mock hasil login
	keranjangDataDummy = keranjangs.Keranjang{
		Id:     1,
		IdUser: 1,
		Status: false,
	}

	keranjangDataDummyGetAllkeranjangs = []keranjangs.Keranjang{
		{
			Id:     1,
			IdUser: 1,
			Status: false,
		},
		{
			Id:     1,
			IdUser: 2,
			Status: false,
		},
	}
}

func TestInsertKeranjang(t *testing.T) {
	setup()
	t.Run("Success insert", func(t *testing.T) {
		keranjangRepoInterfaceMock.On("InsertKeranjang", mock.AnythingOfType("keranjangs.Keranjang"), mock.Anything).Return(keranjangDataDummy, nil).Once()

		var requestInsertKeranjang = keranjangs.Keranjang{
			IdUser: 1,
		}

		keranjang, err := keranjangUseCaseInterface.InsertKeranjang(requestInsertKeranjang, context.Background())

		assert.Equal(t, nil, err)
		assert.Equal(t, keranjangDataDummy, keranjang)
	})

	t.Run("Insert failed", func(t *testing.T) {
		keranjangRepoInterfaceMock.On("InsertKeranjang", mock.AnythingOfType("keranjangs.Keranjang"), mock.Anything).Return(keranjangs.Keranjang{}, errors.New("Insert failed")).Once()

		var requestInsertKeranjang = keranjangs.Keranjang{
			IdUser: 1,
		}
		keranjang, err := keranjangUseCaseInterface.InsertKeranjang(requestInsertKeranjang, context.Background())

		assert.Equal(t, errors.New("Insert failed"), err)
		assert.Equal(t, keranjangs.Keranjang{}, keranjang)
	})
}

func TestGetAllKeranjang(t *testing.T) {
	setup()
	t.Run("Success Get All Keranjangs", func(t *testing.T) {
		keranjangRepoInterfaceMock.On("GetAllKeranjang", mock.Anything, mock.Anything).Return(keranjangDataDummyGetAllkeranjangs, nil).Once()

		keranjang, err := keranjangUseCaseInterface.GetAllKeranjang(context.Background())

		assert.Equal(t, nil, err)
		assert.Equal(t, keranjangDataDummyGetAllkeranjangs, keranjang)
	})

	t.Run("Keranjang not found in database", func(t *testing.T) {
		keranjangRepoInterfaceMock.On("GetAllKeranjang", mock.Anything, mock.Anything).Return([]keranjangs.Keranjang{}, errors.New("There is no keranjang column"))

		keranjang, err := keranjangUseCaseInterface.GetAllKeranjang(context.Background())

		assert.Equal(t, errors.New("There is no keranjang column"), err)
		assert.Equal(t, []keranjangs.Keranjang{}, keranjang)
	})
}

func TestGetKeranjangDetail(t *testing.T) {
	setup()
	t.Run("Success Get keranjang Detail", func(t *testing.T) {
		keranjangRepoInterfaceMock.On("GetKeranjangDetail", mock.Anything, mock.Anything).Return(keranjangDataDummy, nil).Once()

		keranjang, err := keranjangUseCaseInterface.GetKeranjangDetail(1, context.Background())

		assert.Equal(t, nil, err)
		assert.Equal(t, keranjangDataDummy, keranjang)
	})

	t.Run("Keranjangs not found in database", func(t *testing.T) {
		keranjangRepoInterfaceMock.On("GetKeranjangDetail", mock.Anything, mock.Anything).Return(keranjangs.Keranjang{}, errors.New("Keranjang not found")).Once()

		keranjang, err := keranjangUseCaseInterface.GetKeranjangDetail(-1, context.Background())

		assert.Equal(t, errors.New("Keranjang not found"), err)
		assert.Equal(t, keranjangs.Keranjang{}, keranjang)
	})
}

func TestEditKeranjang(t *testing.T) {
	setup()
	t.Run("Success Edit", func(t *testing.T) {
		keranjangRepoInterfaceMock.On("EditKeranjang", mock.AnythingOfType("keranjangs.Keranjang"), mock.Anything, mock.Anything).Return(keranjangDataDummyEdit, nil).Once()

		var requestEditKeranjang = keranjangs.Keranjang{
			Status: true,
		}

		keranjang, err := keranjangUseCaseInterface.EditKeranjang(requestEditKeranjang, 1, context.Background())

		assert.Equal(t, nil, err)
		assert.Equal(t, keranjangDataDummyEdit, keranjang)
	})

	t.Run("keranjang not found", func(t *testing.T) {
		keranjangRepoInterfaceMock.On("EditKeranjang", mock.AnythingOfType("keranjangs.Keranjang"), mock.Anything, mock.Anything).Return(keranjangs.Keranjang{}, errors.New("keranjang not found")).Once()

		var requestEditKeranjang = keranjangs.Keranjang{
			Status: true,
		}
		keranjang, err := keranjangUseCaseInterface.EditKeranjang(requestEditKeranjang, 1, context.Background())

		assert.Equal(t, errors.New("keranjang not found"), err)
		assert.Equal(t, keranjangs.Keranjang{}, keranjang)
	})
}

func TestDeleteKeranjang(t *testing.T) {
	setup()
	t.Run("Success delete", func(t *testing.T) {
		keranjangRepoInterfaceMock.On("DeleteKeranjang", mock.Anything, mock.Anything).Return(keranjangDataDummy, nil).Once()

		keranjang, err := keranjangUseCaseInterface.DeleteKeranjang(1, context.Background())

		assert.Equal(t, nil, err)
		assert.Equal(t, keranjangDataDummy, keranjang)
	})

	t.Run("keranjangs not found", func(t *testing.T) {
		keranjangRepoInterfaceMock.On("DeleteKeranjang", mock.Anything, mock.Anything).Return(keranjangs.Keranjang{}, errors.New("keranjang not found")).Once()

		keranjang, err := keranjangUseCaseInterface.DeleteKeranjang(-1, context.Background())

		assert.Equal(t, errors.New("keranjang not found"), err)
		assert.Equal(t, keranjangs.Keranjang{}, keranjang)
	})
}
