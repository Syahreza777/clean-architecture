package products

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	Usecase Usecase
}

func (handler Handler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	products, err := handler.Usecase.UcGetAllProducts()
	if err != nil {
		messageErr, _ := json.Marshal(map[string]string{"message": "data not found"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(messageErr)
		return
	}

	response, err := json.Marshal(products)
	if err != nil {
		messageErr, _ := json.Marshal(map[string]string{"message": "data cannot be converted to json"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(messageErr)
		return
	}

	w.Write(response)
}

func (handler Handler) GetProductById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]

	product, err := handler.Usecase.UcGetProductById(id)
	if err != nil {
		messageErr, _ := json.Marshal(map[string]string{"message": "data not found"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(messageErr)
		return
	}

	response, err := json.Marshal(product)
	if err != nil {
		messageErr, _ := json.Marshal(map[string]string{"message": "data cannot be converted to json"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(messageErr)
		return
	}

	w.Write(response)

}

func (handler Handler) AddProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var product Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		messageErr, _ := json.Marshal(map[string]string{"message": "failed to decode json"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(messageErr)
		return
	}

	err = handler.Usecase.UcAddProduct(product)
	if err != nil {
		messageErr, _ := json.Marshal(map[string]string{"message": "data was not successfully added"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(messageErr)
		return
	}

	response := map[string]string{"message": "data added successfully"}
	json.NewEncoder(w).Encode(response)

}

func (handler Handler) EditProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]

	var product Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		messageErr, _ := json.Marshal(map[string]string{"message": "failed to decode json"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(messageErr)
		return
	}

	err = handler.Usecase.UcEditProduct(id, product)
	if err != nil {
		messageErr, _ := json.Marshal(map[string]string{"message": "data cannot be changed"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(messageErr)
		return
	}

	response := map[string]string{"message": "data changed successfully"}
	json.NewEncoder(w).Encode(response)

}

func (handler Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]

	err := handler.Usecase.UcDeleteProduct(id)
	if err != nil {
		messageErr, _ := json.Marshal(map[string]string{"message": "data was not successfully deleted"})
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(messageErr)
		return
	}

	response := map[string]string{"message": "data deleted successfully"}
	json.NewEncoder(w).Encode(response)

}
