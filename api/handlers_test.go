package api_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/EgorSkurihin/Hokku/api"
	"github.com/EgorSkurihin/Hokku/config"
	"github.com/EgorSkurihin/Hokku/store/test_store"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func testAPIServer() *api.APIServer {
	store := test_store.New()
	conf := &config.Server{
		Addr:  ":1323",
		Debug: true,
	}
	return api.New(conf, store)
}

func TestHealthCheck(t *testing.T) {
	api := testAPIServer()
	req := httptest.NewRequest(echo.GET, "/health", nil)
	rec := httptest.NewRecorder()
	c := api.Echo.NewContext(req, rec)

	if assert.NoError(t, api.HealthCheck(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestGetHokkus(t *testing.T) {
	api := testAPIServer()
	cases := []struct {
		name         string
		url          string
		expectedCode int
		expectedBody []byte
		isValid      bool
	}{
		{
			name:         "without params",
			url:          "/hokkus",
			expectedCode: http.StatusOK,
			expectedBody: test_store.MockHokkus(0, len(test_store.Hokkus)),
			isValid:      true,
		},
		{
			name:         "correct limit",
			url:          "/hokkus?limit=2",
			expectedCode: http.StatusOK,
			expectedBody: test_store.MockHokkus(0, 2),
			isValid:      true,
		},
		{
			name:         "correct offset",
			url:          "/hokkus?offset=2",
			expectedCode: http.StatusOK,
			expectedBody: test_store.MockHokkus(2, len(test_store.Hokkus)),
			isValid:      true,
		},
		{
			name:         "wrong limit",
			url:          "/hokkus?limit=qwe",
			expectedCode: http.StatusBadRequest,
			isValid:      false,
		},
		{
			name:         "wrong offset",
			url:          "/hokkus?offset=qwe",
			expectedCode: http.StatusBadRequest,
			isValid:      false,
		},
	}
	for _, cs := range cases {
		t.Run(cs.name, func(t *testing.T) {
			req := httptest.NewRequest(echo.GET, cs.url, nil)
			rec := httptest.NewRecorder()
			c := api.Echo.NewContext(req, rec)
			if cs.isValid {
				t.Log(cs.name)
				assert.NoError(t, api.GetHokkus(c))
				assert.Equal(t, cs.expectedBody, rec.Body.Bytes())

			}
			if !cs.isValid {
				assert.Error(t, api.GetHokkus(c))
				//assert.Equal(t, cs.expectedCode, rec.Code)
			}
		})
	}
}

func TestGetHokkusByTheme(t *testing.T) {
	api := testAPIServer()
	cases := []struct {
		name         string
		themeId      string
		limit        string
		offset       string
		expectedCode int
		expectedBody []byte
		isValid      bool
	}{
		{
			name:         "correct themeId",
			themeId:      "1",
			expectedCode: http.StatusOK,
			expectedBody: test_store.MockHokkusByTheme(1, 0, -1),
			isValid:      true,
		},
		{
			name:         "wrong themeId",
			themeId:      "qwe",
			expectedCode: http.StatusBadRequest,
			isValid:      false,
		},

		{
			name:         "correct limit",
			themeId:      "1",
			limit:        "2",
			expectedCode: http.StatusOK,
			expectedBody: test_store.MockHokkusByTheme(1, 0, 2),
			isValid:      true,
		},
		{
			name:         "correct offset",
			themeId:      "1",
			offset:       "2",
			expectedCode: http.StatusOK,
			expectedBody: test_store.MockHokkusByTheme(1, 2, -1),
			isValid:      true,
		},
		{
			name:         "wrong limit",
			themeId:      "1",
			limit:        "qwe",
			expectedCode: http.StatusBadRequest,
			isValid:      false,
		},
		{
			name:         "wrong offset",
			themeId:      "1",
			offset:       "qwe",
			expectedCode: http.StatusBadRequest,
			isValid:      false,
		},
	}
	for _, cs := range cases {
		t.Run(cs.name, func(t *testing.T) {
			q := make(url.Values)
			q.Set("limit", cs.limit)
			q.Set("offset", cs.offset)
			req := httptest.NewRequest(echo.GET, "/hokkus/byTheme?"+q.Encode(), nil)
			rec := httptest.NewRecorder()
			c := api.Echo.NewContext(req, rec)
			c.SetParamNames("themeId")
			c.SetParamValues(cs.themeId)
			if cs.isValid {
				assert.NoError(t, api.GetHokkusByTheme(c))
				assert.Equal(t, cs.expectedBody, rec.Body.Bytes())

			}
			if !cs.isValid {
				assert.Error(t, api.GetHokkusByTheme(c))
				//assert.Equal(t, cs.expectedCode, rec.Code)
			}
		})
	}
}

func TestGetHokkusByAuthor(t *testing.T) {
	api := testAPIServer()
	cases := []struct {
		name         string
		authorId     string
		limit        string
		offset       string
		expectedCode int
		expectedBody []byte
		isValid      bool
	}{
		{
			name:         "correct authorId",
			authorId:     "1",
			expectedCode: http.StatusOK,
			expectedBody: test_store.MockHokkusByUser(1, 0, -1),
			isValid:      true,
		},
		{
			name:         "wrong themeId",
			authorId:     "qwe",
			expectedCode: http.StatusBadRequest,
			isValid:      false,
		},
		{
			name:         "correct limit",
			authorId:     "1",
			limit:        "2",
			expectedCode: http.StatusOK,
			expectedBody: test_store.MockHokkusByUser(1, 0, 2),
			isValid:      true,
		},
		{
			name:         "correct offset",
			authorId:     "1",
			offset:       "2",
			expectedCode: http.StatusOK,
			expectedBody: test_store.MockHokkusByUser(1, 2, -1),
			isValid:      true,
		},
		{
			name:         "wrong limit",
			authorId:     "1",
			limit:        "qwe",
			expectedCode: http.StatusBadRequest,
			isValid:      false,
		},
		{
			name:         "wrong offset",
			authorId:     "1",
			offset:       "qwe",
			expectedCode: http.StatusBadRequest,
			isValid:      false,
		},
	}
	for _, cs := range cases {
		t.Run(cs.name, func(t *testing.T) {
			q := make(url.Values)
			q.Set("limit", cs.limit)
			q.Set("offset", cs.offset)
			req := httptest.NewRequest(echo.GET, "/hokkus/byAuthor?"+q.Encode(), nil)
			rec := httptest.NewRecorder()
			c := api.Echo.NewContext(req, rec)
			c.SetParamNames("authorId")
			c.SetParamValues(cs.authorId)
			if cs.isValid {
				assert.NoError(t, api.GetHokkusByAuthor(c))
				assert.Equal(t, cs.expectedBody, rec.Body.Bytes())

			}
			if !cs.isValid {
				assert.Error(t, api.GetHokkusByAuthor(c))
				//assert.Equal(t, cs.expectedCode, rec.Code)
			}
		})
	}
}

func TestGetHokku(t *testing.T) {
	api := testAPIServer()
	cases := []struct {
		name         string
		id           string
		expectedBody []byte
		isValid      bool
	}{
		{
			name:         "valid",
			id:           "1",
			expectedBody: test_store.MockHokku(1),
			isValid:      true,
		},
		{
			name:    "id not found",
			id:      "1000",
			isValid: false,
		},
		{
			name:    "invalid id",
			id:      "id",
			isValid: false,
		},
	}
	for _, cs := range cases {
		t.Run(cs.name, func(t *testing.T) {
			req := httptest.NewRequest(echo.GET, "/hokku/", nil)
			rec := httptest.NewRecorder()
			c := api.Echo.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues(cs.id)
			if cs.isValid {
				assert.NoError(t, api.GetHokku(c))
				assert.Equal(t, cs.expectedBody, rec.Body.Bytes())

			}
			if !cs.isValid {
				assert.Error(t, api.GetHokku(c))
				//assert.Equal(t, cs.expectedCode, rec.Code)
			}
		})
	}
}

func TestPostHokku(t *testing.T) {
	api := testAPIServer()
	cases := []struct {
		name    string
		reqBody string
		isValid bool
	}{
		{
			name:    "valid",
			reqBody: `{"title":"Example","content":"1","ownerId":1,"themeId":1}`,
			isValid: true,
		},
		{
			name:    "bad body params",
			reqBody: `{Error}`,
			isValid: false,
		},
		{
			name:    "validation error",
			reqBody: `{"title":"","content":"1","ownerId":1,"themeId":0}`,
			isValid: false,
		},
		{
			name:    "bad foreign key constraint",
			reqBody: `{"title":"Example","content":"1","ownerId":100,"themeId":1}`,
			isValid: false,
		},
	}
	for _, cs := range cases {
		t.Run(cs.name, func(t *testing.T) {
			req := httptest.NewRequest(echo.POST, "/hokku", strings.NewReader(cs.reqBody))
			rec := httptest.NewRecorder()
			c := api.Echo.NewContext(req, rec)
			if cs.isValid {
				assert.NoError(t, api.PostHokku(c))
			}
			if !cs.isValid {
				assert.Error(t, api.PostHokku(c))
			}
		})
	}
}

func TestDeleteHokku(t *testing.T) {
	api := testAPIServer()
	cases := []struct {
		name    string
		id      string
		isValid bool
	}{
		{
			name:    "valid",
			id:      "1",
			isValid: true,
		},
		{
			name:    "bad params",
			id:      "qwe",
			isValid: false,
		},
		{
			name:    "validation error",
			id:      "1000",
			isValid: false,
		},
	}
	for _, cs := range cases {
		t.Run(cs.name, func(t *testing.T) {
			req := httptest.NewRequest(echo.DELETE, "/hokku", nil)
			rec := httptest.NewRecorder()
			c := api.Echo.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues(cs.id)
			if cs.isValid {
				assert.NoError(t, api.DeleteHokku(c))
			}
			if !cs.isValid {
				assert.Error(t, api.DeleteHokku(c))
			}
		})
	}
}

func TestPutHokku(t *testing.T) {
	api := testAPIServer()
	cases := []struct {
		name    string
		reqBody string
		id      string
		isValid bool
	}{
		{
			name:    "valid",
			reqBody: `{"title":"Example","content":"1","ownerId":1,"themeId":1}`,
			id:      "1",
			isValid: true,
		},
		{
			name:    "bad body params",
			reqBody: `{Error}`,
			id:      "1",
			isValid: false,
		},
		{
			name:    "not found",
			reqBody: `{"title":"Example","content":"1","ownerId":1,"themeId":1}`,
			id:      "1000",
			isValid: false,
		},
		{
			name:    "id not valid",
			reqBody: `{"title":"Example","content":"1","ownerId":1,"themeId":1}`,
			id:      "qwe",
			isValid: false,
		},
		{
			name:    "validation error",
			reqBody: `{"title":"","content":"1","ownerId":1,"themeId":0}`,
			isValid: false,
		},
	}
	for _, cs := range cases {
		t.Run(cs.name, func(t *testing.T) {
			req := httptest.NewRequest(echo.PUT, "/hokku", strings.NewReader(cs.reqBody))
			rec := httptest.NewRecorder()
			c := api.Echo.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues(cs.id)
			if cs.isValid {
				assert.NoError(t, api.PutHokku(c))
			}
			if !cs.isValid {
				assert.Error(t, api.PutHokku(c))
			}
		})
	}
}

func TestGetUser(t *testing.T) {
	api := testAPIServer()
	cases := []struct {
		name         string
		id           string
		expectedBody []byte
		isValid      bool
	}{
		{
			name:         "valid",
			id:           "1",
			expectedBody: test_store.MockUser(1),
			isValid:      true,
		},
		{
			name:    "id not found",
			id:      "1000",
			isValid: false,
		},
		{
			name:    "invalid id",
			id:      "id",
			isValid: false,
		},
	}
	for _, cs := range cases {
		t.Run(cs.name, func(t *testing.T) {
			req := httptest.NewRequest(echo.GET, "/user/", nil)
			rec := httptest.NewRecorder()
			c := api.Echo.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues(cs.id)
			if cs.isValid {
				assert.NoError(t, api.GetUser(c))
				assert.Equal(t, cs.expectedBody, rec.Body.Bytes())

			}
			if !cs.isValid {
				assert.Error(t, api.GetUser(c))
			}
		})
	}
}

func TestPostUser(t *testing.T) {
	api := testAPIServer()
	cases := []struct {
		name    string
		reqBody string
		isValid bool
	}{
		{
			name:    "valid",
			reqBody: `{"email":"example@email.com","password":"123123231","name":"test"}`,
			isValid: true,
		},
		{
			name:    "bad params",
			reqBody: `{Error}`,
			isValid: false,
		},
		{
			name:    "validation error",
			reqBody: `{"email":"example.com","password":"1231","name":"test"}`,
			isValid: false,
		},
		{
			name:    "already exists",
			reqBody: `{"email":"example1@email.com","password":"123123231","name":"test"}`,
			isValid: false,
		},
	}
	for _, cs := range cases {
		t.Run(cs.name, func(t *testing.T) {
			req := httptest.NewRequest(echo.POST, "/user", strings.NewReader(cs.reqBody))
			rec := httptest.NewRecorder()
			c := api.Echo.NewContext(req, rec)
			if cs.isValid {
				assert.NoError(t, api.PostUser(c))
			}
			if !cs.isValid {
				assert.Error(t, api.PostUser(c))
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {
	api := testAPIServer()
	cases := []struct {
		name    string
		id      string
		isValid bool
	}{
		{
			name:    "valid",
			id:      "1",
			isValid: true,
		},
		{
			name:    "bad params",
			id:      "qwe",
			isValid: false,
		},
		{
			name:    "validation error",
			id:      "1000",
			isValid: false,
		},
	}
	for _, cs := range cases {
		t.Run(cs.name, func(t *testing.T) {
			req := httptest.NewRequest(echo.DELETE, "/user", nil)
			rec := httptest.NewRecorder()
			c := api.Echo.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues(cs.id)
			if cs.isValid {
				assert.NoError(t, api.DeleteUser(c))
			}
			if !cs.isValid {
				assert.Error(t, api.DeleteUser(c))
			}
		})
	}
}

func TestPutUser(t *testing.T) {
	api := testAPIServer()
	cases := []struct {
		name    string
		reqBody string
		id      string
		isValid bool
	}{
		{
			name:    "valid",
			reqBody: `{"email":"example@email.com","password":"123123231","name":"test"}`,
			id:      "1",
			isValid: true,
		},
		{
			name:    "bad body params",
			reqBody: `{Error}`,
			id:      "1",
			isValid: false,
		},
		{
			name:    "not found",
			reqBody: `{"email":"example@email.com","password":"123123231","name":"test"}`,
			id:      "1000",
			isValid: false,
		},
		{
			name:    "id not valid",
			reqBody: `{"email":"example@email.com","password":"123123231","name":"test"}`,
			id:      "qwe",
			isValid: false,
		},
		{
			name:    "validation error",
			reqBody: `{"email":"example.com","password":"1231","name":"test"}`,
			id:      "1",
			isValid: false,
		},
	}
	for _, cs := range cases {
		t.Run(cs.name, func(t *testing.T) {
			req := httptest.NewRequest(echo.PUT, "/user", strings.NewReader(cs.reqBody))
			rec := httptest.NewRecorder()
			c := api.Echo.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues(cs.id)
			if cs.isValid {
				assert.NoError(t, api.PutUser(c))
			}
			if !cs.isValid {
				assert.Error(t, api.PutUser(c))
			}
		})
	}
}

func TestGetThemes(t *testing.T) {
	api := testAPIServer()
	cases := []struct {
		name         string
		expectedBody []byte
		isValid      bool
	}{
		{
			name:         "valid",
			expectedBody: test_store.MockThemes(),
			isValid:      true,
		},
	}
	for _, cs := range cases {
		t.Run(cs.name, func(t *testing.T) {
			req := httptest.NewRequest(echo.GET, "/themes", nil)
			rec := httptest.NewRecorder()
			c := api.Echo.NewContext(req, rec)
			if cs.isValid {
				t.Log(cs.name)
				assert.NoError(t, api.GetThemes(c))
				assert.Equal(t, cs.expectedBody, rec.Body.Bytes())

			}
			if !cs.isValid {
				assert.Error(t, api.GetThemes(c))
			}
		})
	}
}

func TestLogin(t *testing.T) {
	api := testAPIServer()
	cases := []struct {
		name    string
		reqBody string
		isValid bool
	}{
		{
			name:    "valid",
			reqBody: `{"email":"example1@email.com","password":"Admin"}`,
			isValid: true,
		},
		{
			name:    "wrong email params",
			reqBody: `{"email":"wrong@email.com","password":"qweqweqwe"}`,
			isValid: false,
		},
	}
	for _, cs := range cases {
		t.Run(cs.name, func(t *testing.T) {
			req := httptest.NewRequest(echo.POST, "/login", strings.NewReader(cs.reqBody))
			rec := httptest.NewRecorder()
			c := api.Echo.NewContext(req, rec)
			if cs.isValid {
				assert.NoError(t, api.Login(c))
				assert.NotEmpty(t, rec.Result().Cookies())
			}
			if !cs.isValid {
				assert.Error(t, api.Login(c))
			}
		})
	}
}
