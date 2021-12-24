package test_store

import (
	"encoding/json"

	"github.com/EgorSkurihin/Hokku/models"
)

var (
	Users = []*models.User{
		{Id: 1, Email: "example1@email.com", Name: "Example1", HashedPassword: "$2a$10$0UWpDXKCrmtrUAEXWkczk.hJdHusoMlaZAu8wvbNenU/mR3kF9fsy"},
		{Id: 2, Email: "example2@email.com", Name: "Example2", HashedPassword: "$2a$10$0UWpDXKCrmtrUAEXWkczk.hJdHusoMlaZAu8wvbNenU/mR3kF9fsy"},
		{Id: 3, Email: "example3@email.com", Name: "Example3", HashedPassword: "$2a$10$0UWpDXKCrmtrUAEXWkczk.hJdHusoMlaZAu8wvbNenU/mR3kF9fsy"},
	}
	Hokkus = []*models.Hokku{
		{Id: 1, Title: "Title1", Content: "Content", OwnerId: 1, ThemeId: 1},
		{Id: 2, Title: "Title2", Content: "Content", OwnerId: 2, ThemeId: 2},
		{Id: 3, Title: "Title3", Content: "Content", OwnerId: 3, ThemeId: 1},
		{Id: 4, Title: "Title4", Content: "Content", OwnerId: 1, ThemeId: 2},
		{Id: 5, Title: "Title5", Content: "Content", OwnerId: 2, ThemeId: 1},
	}
	Themes = []*models.Theme{
		{Id: 1, Title: "exampleTheme1"},
		{Id: 2, Title: "exampleTheme2"},
	}
)

func MockThemes() []byte {
	res, _ := json.Marshal(Themes)
	res = append(res, 10)
	return res
}

func MockUsers() []byte {
	res, _ := json.Marshal(Users)
	res = append(res, 10)
	return res
}

func MockHokkus(left, right int) []byte {
	res, _ := json.Marshal(Hokkus[left:right])
	res = append(res, 10)
	return res
}

func MockHokkusByTheme(themeId, left, right int) []byte {
	hs := make([]*models.Hokku, 0)
	for _, h := range Hokkus {
		if h.ThemeId == themeId {
			hs = append(hs, h)
		}
	}
	var res []byte
	if right == -1 {
		res, _ = json.Marshal(hs[left:])
	} else {
		res, _ = json.Marshal(hs[left:right])
	}
	res = append(res, 10)
	return res
}

func MockHokkusByUser(userId, left, right int) []byte {
	hs := make([]*models.Hokku, 0)
	for _, h := range Hokkus {
		if h.OwnerId == userId {
			hs = append(hs, h)
		}
	}
	var res []byte
	if right == -1 {
		res, _ = json.Marshal(hs[left:])
	} else {
		res, _ = json.Marshal(hs[left:right])
	}
	res = append(res, 10)
	return res
}

func MockHokku(id int) []byte {
	res, _ := json.Marshal(Hokkus[id-1])
	res = append(res, 10)
	return res
}

func MockUser(id int) []byte {
	res, _ := json.Marshal(Users[id-1])
	res = append(res, 10)
	return res
}

func MockTheme(id int) []byte {
	res, _ := json.Marshal(Themes[id-1])
	res = append(res, 10)
	return res
}
