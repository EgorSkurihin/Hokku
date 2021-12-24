package mysql_store

import (
	"testing"

	"github.com/EgorSkurihin/Hokku/config"
)

func TestMysqlStore(t *testing.T) (*MySqlStore, func(...string)) {
	t.Helper()

	config, err := config.New("../../config/config.toml")
	if err != nil {
		t.Fatal(err)
	}
	config.Store.DBName = "hokkutest"
	store := New(&config.Store)
	if err := store.Open(); err != nil {
		t.Fatal(err)
	}
	return store, func(tables ...string) {
		if len(tables) > 0 {
			for _, table := range tables {
				if _, err := store.DB.Exec("SET FOREIGN_KEY_CHECKS = 0;"); err != nil {
					t.Fatal(err)
				}
				if _, err := store.DB.Exec("TRUNCATE " + table); err != nil {
					t.Fatal(err)
				}
				if _, err := store.DB.Exec("SET FOREIGN_KEY_CHECKS = 1;"); err != nil {
					t.Fatal(err)
				}
			}
		}
		store.Close()
	}
}
