package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis"

	"github.com/gorilla/mux"
	"github.com/reynld/carbtographer/pkg/cache"
	"github.com/reynld/carbtographer/pkg/database"

	"github.com/gorilla/handlers"
)

// Server struct
type Server struct {
	DB     *sql.DB
	Router *mux.Router
	Cache  *redis.Client
}

// Initialize maps DB and Router instance to Server struct
func (s *Server) Initialize() {
	s.InitializeRouter()
	s.DB = database.InitializeDB()
	s.Cache = cache.InitializeCache()
}

// Run runs the server router
func (s *Server) Run() {
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))

	fmt.Printf("server live on port%s\n", port)
	log.Fatal(http.ListenAndServe(port, handlers.CORS()(s.Router)))
}
