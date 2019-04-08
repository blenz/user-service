package service

import (
	"github.com/blenz/user-service/internal/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

type service struct {
	router *mux.Router
	db     *gorm.DB
}

func New(db *gorm.DB) *service {

	s := &service{
		router: mux.NewRouter(),
		db:     db,
	}

	s.router.HandleFunc("/user", s.getAllUsers())

	return s
}

func (s *service) getAllUsers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("derp"))
		s.db.Create(&models.User{Username: "derp"})
	}
}

func (s *service) Run() {
	log.Print("Listening...")
	log.Fatalln(http.ListenAndServe(":5000", s.router))
}
