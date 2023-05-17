package users

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	Usecase Usecase
}

func (handler Handler) Login(w http.ResponseWriter, r *http.Request) {
	var userInput User
	if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
		messageErr, _ := json.Marshal(map[string]string{"message": "failed to decode json"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(messageErr)
		return
	}
	defer r.Body.Close()

	signedToken, err := handler.Usecase.Login(userInput.Username, userInput.Password)
	if err != nil {
		messageErr, _ := json.Marshal(map[string]string{"message": err.Error()})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(messageErr)
		return
	}

	response := map[string]string{"token": signedToken}
	json.NewEncoder(w).Encode(response)

}

func (handler Handler) AddUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// mengambil inputan dari json
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		messageErr, _ := json.Marshal(map[string]string{"message": "failed to decode json"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(messageErr)
		return
	}
	defer r.Body.Close()

	// hass pass menggunakan bcrypt
	err := handler.Usecase.AddUsers(&user)
	if err != nil {
		messageErr, _ := json.Marshal(map[string]string{"message": err.Error()})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(messageErr)
		return
	}

	response := map[string]string{"message": "user data added successfully"}
	json.NewEncoder(w).Encode(response)
}

func (handler Handler) FindUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user User
	users, err := handler.Usecase.FindUsers(&user)
	if err != nil {
		messageErr, _ := json.Marshal(map[string]string{"message": "data not found"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(messageErr)
		return
	}

	response, err := json.Marshal(users)
	if err != nil {
		messageErr, _ := json.Marshal(map[string]string{"message": "data cannot be converted to json"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(messageErr)
		return
	}

	w.Write(response)
}

func (handler Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]

	err := handler.Usecase.DeleteUser(id)
	if err != nil {
		messageErr, _ := json.Marshal(map[string]string{"message": "data was not successfully deleted"})
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(messageErr)
		return
	}

	response := map[string]string{"message": "data deleted successfully"}
	json.NewEncoder(w).Encode(response)
}
