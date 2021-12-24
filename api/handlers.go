package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/EgorSkurihin/Hokku/models"
	"github.com/EgorSkurihin/Hokku/store"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /health [get]
func (api *APIServer) HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": "It`s alive!",
	})
}

// @Summary Get all hokkus
// @Description Get all hokkus
// @Tags Open routes
// @Accept json
// @Produce json
// @Param limit query int false "Sample size"
// @Param offset query int false "Number of items to skip"
// @Success 200 {array} models.Hokku
// @Failure 400 {object} echo.HTTPError "Bad query parameters"
// @Failure 500 {object} echo.HTTPError "Unexpected error"
// @Router /hokkus [get]
func (api *APIServer) GetHokkus(c echo.Context) error {
	var limit, offset int
	var err error
	l := c.QueryParam("limit")
	if l != "" {
		limit, err = strconv.Atoi(l)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Limit must be a number")
		}
	}
	if limit == 0 {
		limit = 10
	}
	o := c.QueryParam("offset")
	if o != "" {
		offset, err = strconv.Atoi(c.QueryParam("offset"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Offset must be a number")
		}
	}
	result, err := api.store.GetHokkus(limit, offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unexpected error")
	}
	return c.JSON(http.StatusOK, result)
}

// @Summary Get hokkus by athor
// @Description Get all hokkus of current author
// @Tags Open routes
// @Accept json
// @Produce json
// @Param limit query int false "Sample size"
// @Param offset query int false "Number of items to skip"
// @Param authorId path int true "Author id"
// @Success 200 {array} models.Hokku
// @Failure 400 {object} echo.HTTPError "Bad query parameters"
// @Failure 500 {object} echo.HTTPError "Unexpected error"
// @Router /hokkus/byAuthor/{authorId} [get]
func (api *APIServer) GetHokkusByAuthor(c echo.Context) error {
	var limit, offset int
	var err error
	l := c.QueryParam("limit")
	if l != "" {
		limit, err = strconv.Atoi(l)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Limit must be a number")
		}
	}
	if limit == 0 {
		limit = 10
	}
	o := c.QueryParam("offset")
	if o != "" {
		offset, err = strconv.Atoi(c.QueryParam("offset"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Offset must be a number")
		}
	}
	authorId, err := strconv.Atoi(c.Param("authorId"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request. ThemeID must be an integer")
	}
	result, err := api.store.GetHokkusByAuthor(authorId, limit, offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unexpected error")
	}
	return c.JSON(http.StatusOK, result)
}

// @Summary Get hokkus by theme
// @Description Get all hokkus of current author
// @Tags Open routes
// @Accept json
// @Produce json
// @Param limit query int false "Sample size"
// @Param offset query int false "Number of items to skip"
// @Param themeId path int true "thme Id"
// @Success 200 {array} models.Hokku
// @Failure 400 {object} echo.HTTPError "Bad query parameters"
// @Failure 500 {object} echo.HTTPError "Unexpected error"
// @Router /hokkus/byTheme/{themeId} [get]
func (api *APIServer) GetHokkusByTheme(c echo.Context) error {
	var limit, offset int
	var err error
	l := c.QueryParam("limit")
	if l != "" {
		limit, err = strconv.Atoi(l)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Limit must be a number")
		}
	}
	if limit == 0 {
		limit = 10
	}
	o := c.QueryParam("offset")
	if o != "" {
		offset, err = strconv.Atoi(c.QueryParam("offset"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Offset must be a number")
		}
	}
	themeId, err := strconv.Atoi(c.Param("themeId"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request. ThemeID must be an integer")
	}
	result, err := api.store.GetHokkusByTheme(themeId, limit, offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unexpected error")
	}
	return c.JSON(http.StatusOK, result)
}

// @Summary Get hokku
// @Description Get hokku by ID
// @Tags Open routes
// @Accept json
// @Produce json
// @Param id path  int  true  "id of hokku"
// @Success 200 {object} models.Hokku  "OK"
// @Failure 400 {object} echo.HTTPError "Bad request. Hokku ID must be an integer and larger than 0"
// @Failure 404 {object} echo.HTTPError "A hokku with the specified ID was not found"
// @Failure 500 {object} echo.HTTPError "Unexpected error"
// @Router /hokku/{id} [get]
func (api *APIServer) GetHokku(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request. Id must be an integer")
	}
	hokku, err := api.store.GetHokku(id)
	if err != nil {
		if errors.Is(err, store.ErrNoRecord) {
			return echo.NewHTTPError(http.StatusNotFound, "A hokku with the specified ID was not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Unexpected error")
	}
	return c.JSON(http.StatusOK, hokku)
}

// @Summary Post hokku
// @Security cookieAuth
// @Description Create new hokku in Store. Reurn location of new object in header
// @Tags Restricted routes
// @Accept json
// @Produce json
// @Param hokku body models.Hokku true "New Hokku"
// @Success 201 "Created"
// @Failure 400 {object} echo.HTTPError "Dont pass validation"
// @Failure 401 {object} echo.HTTPError "The request requires user authentication"
// @Failure 409 {object} echo.HTTPError "Foreign key constraint fails"
// @Failure 500 {object} echo.HTTPError "Unexpected error"
// @Router /restricted/hokku [post]
func (api *APIServer) PostHokku(c echo.Context) error {
	h := &models.Hokku{}
	decoder := json.NewDecoder(c.Request().Body)
	if err := decoder.Decode(&h); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request params")
	}
	if err := h.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Don`t pass the validation")
	}
	id, err := api.store.CreateHokku(h)
	if err != nil {
		if errors.Is(err, store.ErrForeignKeyConstraint) {
			return echo.NewHTTPError(http.StatusConflict, "Foreign key constraint fails")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Unexpected error")
	}
	c.Response().Header().Set("Location", fmt.Sprintf("/hokku/%d", id))
	return c.NoContent(http.StatusCreated)
}

