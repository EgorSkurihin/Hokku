package mysql_store_test

import (
	"testing"

	"github.com/EgorSkurihin/Hokku/models"
	"github.com/EgorSkurihin/Hokku/store/mysql_store"
	"github.com/EgorSkurihin/Hokku/store/test_store"
	"github.com/stretchr/testify/assert"
)

func AddTestData(t *testing.T, s *mysql_store.MySqlStore) {
	lastThemeId := 0
	lastUserId := 0
	var err error

	for _, u := range test_store.Users {
		lastUserId, err = s.CreateUser(u)
		assert.NoError(t, err)
	}
	for _, th := range test_store.Themes {
		lastThemeId, err = s.CreateTheme(th)
		assert.NoError(t, err)
	}
	for i, h := range test_store.Hokkus {
		h.ThemeId = lastThemeId - (i % lastThemeId)
		h.OwnerId = lastUserId - (i % lastUserId)
		_, err := s.CreateHokku(h)
		assert.NoError(t, err)
	}
}

func TestGetHokkus(t *testing.T) {
	s, teardown := mysql_store.TestMysqlStore(t)
	defer teardown("users", "themes", "hokkus")
	AddTestData(t, s)

	res, err := s.GetHokkus(10, 0)
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestGetHokkusByTheme(t *testing.T) {
	s, teardown := mysql_store.TestMysqlStore(t)
	defer teardown("users", "themes", "hokkus")
	AddTestData(t, s)

	themeId := 1
	s.DB.QueryRow("SELECT MAX(id) FROM themes;").Scan(&themeId)
	res, err := s.GetHokkusByTheme(themeId, 10, 0)
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestGetHokkusByAuthor(t *testing.T) {
	s, teardown := mysql_store.TestMysqlStore(t)
	defer teardown("users", "themes", "hokkus")
	AddTestData(t, s)

	userId := 0
	s.DB.QueryRow("SELECT MAX(id) FROM users;").Scan(&userId)
	res, err := s.GetHokkusByAuthor(userId-1, 10, 0)
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestGetHokku(t *testing.T) {
	s, teardown := mysql_store.TestMysqlStore(t)
	defer teardown("users", "themes", "hokkus")
	AddTestData(t, s)

	res, err := s.GetHokku(1)
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestDeleteHokku(t *testing.T) {
	s, teardown := mysql_store.TestMysqlStore(t)
	defer teardown("users", "themes", "hokkus")
	AddTestData(t, s)

	err := s.DeleteHokku(1)
	assert.NoError(t, err)
}

func TestUpdateHokku(t *testing.T) {
	s, teardown := mysql_store.TestMysqlStore(t)
	defer teardown("users", "themes", "hokkus")
	AddTestData(t, s)

	h := &models.Hokku{Id: 1, Title: "qwewqewe", Content: "asdqweasd"}
	err := s.UpdateHokku(h)
	assert.NoError(t, err)
}

func TestGetUsers(t *testing.T) {
	s, teardown := mysql_store.TestMysqlStore(t)
	defer teardown("users", "themes", "hokkus")
	AddTestData(t, s)

	res, err := s.GetUsers()
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestGetUser(t *testing.T) {
	s, teardown := mysql_store.TestMysqlStore(t)
	defer teardown("users", "themes", "hokkus")
	AddTestData(t, s)

	res, err := s.GetUser(1)
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestDeleteUser(t *testing.T) {
	s, teardown := mysql_store.TestMysqlStore(t)
	defer teardown("users", "themes", "hokkus")
	AddTestData(t, s)

	err := s.DeleteUser(1)
	assert.NoError(t, err)
}

func TestUpdateUser(t *testing.T) {
	s, teardown := mysql_store.TestMysqlStore(t)
	defer teardown("users", "themes", "hokkus")
	AddTestData(t, s)

	u := &models.User{Id: 1, Name: "qwewqewe", Email: "qwe@email.com"}
	err := s.UpdateUser(u)
	assert.NoError(t, err)
}

func TestGetThemes(t *testing.T) {
	s, teardown := mysql_store.TestMysqlStore(t)
	defer teardown("users", "themes", "hokkus")
	AddTestData(t, s)

	res, err := s.GetThemes()
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestDeleteTheme(t *testing.T) {
	s, teardown := mysql_store.TestMysqlStore(t)
	defer teardown("users", "themes", "hokkus")
	AddTestData(t, s)

	err := s.DeleteTheme(1)
	assert.NoError(t, err)
}
