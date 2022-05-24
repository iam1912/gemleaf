package utils

import (
	"os"
	"path/filepath"

	"gorm.io/gorm"
)

func Transaction(db *gorm.DB, callback func(db *gorm.DB) error) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	if err := callback(tx); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func InferRootDir(name string) string {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	var infer func(d string) string
	infer = func(d string) string {
		if Exists(filepath.Join(d, name)) {
			return d
		}
		return infer(filepath.Dir(d))
	}
	return infer(cwd)
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
