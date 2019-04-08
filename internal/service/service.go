package service

import (
	"github.com/blenz/user-service/internal/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/matryer/respond"
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

	s.router.HandleFunc("/users", s.getUsers())
	s.router.HandleFunc("/users/{id}", s.getUser())

	return s
}

func (s *service) getUsers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		users := []*models.User{}
		s.db.Find(&users)

		respond.With(w, r, http.StatusOK, users)
	}
}

func (s *service) getUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		userID := mux.Vars(r)["id"]

		var user models.User
		result := s.db.First(&user, userID)

		if result.Error != nil {
			respond.With(w, r, http.StatusBadRequest, 
				error(http.StatusBadRequest, result.Error.Error()))
			return
		}
		if result.RecordNotFound() {
			respond.With(w, r, http.StatusNotFound, 
				error(http.StatusNotFound, "Not Found"))
			return
		}

		respond.With(w, r, http.StatusOK, user)
	}
}

func (s *service) Run() {
	log.Print("Listening...")
	log.Fatalln(http.ListenAndServe(":5000", s.router))
}
