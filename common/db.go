package common

import (
	"gorm.io/gorm"
	"io"
)

func InitGormDB() *gorm.DB {
	return nil
}

type MyReader struct {
	Reader io.Reader
}

func NewMyReader(r io.Reader) *MyReader {
	return nil
}
