package main

import (
	"carwash-bot/internal/bot"
	"carwash-bot/internal/storage"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	// 1. Загружаем переменные окружения
	if err := godotenv.Load(); err != nil {
		log.Fatal("Ошибка загрузки .env файла:", err)
	}

	// 2. Инициализируем БД
	db, err := storage.InitDB(os.Getenv("DB_PATH"))
	if err != nil {
		log.Fatal("Ошибка инициализации БД:", err)
	}
	defer db.Close()

	// 3. Запускаем бота
	bot.Start(
		os.Getenv("BOT_TOKEN"),
		os.Getenv("CHANNEL_ID"),
		db,
	)
}
