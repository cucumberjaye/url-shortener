package logger

import (
	"io"
	"log"
	"os"
)

// Разного уровня логгеры
var (
	InfoLogger    *log.Logger // для вывода информации
	WarningLogger *log.Logger // для вывода информации, которая может породить ошибки
	ErrorLogger   *log.Logger // для вывода ошибок
)

// Создает пользовательские логеры
func New() {

	InfoLogger = log.New(os.Stdout, "INFO:  ", log.Ldate|log.Ltime|log.Lmsgprefix)
	WarningLogger = log.New(os.Stdout, "WARNING:  ", log.Ldate|log.Ltime|log.Lmsgprefix)
	ErrorLogger = log.New(os.Stderr, "ERROR:  ", log.Ldate|log.Ltime|log.Lmsgprefix)
}

// Отключает логирование
func Discard() {
	InfoLogger.SetOutput(io.Discard)
	WarningLogger.SetOutput(io.Discard)
	ErrorLogger.SetOutput(io.Discard)
}
