package main

import (
	"context"
	"net/http"
	"os"
	"time"
	"wc/pkg/handler"
	"wc/pkg/repository"
	"wc/pkg/service"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Error loading env variables: %s\n", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode: os.Getenv("DB_SSLMode"),
	})
	if err != nil {
		logrus.Fatalf("Failed to initialize db: %s\n", err.Error())
	}
	defer db.Close()

	//инициализируем репозиторий-сервис-хендлер
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(Server)

	if err := srv.Run(os.Getenv("APP_PORT"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error while running http server: %s\n", err.Error())
	}

}

// проработать с сервером
// Server - сервер REST-API

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return s.httpServer.ListenAndServe()
}

// Shutdown - grace-full- выключение
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
