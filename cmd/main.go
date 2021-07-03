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
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/p12s/fintech-link-shorter/docs"
)

// @title Link shorter API
// @version 0.0.1
// @description This is an API Server for link shorter

// @host localhost:80
// @BasePath /
// @query.collection.format multi

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @termsOfService http://swagger.io/terms/

// @x-extension-openapi {"example": "value on a json format"}
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
	handlers := handler.NewHandler(services)

	srv := new(Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), *handlers); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

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

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler handler.Handler) error {
	http.HandleFunc("/long", handler.Long)
	http.HandleFunc("/short", handler.Short)
	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)
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
