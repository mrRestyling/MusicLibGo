package main

import (
	"context"
	"music/internal/handlers"
	"music/internal/service"
	"music/internal/storage"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	// База данных
	connectionDB := storage.ConnectDB()
	defer connectionDB.Close() // TODO Перенести в GS
	db := storage.New(connectionDB)

	// Сервисный слой
	serv := service.New(db)

	// Хендлеры
	h := handlers.New(serv)
	h.Routes()

	// Запуск сервера
	go h.E.Start(":8080")

	// GS - Graceful Shutdown
	GS := make(chan os.Signal, 1)
	signal.Notify(GS, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-GS

	// TODO log

	err := db.Db.Close()
	if err != nil {
		// TODO log
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = h.E.Shutdown(ctx); err != nil {
		// TODO log
		panic(err)
	}

}
