package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/DmitriiTrifonov/book-trck/internal/pkg/types"

	"github.com/DmitriiTrifonov/book-trck/internal/config"
	"github.com/DmitriiTrifonov/book-trck/internal/pkg/database"
	"github.com/DmitriiTrifonov/book-trck/internal/pkg/repo/users"
	"github.com/kelseyhightower/envconfig"
)

const configPrefix = "BOOK_TRCK"

func main() {
	var cfg config.Config
	err := envconfig.Process(configPrefix, &cfg)
	if err != nil {
		log.Fatal(err)
	}

	dataStorage, err := database.NewDataStorage(cfg.GetDSN())
	if err != nil {
		log.Fatal(err)
	}

	userRepo, err := users.NewUserRepo(dataStorage)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/singup", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		user := types.User{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&user)
		err = userRepo.Add(user)

		responseText := fmt.Sprintf("Registered: %s", user.UserName)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			responseText = err.Error()
		}
		fmt.Fprintf(w, "%s", responseText)
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		user := types.User{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&user)
		_, err = userRepo.Get(user)

		responseText := fmt.Sprintf("Found: %s", user.UserName)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			responseText = err.Error()
		}
		fmt.Fprintf(w, "%s", responseText)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
