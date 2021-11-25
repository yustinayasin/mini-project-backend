package kemejas_test

import (
	"context"
	"errors"
	"kemejaku/business/kemejas"
	"kemejaku/business/kemejas/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

//buat mock yang seolah-olah interface dari database
var kemejaRepoInterfaceMock mocks.KemejaRepoInterface
var kemejaUseCaseInterface kemejas.KemejaUseCaseInterface
var kemejaDataDummy, kemejaDataDummyEdit kemejas.Kemeja
var kemejaDataDummyGetAllKemejas []kemejas.Kemeja

func setup() {
	kemejaUseCaseInterface = kemejas.NewKemejaUsecase(&kemejaRepoInterfaceMock, time.Hour*1)

	//data mock hasil login
	kemejaDataDummy = kemejas.Kemeja{
		Id:        1,
		Nama:      "Basic shirt",
		Deskripsi: "Bahan rayon",
		Stock_L:   10,
		Stock_M:   10,
		Stock_S:   10,
	}

	kemejaDataDummyGetAllKemejas = []kemejas.Kemeja{
		{
			Id:        1,
			Nama:      "Basic shirt",
			Deskripsi: "Bahan rayon",
			Stock_L:   10,
			Stock_M:   10,
			Stock_S:   10,
		},
		{
			Id:        2,
			Nama:      "Basic shirt",
			Deskripsi: "Bahan rayon",
			Stock_L:   10,
			Stock_M:   10,
			Stock_S:   10,
		},
	}
}

func TestInsertKemeja(t *testing.T) {
	setup()
	t.Run("Success insert", func(t *testing.T) {
		kemejaRepoInterfaceMock.On("InsertKemeja", mock.AnythingOfType("kemejas.Kemeja"), mock.Anything).Return(kemejaDataDummy, nil).Once()

		var requestInsertKemeja = kemejas.Kemeja{
			Nama:      "Basic shirt",
			Deskripsi: "Bahan rayon",
			Stock_L:   10,
			Stock_M:   10,
			Stock_S:   10,
		}

		kemeja, err := kemejaUseCaseInterface.InsertKemeja(requestInsertKemeja, context.Background())

		assert.Equal(t, nil, err)
		assert.Equal(t, kemejaDataDummy, kemeja)
	})

	t.Run("Insert failed", func(t *testing.T) {
		kemejaRepoInterfaceMock.On("InsertKemeja", mock.AnythingOfType("kemejas.Kemeja"), mock.Anything).Return(kemejas.Kemeja{}, errors.New("Insert failed")).Once()

		var requestInsertKemeja = kemejas.Kemeja{
			Nama:      "Basic shirt",
			Deskripsi: "Bahan rayon",
			Stock_L:   10,
			Stock_M:   10,
			Stock_S:   10,
		}
		kemeja, err := kemejaUseCaseInterface.InsertKemeja(requestInsertKemeja, context.Background())

		assert.Equal(t, errors.New("Insert failed"), err)
		assert.Equal(t, kemejas.Kemeja{}, kemeja)
	})
}

func TestGetAllkemeja(t *testing.T) {
	setup()
	t.Run("Success Get All kemejas", func(t *testing.T) {
		kemejaRepoInterfaceMock.On("GetAllKemeja", mock.Anything, mock.Anything).Return(kemejaDataDummyGetAllKemejas, nil).Once()

		kemeja, err := kemejaUseCaseInterface.GetAllKemeja(context.Background())

		assert.Equal(t, nil, err)
		assert.Equal(t, kemejaDataDummyGetAllKemejas, kemeja)
	})

	t.Run("Kemeja not found in database", func(t *testing.T) {
		kemejaRepoInterfaceMock.On("GetAllKemeja", mock.Anything, mock.Anything).Return([]kemejas.Kemeja{}, errors.New("There is no kemeja column"))

		kemeja, err := kemejaUseCaseInterface.GetAllKemeja(context.Background())

		assert.Equal(t, errors.New("There is no kemeja column"), err)
		assert.Equal(t, []kemejas.Kemeja{}, kemeja)
	})
}

func TestGetkemejaDetail(t *testing.T) {
	setup()
	t.Run("Success Get kemeja Detail", func(t *testing.T) {
		kemejaRepoInterfaceMock.On("GetKemejaDetail", mock.Anything, mock.Anything).Return(kemejaDataDummy, nil).Once()

		kemeja, err := kemejaUseCaseInterface.GetKemejaDetail(1, context.Background())

		assert.Equal(t, nil, err)
		assert.Equal(t, kemejaDataDummy, kemeja)
	})

	t.Run("Kemejas not found in database", func(t *testing.T) {
		kemejaRepoInterfaceMock.On("GetKemejaDetail", mock.Anything, mock.Anything).Return(kemejas.Kemeja{}, errors.New("Kemeja not found")).Once()

		kemeja, err := kemejaUseCaseInterface.GetKemejaDetail(-1, context.Background())

		assert.Equal(t, errors.New("Kemeja not found"), err)
		assert.Equal(t, kemejas.Kemeja{}, kemeja)
	})
}

func TestEditkemeja(t *testing.T) {
	setup()
	t.Run("Success Edit", func(t *testing.T) {
		kemejaRepoInterfaceMock.On("EditKemeja", mock.AnythingOfType("kemejas.Kemeja"), mock.Anything, mock.Anything).Return(kemejaDataDummyEdit, nil).Once()

		var requestEditKemeja = kemejas.Kemeja{
			Nama:      "Basic shirt",
			Deskripsi: "Bahan rayon",
			Stock_L:   10,
			Stock_M:   10,
			Stock_S:   10,
		}

		kemeja, err := kemejaUseCaseInterface.EditKemeja(requestEditKemeja, 1, context.Background())

		assert.Equal(t, nil, err)
		assert.Equal(t, kemejaDataDummyEdit, kemeja)
	})

	t.Run("Kemeja not found", func(t *testing.T) {
		kemejaRepoInterfaceMock.On("EditKemeja", mock.AnythingOfType("kemejas.Kemeja"), mock.Anything, mock.Anything).Return(kemejas.Kemeja{}, errors.New("kemeja not found")).Once()

		var requestEditKemeja = kemejas.Kemeja{
			Nama:      "Basic shirt",
			Deskripsi: "Bahan rayon",
			Stock_L:   10,
			Stock_M:   10,
			Stock_S:   10,
		}
		kemeja, err := kemejaUseCaseInterface.EditKemeja(requestEditKemeja, 1, context.Background())

		assert.Equal(t, errors.New("kemeja not found"), err)
		assert.Equal(t, kemejas.Kemeja{}, kemeja)
	})
}

func TestDeletekemeja(t *testing.T) {
	setup()
	t.Run("Success delete", func(t *testing.T) {
		kemejaRepoInterfaceMock.On("DeleteKemeja", mock.Anything, mock.Anything).Return(kemejaDataDummy, nil).Once()

		kemeja, err := kemejaUseCaseInterface.DeleteKemeja(1, context.Background())

		assert.Equal(t, nil, err)
		assert.Equal(t, kemejaDataDummy, kemeja)
	})

	t.Run("kemejas not found", func(t *testing.T) {
		kemejaRepoInterfaceMock.On("DeleteKemeja", mock.Anything, mock.Anything).Return(kemejas.Kemeja{}, errors.New("kemeja not found")).Once()

		kemeja, err := kemejaUseCaseInterface.DeleteKemeja(-1, context.Background())

		assert.Equal(t, errors.New("kemeja not found"), err)
		assert.Equal(t, kemejas.Kemeja{}, kemeja)
	})
}
