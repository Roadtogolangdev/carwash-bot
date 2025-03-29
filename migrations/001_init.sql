-- Таблица моек
CREATE TABLE IF NOT EXISTS car_washes (
                                          id INTEGER PRIMARY KEY AUTOINCREMENT,
                                          address TEXT NOT NULL,
                                          open_time INTEGER NOT NULL DEFAULT 8,
                                          close_time INTEGER NOT NULL DEFAULT 20
);

-- Таблица слотов
CREATE TABLE IF NOT EXISTS slots (
                                     id INTEGER PRIMARY KEY AUTOINCREMENT,
                                     car_wash_id INTEGER NOT NULL,
                                     date TEXT NOT NULL, -- Формат YYYY-MM-DD
                                     hour INTEGER NOT NULL, -- 8-20
                                     is_booked BOOLEAN NOT NULL DEFAULT FALSE,
                                     FOREIGN KEY (car_wash_id) REFERENCES car_washes(id)
    );

-- Начальные данные
INSERT OR IGNORE INTO car_washes (id, address) VALUES
    (1, 'Улица 1, дом 5'),
    (2, 'Улица 2, дом 10');