// @Summary Delete hokku
// @Security cookieAuth
// @Description Delete hokku from Store
// @Tags Restricted routes
// @Accept json
// @Produce json
// @Param id path  int  true  "id of hokku"
// @Success 204 "Deleted succesfuly"
// @Failure 400 {object} echo.HTTPError "Bad request. Hokku ID must be an integer and larger than 0"
// @Failure 401 {object} echo.HTTPError "The request requires user authentication"
// @Failure 404 {object} echo.HTTPError "A hokku with the specified ID was not found"
// @Failure 500 {object} echo.HTTPError "Unexpected error"
// @Router /restricted/hokku/{id} [delete]
func (api *APIServer) DeleteHokku(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request. Id must be an integer")
	}
	if err = api.store.DeleteHokku(id); err != nil {
		if errors.Is(err, store.ErrNoRecord) {
			return echo.NewHTTPError(http.StatusNotFound, "A hokku with the specified ID was not exist")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Unexpected error")
	}
	return c.NoContent(http.StatusNoContent)
}

// @Summary Put hokku
// @Security cookieAuth
// @Description Update hokku in store
// @Tags Restricted routes
// @Accept json
// @Produce json
// @Param id path  int  true  "id of hokku"
// @Param hokku body models.Hokku true "Put Hokku"
// @Success 204 "OK"
// @Failure 401 {object} echo.HTTPError "The request requires user authentication"
// @Failure 404 {object} echo.HTTPError "Not Found"
// @Failure 500 {object} echo.HTTPError "Unexpected error"
// @Router /restricted/hokku [put]
func (api *APIServer) PutHokku(c echo.Context) error {
	h := &models.Hokku{}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	decoder := json.NewDecoder(c.Request().Body)
	if err := decoder.Decode(&h); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request params")
	}
	if err := h.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Don`t pass the validation")
	}
	h.Id = id
	if err := api.store.UpdateHokku(h); err != nil {
		if errors.Is(err, store.ErrNoRecord) {
			return echo.NewHTTPError(http.StatusNotFound, "A hokku with the specified ID was not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Unexpected error")
	}
	return c.NoContent(http.StatusNoContent)
}

// @Summary Get user
// @Description Get user by ID
// @Tags Open routes
// @Accept json
// @Produce json
// @Param id path  int  true  "id of user"
// @Success 200 {object} models.User  "OK"
// @Failure 400 {object} echo.HTTPError "Bad request. User ID must be an integer and larger than 0"
// @Failure 404 {object} echo.HTTPError "A user with the specified ID was not found"
// @Failure 500 {object} echo.HTTPError "Unexpected error"
// @Router /user/{id} [get]
func (api *APIServer) GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request. Id must be an integer")
	}
	user, err := api.store.GetUser(id)
	if err != nil {
		if errors.Is(err, store.ErrNoRecord) {
			return echo.NewHTTPError(http.StatusNotFound, "A user with the specified ID was not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Unexpected error")
	}
	return c.JSON(http.StatusOK, user)
}

// @Summary Post user
// @Description Create new user in Store. Return location of new user in header
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body models.User true "New User"
// @Success 201 "Created"
// @Failure 400 {object} echo.HTTPError "Dont pass validation"
// @Failure 409 {object} echo.HTTPError "User with this email already exists"
// @Failure 500 {object} echo.HTTPError "Unexpected error"
// @Router /user [post]
func (api *APIServer) PostUser(c echo.Context) error {
	u := &models.User{}
	decoder := json.NewDecoder(c.Request().Body)
	if err := decoder.Decode(&u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request params")
	}
	if err := u.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Don`t pass the validation")
	}
	if err := u.BeforeCreate(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unexpected error")
	}
	id, err := api.store.CreateUser(u)
	if err != nil {
		if errors.Is(err, store.ErrAlreadyExist) {
			return echo.NewHTTPError(http.StatusConflict, "User with this email already exists")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Unexpected error")
	}
	c.Response().Header().Set("Location", fmt.Sprintf("/user/%d", id))
	return c.NoContent(http.StatusCreated)
}

