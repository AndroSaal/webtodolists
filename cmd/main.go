package main

import (
	todo "ToDoApp"
	"ToDoApp/pkg/handler"
	"ToDoApp/pkg/repository"
	"ToDoApp/pkg/service"
	"log"
	"github.com/spf13/viper"
)

func main() {
	//Инициализация конфига
	if err := InitConfig(); err != nil {
		log.Fatalf("error occured while init config: %s", err.Error())
	}
	repos := repository.NewRepository()
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
