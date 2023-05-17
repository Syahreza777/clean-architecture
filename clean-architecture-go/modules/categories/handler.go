package categories

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct{
	Usecase Usecase
}

func (handler Handler) GetAllCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	categories, err := handler.Usecase.GetAllCategories()
	if err != nil {
		messageErr, _ := json.Marshal(map[string]string{"message": "data not found"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(messageErr)
		return
	}

	response, err := json.Marshal(categories)
	if err != nil {
		messageErr, _ := json.Marshal(map[string]string{"message": "data cannot be converted to json"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(messageErr)
		return
	}

	w.Write(response)
}

func (handler Handler) GetCategoryById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]

	category, err := handler.Usecase.GetCategoryById(id)
	if err != nil {
		messageErr, _ := json.Marshal(map[string]string{"message": "data not found"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(messageErr)
		return
	}

	response, err := json.Marshal(category)
	if err != nil {
		messageErr, _ := json.Marshal(map[string]string{"message": "data cannot be converted to json"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(messageErr)
		return
	}

	w.Write(response)

}

func (handler Handler) AddCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var category Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		messageErr, _ := json.Marshal(map[string]string{"message": "failed to decode json"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(messageErr)
		return
	}

	err = handler.Usecase.AddCategory(category)
	if err != nil {
		messageErr, _ := json.Marshal(map[string]string{"message": "data was not successfully added"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(messageErr)
		return
	}

	response := map[string]string{"message": "data added successfully"}
	json.NewEncoder(w).Encode(response)

}

func (handler Handler) EditCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]

	var category Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		messageErr, _ := json.Marshal(map[string]string{"message": "failed to decode json"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(messageErr)
		return
	}

	err = handler.Usecase.EditCategory(id, category)
	if err != nil {
		messageErr, _ := json.Marshal(map[string]string{"message": "data cannot be changed"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(messageErr)
		return
	}

	response := map[string]string{"message": "data changed successfully"}
	json.NewEncoder(w).Encode(response)
}

func (handler Handler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]

	err := handler.Usecase.DeleteCategory(id)
	if err != nil {
		messageErr, _ := json.Marshal(map[string]string{"message": "data was not successfully deleted"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(messageErr)
		return
	}

	response := map[string]string{"message": "data deleted successfully"}
	json.NewEncoder(w).Encode(response)

}