// @Summary Put user
// @Security cookieAuth
// @Description Update user in store
// @Tags Restricted routes
// @Accept json
// @Produce json
// @Param id path  int  true  "id of user"
// @Param user body models.User true "Put User"
// @Success 204 "OK"
// @Failure 401 {object} echo.HTTPError "The request requires user authentication"
// @Failure 404 {object} echo.HTTPError "A user with the specified ID was not found"
// @Failure 500 {object} echo.HTTPError "Unexpected error"
// @Router /restricted/user [put]
func (api *APIServer) PutUser(c echo.Context) error {
	u := &models.User{}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	decoder := json.NewDecoder(c.Request().Body)
	if err := decoder.Decode(&u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request params")
	}
	if err := u.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Don`t pass the validation")
	}
	u.Id = id
	if err := api.store.UpdateUser(u); err != nil {
		if errors.Is(err, store.ErrNoRecord) {
			return echo.NewHTTPError(http.StatusNotFound, "A user with the specified ID was not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Unexpected error")
	}
	return c.NoContent(http.StatusNoContent)
}

// @Summary Delete user
// @Security cookieAuth
// @Description Delete user from Store
// @Tags Restricted routes
// @Accept json
// @Produce json
// @Param id path  int  true  "id of user"
// @Success 204 "Deleted succesfuly"
// @Failure 400 {object} echo.HTTPError "Bad request. User ID must be an integer and larger than 0"
// @Failure 401 {object} echo.HTTPError "The request requires user authentication"
// @Failure 404 {object} echo.HTTPError "A user with the specified ID was not found"
// @Failure 500 {object} echo.HTTPError "Unexpected error"
// @Router /restricted/user/{id} [delete]
func (api *APIServer) DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	if err = api.store.DeleteUser(id); err != nil {
		if errors.Is(err, store.ErrNoRecord) {
			return echo.NewHTTPError(http.StatusNotFound)
		}
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusNoContent)
}

// @Summary Get all themes
// @Description Get all themes
// @Tags Open routes
// @Accept json
// @Produce json
// @Success 200 {array} models.Theme
// @Failure 500 {object} echo.HTTPError "Unexpected error"
// @Router /themes [get]
func (api *APIServer) GetThemes(c echo.Context) error {
	var err error
	result, err := api.store.GetThemes()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unexpected error")
	}
	return c.JSON(http.StatusOK, result)
}

// @Summary Authenticate
// @Description Login
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body models.User true "The user object can only contain email and password"
// @Success 200
// @Failure 400 {object} echo.HTTPError "Wrong email or passowrd"
// @Failure 404 {object} echo.HTTPError "A user with the specified Email was not found"
// @Failure 500 {object} echo.HTTPError "Unexpected error"
// @Router /login [post]
func (api *APIServer) Login(c echo.Context) error {
	formUser := &models.User{}
	decoder := json.NewDecoder(c.Request().Body)
	if err := decoder.Decode(&formUser); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request params")
	}
	dbUser, err := api.store.GetUserByEmail(formUser.Email)
	if err != nil {
		if errors.Is(err, store.ErrNoRecord) {
			return echo.NewHTTPError(http.StatusNotFound)
		}
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	err = bcrypt.CompareHashAndPassword([]byte(dbUser.HashedPassword), []byte(formUser.OpenPassword))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Wrong email or passowrd")
	}
	session, _ := api.sessionStore.Get(c.Request(), "session")
	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	session.Values["userId"] = dbUser.Id
	session.Save(c.Request(), c.Response())
	return c.NoContent(http.StatusOK)
}

/*
// @Summary Post theme
// @Description Create new user in Store. Return location of new user in header
// @Tags theme
// @Accept json
// @Produce json
// @Param theme body models.Theme true "New theme"
// @Success 201 "Created"
// @Failure 400 {object} echo.HTTPError "Dont pass validation"
// @Failure 409 {object} echo.HTTPError "Theme with this title already exists"
// @Failure 500 {object} echo.HTTPError "Unexpected error"
// @Router /admin/theme [post]
func (api *APIServer) PostTheme(c echo.Context) error {
	t := &models.Theme{}
	decoder := json.NewDecoder(c.Request().Body)
	if err := decoder.Decode(&t); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request params")
	}
	if err := t.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Don`t pass the validation")
	}
	id, err := api.store.CreateTheme(t)
	if err != nil {
		if errors.Is(err, store.ErrAlreadyExist) {
			return echo.NewHTTPError(http.StatusConflict, "Theme with this title already exists")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Unexpected error")
	}
	c.Response().Header().Set("Location", fmt.Sprintf("/theme/%d", id))
	return c.NoContent(http.StatusCreated)
}

// @Summary Delete theme
// @Description Delete theme from Store
// @Tags theme
// @Accept json
// @Produce json
// @Param id path  int  true  "id of theme"
// @Success 204 "Deleted succesfuly"
// @Failure 400 {object} echo.HTTPError "Bad request. Theme ID must be an integer and larger than 0"
// @Failure 404 {object} echo.HTTPError "A theme with the specified ID was not found"
// @Failure 500 {object} echo.HTTPError "Unexpected error"
// @Router /admin/theme/{id} [delete]
func (api *APIServer) DeleteTheme(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	if err = api.store.DeleteTheme(id); err != nil {
		if errors.Is(err, store.ErrNoRecord) {
			return echo.NewHTTPError(http.StatusNotFound)
		}
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusNoContent)
} */
