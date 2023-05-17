package main

import (
	"fmt"
	"net/http"
	"relasi-go/middleware"
	"relasi-go/modules/categories"
	"relasi-go/modules/products"
	"relasi-go/modules/users"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/relasi-go"))
	if err != nil {
		panic("Failed to connect database")
	}

	db.AutoMigrate(&categories.Category{}, &products.Product{}, &users.User{})

	categoryRepo := categories.Repository{DB: db}
	categoryUsecase := categories.Usecase{Repo: categoryRepo}
	categoryHandler := categories.Handler{Usecase: categoryUsecase}

	productRepo := products.Repository{DB: db}
	productUsecase := products.Usecase{Repo: productRepo}
	productHandler := products.Handler{Usecase: productUsecase}

	usersRepo := users.Repository{DB: db}
	usersUsecase := users.Usecase{Repo: usersRepo}
	usersHandler := users.Handler{Usecase: usersUsecase}

	router := mux.NewRouter()

	//login
	router.HandleFunc("/login", usersHandler.Login).Methods("POST")

	// users
	router.HandleFunc("/register", usersHandler.AddUsers).Methods("POST")
	router.HandleFunc("/register", usersHandler.FindUsers).Methods("GET")
	router.HandleFunc("/register/{id}", usersHandler.DeleteUser).Methods("DELETE")

	// CRUD Product
	router.HandleFunc("/products", middleware.MiddlewareJWTAuthorization(productHandler.GetAllProducts)).Methods("GET")
	router.HandleFunc("/products/{id}", middleware.MiddlewareJWTAuthorization(productHandler.GetProductById)).Methods("GET")
	router.HandleFunc("/products", middleware.MiddlewareJWTAuthorization(productHandler.AddProduct)).Methods("POST")
	router.HandleFunc("/products/{id}", middleware.MiddlewareJWTAuthorization(productHandler.EditProduct)).Methods("PUT")
	router.HandleFunc("/products/{id}", middleware.MiddlewareJWTAuthorization(productHandler.DeleteProduct)).Methods("DELETE")

	// CRUD Category
	router.HandleFunc("/categories", middleware.MiddlewareJWTAuthorization(categoryHandler.GetAllCategory)).Methods("GET")
	router.HandleFunc("/categories/{id}", middleware.MiddlewareJWTAuthorization(categoryHandler.GetCategoryById)).Methods("GET")
	router.HandleFunc("/categories", middleware.MiddlewareJWTAuthorization(categoryHandler.AddCategory)).Methods("POST")
	router.HandleFunc("/categories/{id}", middleware.MiddlewareJWTAuthorization(categoryHandler.EditCategory)).Methods("PUT")
	router.HandleFunc("/categories/{id}", middleware.MiddlewareJWTAuthorization(categoryHandler.DeleteCategory)).Methods("DELETE")

	PORT := ":8080"
	fmt.Println("Starting server at localhost", PORT)
	http.ListenAndServe(PORT, router)

}
