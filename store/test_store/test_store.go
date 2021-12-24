package test_store

import (
	"github.com/EgorSkurihin/Hokku/models"
	"github.com/EgorSkurihin/Hokku/store"
)

type TestStore struct {
	Users  []*models.User
	Hokkus []*models.Hokku
	Themes []*models.Theme
}

func New() *TestStore {
	return &TestStore{
		Users:  Users,
		Hokkus: Hokkus,
		Themes: Themes,
	}
}

func (s *TestStore) Open() error {
	return nil
}

func (s *TestStore) Close() {}

func (s *TestStore) GetThemes() ([]*models.Theme, error) {
	return s.Themes, nil
}

func (s *TestStore) CreateTheme(theme *models.Theme) (int, error) {
	for _, t := range s.Themes {
		if t.Title == theme.Title {
			return 0, store.ErrAlreadyExist
		}
	}
	theme.Id = len(s.Users) + 1
	s.Themes = append(s.Themes, theme)
	return theme.Id, nil
}

func (s *TestStore) DeleteTheme(id int) error {
	if id < 0 || id > len(s.Themes)-1 {
		return store.ErrNoRecord
	}
	s.Themes[id] = s.Themes[len(s.Themes)-1]
	s.Themes = s.Themes[:len(s.Themes)-1]
	return nil
}

func (s *TestStore) GetUsers() ([]*models.User, error) {
	return s.Users, nil
}

func (s *TestStore) GetUser(id int) (*models.User, error) {
	if id < 0 || id > len(s.Users) {
		return nil, store.ErrNoRecord
	}
	return s.Users[id-1], nil
}

func (s *TestStore) GetUserByEmail(email string) (*models.User, error) {
	for _, u := range s.Users {
		if u.Email == email {
			return u, nil
		}
	}
	return nil, store.ErrNoRecord
}

func (s *TestStore) CreateUser(user *models.User) (int, error) {
	for _, u := range s.Users {
		if u.Email == user.Email {
			return 0, store.ErrAlreadyExist
		}
	}
	user.Id = len(s.Users) + 1
	s.Users = append(s.Users, user)
	return user.Id, nil
}

func (s *TestStore) DeleteUser(id int) error {
	if id < 0 || id > len(s.Users)-1 {
		return store.ErrNoRecord
	}
	s.Users[id] = s.Users[len(s.Users)-1]
	s.Users = s.Users[:len(s.Users)-1]
	return nil
}

func (s *TestStore) UpdateUser(user *models.User) error {
	if user.Id < 0 || user.Id > len(s.Users) {
		return store.ErrNoRecord
	}
	s.Users[user.Id-1] = user
	return nil
}

func (s *TestStore) GetHokkus(limit, offset int) ([]*models.Hokku, error) {
	var res []*models.Hokku
	if limit != 0 {
		if offset >= len(s.Hokkus) {
			res = make([]*models.Hokku, 0)
		}
		if limit+offset >= len(s.Hokkus) {
			res = s.Hokkus[offset:]
		}
		if limit+offset < len(s.Hokkus) {
			res = s.Hokkus[offset : offset+limit]
		}
	} else {
		res = s.Hokkus
	}
	return res, nil
}

func (s *TestStore) GetHokkusByAuthor(authorId, limit, offset int) ([]*models.Hokku, error) {
	res := make([]*models.Hokku, 0)
	for _, h := range s.Hokkus {
		if h.OwnerId == authorId {
			res = append(res, h)
		}
	}
	if limit != 0 {
		if limit+offset >= len(res) {
			if offset >= len(res) {
				res = make([]*models.Hokku, 0)
			} else {
				res = res[offset:]
			}
		}
		if limit+offset < len(res) {
			res = res[offset : offset+limit]
		}
	}
	return res, nil
}

func (s *TestStore) GetHokkusByTheme(themeId, limit, offset int) ([]*models.Hokku, error) {
	res := make([]*models.Hokku, 0)
	for _, h := range s.Hokkus {
		if h.ThemeId == themeId {
			res = append(res, h)
		}
	}
	if limit != 0 {
		if limit+offset >= len(res) {
			if offset >= len(res) {
				res = make([]*models.Hokku, 0)
			} else {
				res = res[offset:]
			}
		}
		if limit+offset < len(res) {
			res = res[offset : offset+limit]
		}
	}
	return res, nil
}

func (s *TestStore) GetHokku(id int) (*models.Hokku, error) {
	if id < 0 || id > len(s.Hokkus) {
		return nil, store.ErrNoRecord
	}
	return s.Hokkus[id-1], nil
}

func (s *TestStore) CreateHokku(hokku *models.Hokku) (int, error) {
	if hokku.ThemeId > len(s.Themes) || hokku.ThemeId < 1 {
		return 0, store.ErrForeignKeyConstraint
	}
	if hokku.OwnerId > len(s.Users) || hokku.OwnerId < 1 {
		return 0, store.ErrForeignKeyConstraint
	}
	hokku.Id = len(s.Hokkus) + 1
	s.Hokkus = append(s.Hokkus, hokku)
	return hokku.Id, nil
}

func (s *TestStore) DeleteHokku(id int) error {
	if id < 0 || id > len(s.Hokkus)-1 {
		return store.ErrNoRecord
	}
	s.Hokkus[id] = s.Hokkus[len(s.Hokkus)-1]
	s.Hokkus = s.Hokkus[:len(s.Hokkus)-1]
	return nil
}

func (s *TestStore) UpdateHokku(hokku *models.Hokku) error {
	if hokku.Id < 0 || hokku.Id > len(s.Hokkus) {
		return store.ErrNoRecord
	}
	s.Hokkus[hokku.Id-1] = hokku
	return nil
}
