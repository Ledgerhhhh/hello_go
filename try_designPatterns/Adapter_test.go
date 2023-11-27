package try_designPatterns

import (
	"fmt"
	"testing"
)

type Storage interface {
	Save(data string)
}

type FileStorage struct{}

func (f FileStorage) Save(data string) {
	fmt.Println("保存数据到文件:", data)
}

type DatabaseStorage struct{}

func (d DatabaseStorage) StoreData(data string) {
	fmt.Println("存储数据到数据库:", data)
}

type DatabaseAdapter struct {
	DatabaseStorage *DatabaseStorage
}

func (da DatabaseAdapter) Save(data string) {
	da.DatabaseStorage.StoreData(data)
}

func TestStorage(t *testing.T) {
	// 使用文件系统存储
	fileStorage := &FileStorage{}
	var storage Storage = fileStorage
	storage.Save("Hello from FileStorage")

	// 使用适配器切换数据库存储
	databaseStorage := &DatabaseStorage{}
	databaseAdapter := &DatabaseAdapter{DatabaseStorage: databaseStorage}
	databaseAdapter.Save("Hello from DatabaseAdapter")
}
