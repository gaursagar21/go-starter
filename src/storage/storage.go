package storage

import (
	"context"
	placeStorage "github.com/gaursagarMT/starter/src/storage/place"
	"github.com/gaursagarMT/starter/src/storage/user"
)

type IStorage interface {
	CreateTables(ctx context.Context) error
	userStorage.IUserStorage
	placeStorage.IPlaceStorage
}
