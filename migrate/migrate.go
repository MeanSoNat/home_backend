package main

import (
	"fmt"
	"log"

	"github.com/Makeyabe/Home_Backend/initializers"
	"github.com/Makeyabe/Home_Backend/model"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	err := initializers.DB.AutoMigrate(&model.Student{}, &model.Teacher{}, &model.Booking{}, &model.Summary{}, &model.Formvisit{})
	if err != nil {
		log.Fatal("? Migration failed", err)
	}

	fmt.Println("? Migration complete")
}
