package config

// Config структура для хранения конфигурации приложения
type Config struct {
	DSN string // Data Source Name для подключения к PostgreSQL
}

// New инициализирует и возвращает новый экземпляр конфигурации
func New() *Config {
	return &Config{
		DSN: "host=localhost port=5433 user=postgres dbname=tdp sslmode=disable password=postgres",
	}
}
