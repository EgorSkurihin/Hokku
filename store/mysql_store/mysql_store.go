package mysql_store

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/EgorSkurihin/Hokku/config"
	"github.com/EgorSkurihin/Hokku/models"
	"github.com/EgorSkurihin/Hokku/store"

	"github.com/go-sql-driver/mysql"
)

type MySqlStore struct {
	dsn string
	DB  *sql.DB
}

func New(conf *config.Store) *MySqlStore {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		conf.User, conf.Password, conf.Host, conf.Port, conf.DBName)
	return &MySqlStore{
		dsn: dsn,
	}
}

func (s *MySqlStore) Open() error {
	db, err := sql.Open("mysql", s.dsn)
	if err != nil {
		return err
	}
	s.DB = db
	err = s.DB.Ping()
	if err != nil {
		return err
	}
	return nil
}

func (s *MySqlStore) Close() {
	s.DB.Close()
}

func (s *MySqlStore) GetThemes() ([]*models.Theme, error) {
	themes := []*models.Theme{}
	rows, err := s.DB.Query("SELECT * FROM themes;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		t := &models.Theme{}
		err := rows.Scan(
			&t.Id,
			&t.Title,
		)
		if err != nil {
			return nil, err
		}
		themes = append(themes, t)
	}
	return themes, nil
}

func (s *MySqlStore) CreateTheme(theme *models.Theme) (int, error) {
	stmt := "INSERT INTO themes (title) VALUES (?)"
	res, err := s.DB.Exec(stmt, theme.Title)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (s *MySqlStore) DeleteTheme(id int) error {
	res, err := s.DB.Exec("DELETE FROM themes WHERE id = ?", id)
	if err != nil {
		return err
	}
	affeted, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affeted == 0 {
		return store.ErrNoRecord
	}
	return nil
}

func (s *MySqlStore) GetUsers() ([]*models.User, error) {
	users := []*models.User{}
	rows, err := s.DB.Query("SELECT * FROM users;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		u := &models.User{}
		err := rows.Scan(
			&u.Id,
			&u.Email,
			&u.Name,
			&u.HashedPassword,
			&u.Created,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (s *MySqlStore) GetUser(id int) (*models.User, error) {
	u := &models.User{}
	err := s.DB.QueryRow("SELECT * FROM users WHERE id=?", id).Scan(
		&u.Id,
		&u.Email,
		&u.Name,
		&u.HashedPassword,
		&u.Created,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, store.ErrNoRecord
		}
		return nil, err
	}
	return u, nil
}

func (s *MySqlStore) GetUserByEmail(email string) (*models.User, error) {
	u := &models.User{}
	err := s.DB.QueryRow("SELECT * FROM users WHERE email=?", email).Scan(
		&u.Id,
		&u.Email,
		&u.Name,
		&u.HashedPassword,
		&u.Created,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, store.ErrNoRecord
		}
		return nil, err
	}
	return u, nil
}

func (s *MySqlStore) CreateUser(user *models.User) (int, error) {
	stmt := "INSERT INTO users (name, email, password, created) VALUES (?, ?, ?, NOW())"
	res, err := s.DB.Exec(stmt, user.Name, user.Email, user.HashedPassword)
	if err != nil {
		me, ok := err.(*mysql.MySQLError)
		if ok {
			if me.Number == 1062 {
				return 0, store.ErrAlreadyExist
			}
		}
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (s *MySqlStore) DeleteUser(id int) error {
	res, err := s.DB.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return err
	}
	affeted, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affeted == 0 {
		return store.ErrNoRecord
	}
	return nil
}

func (s *MySqlStore) UpdateUser(user *models.User) error {
	stmt := `UPDATE users SET name = ?, email = ? WHERE id = ?`
	res, err := s.DB.Exec(stmt, user.Name, user.Email, user.Id)
	if err != nil {
		return err
	}
	affeted, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affeted == 0 {
		return store.ErrNoRecord
	}
	return nil
}

func (s *MySqlStore) GetHokkus(limit, offset int) ([]*models.Hokku, error) {
	hs := []*models.Hokku{}
	rows, err := s.DB.Query("SELECT * FROM hokkus LIMIT ? OFFSET ?;", limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		h := &models.Hokku{}
		err := rows.Scan(
			&h.Id,
			&h.Title,
			&h.Content,
			&h.Created,
			&h.OwnerId,
			&h.ThemeId,
		)
		if err != nil {
			return nil, err
		}
		hs = append(hs, h)
	}
	return hs, nil
}

func (s *MySqlStore) GetHokkusByAuthor(authorId, limit, offset int) ([]*models.Hokku, error) {
	hs := []*models.Hokku{}
	rows, err := s.DB.Query("SELECT * FROM hokkus WHERE theme = ? LIMIT ? OFFSET ?;", authorId, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		h := &models.Hokku{}
		err := rows.Scan(
			&h.Id,
			&h.Title,
			&h.Content,
			&h.Created,
			&h.OwnerId,
			&h.ThemeId,
		)
		if err != nil {
			return nil, err
		}
		hs = append(hs, h)
	}
	return hs, nil
}

func (s *MySqlStore) GetHokkusByTheme(themeId, limit, offset int) ([]*models.Hokku, error) {
	hs := []*models.Hokku{}
	rows, err := s.DB.Query("SELECT * FROM hokkus WHERE theme = ? LIMIT ? OFFSET ?;", themeId, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		h := &models.Hokku{}
		err := rows.Scan(
			&h.Id,
			&h.Title,
			&h.Content,
			&h.Created,
			&h.OwnerId,
			&h.ThemeId,
		)
		if err != nil {
			return nil, err
		}
		hs = append(hs, h)
	}
	return hs, nil
}

func (s *MySqlStore) GetHokku(id int) (*models.Hokku, error) {
	h := &models.Hokku{}
	err := s.DB.QueryRow("SELECT * FROM hokkus WHERE id=?;", id).Scan(
		&h.Id,
		&h.Title,
		&h.Content,
		&h.Created,
		&h.OwnerId,
		&h.ThemeId,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, store.ErrNoRecord
		}
		return nil, err
	}
	return h, nil
}

func (s *MySqlStore) CreateHokku(hokku *models.Hokku) (int, error) {
	stmt := "INSERT INTO hokkus (title, content, created, owner, theme) VALUES (?, ?, Now(), ?, ?)"
	res, err := s.DB.Exec(stmt, hokku.Title, hokku.Content, hokku.OwnerId, hokku.ThemeId)
	if err != nil {
		me, ok := err.(*mysql.MySQLError)
		if ok {
			if me.Number == 1452 {
				return 0, store.ErrForeignKeyConstraint
			}
		}
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (s *MySqlStore) DeleteHokku(id int) error {
	res, err := s.DB.Exec("DELETE FROM hokkus WHERE id = ?", id)
	if err != nil {
		return err
	}
	affeted, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affeted == 0 {
		return store.ErrNoRecord
	}
	return nil
}

func (s *MySqlStore) UpdateHokku(hokku *models.Hokku) error {
	stmt := `UPDATE hokkus SET title = ?, content = ?, created = NOW() WHERE id = ?`
	res, err := s.DB.Exec(stmt, hokku.Title, hokku.Content, hokku.Id)
	if err != nil {
		return err
	}
	affeted, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affeted == 0 {
		return store.ErrNoRecord
	}
	return nil
}
