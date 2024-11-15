package main

import (
	todo "ToDoApp"
	"ToDoApp/pkg/handler"
	"ToDoApp/pkg/repository"
	"ToDoApp/pkg/service"
	"log"
	"os"

	//драйвер для работы pgsql - реализация интерфейса из sqlx
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/spf13/viper"
)

func main() {
	//Инициализация конфига
	if err := InitConfig(); err != nil {
		log.Fatalf("error occured while init config: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	//Инициализируем БД
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		log.Fatalf("failed to connect to database: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)

	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while run HTTP server: %s", err.Error())
	}
}

func InitConfig() error {
	//добавление пути к фалйлу конфигурации относительно корневой директории
	viper.AddConfigPath("configs")
	//имя файла конфигурации
	viper.SetConfigName("configs")
	//возфращаем функцию которая инициализирует значения из файла конфигурации
	return viper.ReadInConfig()
}
