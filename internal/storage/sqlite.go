package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

// InitDB инициализирует подключение к SQLite
func InitDB(dbPath string) (*sql.DB, error) {
	// Открываем соединение
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("ошибка подключения к БД: %v", err)
	}

	// Проверяем соединение
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ошибка ping БД: %v", err)
	}

	// Создаём таблицы (если их нет)
	if err := createTables(db); err != nil {
		return nil, err
	}

	return db, nil
}

// createTables создаёт все необходимые таблицы
func createTables(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS car_washes (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			address TEXT NOT NULL,
			open_time INTEGER DEFAULT 8,
			close_time INTEGER DEFAULT 20
		);

		CREATE TABLE IF NOT EXISTS slots (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			car_wash_id INTEGER NOT NULL,
			date TEXT NOT NULL,
			hour INTEGER NOT NULL,
			is_booked BOOLEAN DEFAULT FALSE,
			car_model TEXT,
			car_number TEXT,
			user_id INTEGER,
			FOREIGN KEY (car_wash_id) REFERENCES car_washes(id)
		);

		INSERT OR IGNORE INTO car_washes (id, address) VALUES
			(1, 'Улица 1, дом 5'),
			(2, 'Улица 2, дом 10');
	`)
	return err
}
