package service

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/jinzhu/gorm"
	"github.com/blenz/user-service/internal/models"
)

type service struct {
	router *mux.Router
	db *gorm.DB
}

func New(db *gorm.DB) *service {

	s := &service{
		router: mux.NewRouter(),
		db: db,
	}

	s.router.HandleFunc("/user", s.getAllUsers())

	return s
}

func (s *service) getAllUsers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("derp"))
		s.db.Create(&models.User{ Username: "derp" })
    }
}

func (s *service) Run() {
	log.Print("Listening...")
    log.Fatalln(http.ListenAndServe(":5000", s.router))
}