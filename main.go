package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/reynld/carbtographer/pkg/helpers"
	"github.com/reynld/carbtographer/pkg/routes"
)

var db *gorm.DB
var err error
var port string

func main() {
	godotenv.Load()
	db, err = helpers.InitDB()

	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	port = os.Getenv("PORT")
	if port == "" {
		port = "9001"
	}

	r := mux.NewRouter()
	routes.ConfigureRoutes(db, r)

	fmt.Println("server live on port: " + port)
	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS()(r)))
}