package repositories

import (
	"database/sql"

	"github.com/santhureinoo_agnos_backend/configs"
	"github.com/santhureinoo_agnos_backend/models"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", configs.DBString)
	if err != nil {
		panic(err)
	}
}

func SaveLog(entry models.LogEntry) {
	_, err := db.Exec("INSERT INTO logs(method, path, status_code, duration) VALUES($1, $2, $3, $4)",
		entry.Method, entry.Path, entry.StatusCode, entry.Duration)
	if err != nil {
		panic(err)
	}
}
