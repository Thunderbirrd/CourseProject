package main

import (
	"context"
	"github.com/Thunderbirrd/CourseProject/internal/config"
	"github.com/Thunderbirrd/CourseProject/internal/handler"
	"github.com/Thunderbirrd/CourseProject/internal/repository"
	"github.com/Thunderbirrd/CourseProject/internal/server"
	"github.com/Thunderbirrd/CourseProject/internal/service"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	cfg, err := config.New()
	if err != nil {
		logrus.Errorf("Error initializing jwtConfig: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     cfg.DBHost,
		Port:     cfg.DBPort,
		Username: cfg.DBUsername,
		Password: cfg.DBPassword,
		DBName:   cfg.DBName,
		SSLMode:  cfg.DBSSLMode,
	})
	if err != nil {
		logrus.Errorf("Failed to initialed db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)

	go func() {
		if err := srv.Run(cfg.HttpPort, handlers.InitRoutes()); err != nil {
			logrus.Fatalf("Error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("App started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("App shutting down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("Error occured on server shutting down: %s", err.Error())
	}
	if err := db.Close(); err != nil {
		logrus.Errorf("Error occured on db connection close: %s", err.Error())
	}

}
