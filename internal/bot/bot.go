package bot

import (
	"database/sql"
	"log"

	"github.com/go-telegram/bot"
)

// Start запускает Telegram-бота
func Start(token string, channelID string, db *sql.DB) {
	// Создаём новый экземпляр бота
	b, err := bot.NewBot(token)
	if err != nil {
		log.Fatal("Ошибка создания бота:", err)
	}

	// Регистрируем обработчики
	registerHandlers(b, db)

	log.Println("Бот запущен!")
	b.Start()
}

// registerHandlers добавляет все обработчики команд
func registerHandlers(b *bot.Bot, db *sql.DB) {
	b.Handle("/start", func(c bot.Context) error {
		return c.Send("Добро пожаловать! Нажмите кнопку ниже.",
			bot.WithReplyMarkup(mainMenu()))
	})
}

// mainMenu создаёт клавиатуру главного меню
func mainMenu() bot.ReplyMarkup {
	menu := &bot.ReplyKeyboardMarkup{
		Keyboard: [][]bot.KeyboardButton{
			{
				{Text: "📝 Записаться на мойку"},
			},
			{
				{Text: "🗑 Мои записи"},
			},
		},
		ResizeKeyboard: true,
	}
	return menu
}
