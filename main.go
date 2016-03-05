package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/jinzhu/gorm"
	"github.com/rs/cors"
    _ "github.com/lib/pq"

	"apretasteconnect/config"
	"apretasteconnect/database"
)

func init() {
	config.ConfigData = config.ParseConfig()
}

func main() {
	var err error
	connectionString := fmt.Sprintf("user=%s host=%s dbname=%s password=%s sslmode=disable", config.ConfigData.DatabaseUser, config.ConfigData.DatabaseIP, config.ConfigData.DatabaseName, config.ConfigData.DatabasePassword)
	fmt.Println(connectionString)
	database.DB, err = gorm.Open("postgres", connectionString)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	database.AutoMigrate(database.DB)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "UPDATE"},
	})
	n := negroni.Classic()
	n.Use(c)
//	n.Use(negroni.HandlerFunc(middleware.RestrictedHandler))
	router := NewRouter()

	n.UseHandler(router)
	n.Run("0.0.0.0:8080")

}
