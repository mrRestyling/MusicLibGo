package main

import (
	"context"
	"log"
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

	log.Println("Graceful shutdown server...")

	err := db.Db.Close()
	if err != nil {
		log.Fatal("error close db: ", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = h.E.Shutdown(ctx); err != nil {
		log.Fatal("error shutdown server: ", err)
	}

	log.Println("Server exiting")

}
