package config

var (
	logger *Logger
)

func Init() error {
	return nil
}

func GetLogger(prefix string) *Logger {
	logger = NewLogger(prefix)
	return logger
}
