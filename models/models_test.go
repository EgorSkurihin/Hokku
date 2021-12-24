package models_test

import (
	"testing"

	"github.com/EgorSkurihin/Hokku/models"
	"github.com/stretchr/testify/assert"
)

func testUser() *models.User {
	return &models.User{
		Email:        "example@email.com",
		Name:         "Example",
		OpenPassword: "password",
	}
}

func testHokku() *models.Hokku {
	return &models.Hokku{
		Title:   "Example",
		Content: "Example\nExample\nExample",
		OwnerId: 1,
		ThemeId: 1,
	}
}

func testTheme() *models.Theme {
	return &models.Theme{
		Title: "Example",
	}
}

func TestUserBeforeCreate(t *testing.T) {
	u := testUser()
	err := u.BeforeCreate()
	assert.NoError(t, err)
	assert.Equal(t, "", u.OpenPassword)
	assert.NotEqual(t, "", u.HashedPassword)
}

func TestUserCheckPassword(t *testing.T) {
	u := testUser()
	u.BeforeCreate()
	pass := "password"
	ok := u.CheckPassword(pass)
	assert.Equal(t, ok, true)
	pass = "pwd"
	ok = u.CheckPassword(pass)
	assert.Equal(t, ok, false)
}

func TestUserValidate(t *testing.T) {
	cases := []struct {
		name    string
		u       func() *models.User
		isValid bool
	}{
		{
			name:    "valid",
			u:       testUser,
			isValid: true,
		},
		{
			name: "empty email",
			u: func() *models.User {
				u := testUser()
				u.Email = ""
				return u
			},
			isValid: false,
		},
		{
			name: "short password",
			u: func() *models.User {
				u := testUser()
				u.OpenPassword = ""
				return u
			},
			isValid: false,
		},
		{
			name: "short name",
			u: func() *models.User {
				u := testUser()
				u.Name = ""
				return u
			},
			isValid: false,
		},
	}
	for _, c := range cases {
		if c.isValid {
			assert.NoError(t, c.u().Validate())
		} else {
			assert.Error(t, c.u().Validate())
		}
	}
	u := testUser()
	assert.NoError(t, u.Validate())
}

func TestHokkuValidate(t *testing.T) {
	cases := []struct {
		name    string
		h       func() *models.Hokku
		isValid bool
	}{
		{
			name:    "valid",
			h:       testHokku,
			isValid: true,
		},
		{
			name: "empty title",
			h: func() *models.Hokku {
				h := testHokku()
				h.Title = ""
				return h
			},
			isValid: false,
		},
		{
			name: "emty content",
			h: func() *models.Hokku {
				h := testHokku()
				h.Content = ""
				return h
			},
			isValid: false,
		},
		/* {
			name: "themeId < 1",
			h: func() *models.Hokku {
				h := testHokku()
				h.ThemeId = 0
				return h
			},
			isValid: false,
		},
		{
			name: "authorId < 1",
			h: func() *models.Hokku {
				h := testHokku()
				h.OwnerId = 0
				return h
			},
			isValid: false,
		}, */
	}
	for _, c := range cases {
		if c.isValid {
			assert.NoError(t, c.h().Validate())
		} else {
			assert.Error(t, c.h().Validate())
		}
	}
}

func TestThemeValidate(t *testing.T) {
	cases := []struct {
		name    string
		t       func() *models.Theme
		isValid bool
	}{
		{
			name:    "valid",
			t:       testTheme,
			isValid: true,
		},
		{
			name: "empty title",
			t: func() *models.Theme {
				t := testTheme()
				t.Title = ""
				return t
			},
			isValid: false,
		},
	}
	for _, c := range cases {
		if c.isValid {
			assert.NoError(t, c.t().Validate())
		} else {
			assert.Error(t, c.t().Validate())
		}
	}
}
