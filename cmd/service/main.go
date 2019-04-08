package main

import (
	"github.com/blenz/user-service/internal/service"
	"github.com/blenz/user-service/internal/storage"
)

func main() {

	db := storage.InitDb()
	defer db.Close()

	service.New(db).Run()
}
