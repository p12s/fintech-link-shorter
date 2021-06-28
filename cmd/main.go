package main

import (
	"context"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"github.com/p12s/fintech-link-shorter/pkg/handler"
	"github.com/p12s/fintech-link-shorter/pkg/repository"
	"github.com/p12s/fintech-link-shorter/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler handler.Handler) error {
	http.HandleFunc("/long", handler.Long)
	http.HandleFunc("/short", handler.Short)

	http.NewServeMux()

	s.httpServer = &http.Server{
		Addr:           ":" + port,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error init config: %s\n", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s\n", err.Error())
	}

	db, err := repository.NewSqlite3DB(repository.Config{
		DriverName:     viper.GetString("db.driverName"),
		DataSourceName: viper.GetString("db.dataSourceName"),
		MaxFileSize:    viper.GetInt64("db.maxFileSize"),
	})
	if err != nil {
		logrus.Fatalf("Failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handler := handler.NewHandler(services)

	// server
	srv := new(Server)
	if err := srv.Run(viper.GetString("port"), *handler); err != nil { //, handlers.InitRoutes()
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}

	logrus.Print("Link shorter app started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("Link shorter app shutting down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}