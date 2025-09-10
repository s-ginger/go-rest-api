package sqlite

import (
	"log/slog"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"path/filepath"
)

var DB *gorm.DB

func InitDB(log *slog.Logger, path string)  {
	var err error
	dir := filepath.Dir(path)
	if err = os.MkdirAll(dir, os.ModePerm); err != nil {
		log.Error("Не удалось создать директорию", "dir", dir, "err", err)
		os.Exit(1)
	}

	DB, err = gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		log.Error("Не удалось инициализировать БД", "err", err)
		os.Exit(1)
  	}

	log.Info("Подключение к SQLite установлено", "path", path)
}


func Migrate(log *slog.Logger,models ...interface{}) {
	if DB == nil {
		log.Error("Db is nil")
		os.Exit(1)
	}
	if err := DB.AutoMigrate(models...); err != nil {
		log.Error("Migrate Error", "err", err)
		os.Exit(1)
	}
}





