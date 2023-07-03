package api

import (
	"net/http"

	"github.com/ibilalkayy/backend/handlers"
)

func SetupRoutes() {
	http.HandleFunc("/", handlers.HomePage)
	http.HandleFunc("/loops", handlers.LoopsPage)
	http.HandleFunc("/signup", handlers.SignupPage)
	http.HandleFunc("/login", handlers.LoginPage)
	http.HandleFunc("/logout", handlers.LogoutPage)

	http.HandleFunc("/account", handlers.AccountPage)
	http.HandleFunc("/authors", handlers.AuthorsPage)
	http.HandleFunc("/billing", handlers.BillingPage)

	http.HandleFunc("/books", handlers.BooksPage)
	http.HandleFunc("/cart", handlers.CartPage)
	http.HandleFunc("/checkout", handlers.CheckoutPage)
	http.HandleFunc("/pricing", handlers.PricingPage)
	http.HandleFunc("/authors-signup", handlers.AuthorsSignupPage)
}
