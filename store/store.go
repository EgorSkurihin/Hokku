package store

import (
	"errors"

	"github.com/EgorSkurihin/Hokku/models"
)

var (
	ErrAlreadyExist         = errors.New("duplicate entry")
	ErrNoRecord             = errors.New("no matching entry found")
	ErrForeignKeyConstraint = errors.New("foreign key constraint fails")
)

type Store interface {
	Open() error
	Close()

	GetThemes() ([]*models.Theme, error)
	CreateTheme(*models.Theme) (int, error)
	DeleteTheme(int) error

	GetUsers() ([]*models.User, error)
	GetUser(int) (*models.User, error)
	GetUserByEmail(string) (*models.User, error)
	CreateUser(*models.User) (int, error)
	DeleteUser(int) error
	UpdateUser(*models.User) error

	GetHokkus(int, int) ([]*models.Hokku, error)
	GetHokkusByAuthor(int, int, int) ([]*models.Hokku, error)
	GetHokkusByTheme(int, int, int) ([]*models.Hokku, error)
	GetHokku(int) (*models.Hokku, error)
	CreateHokku(*models.Hokku) (int, error)
	DeleteHokku(int) error
	UpdateHokku(*models.Hokku) error
}
