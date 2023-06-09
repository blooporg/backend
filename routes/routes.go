package routes

import (
	"net/http"

	"github.com/ibilalkayy/backend/controller"
	"github.com/ibilalkayy/backend/controller/account"
	"github.com/ibilalkayy/backend/controller/author_signup"
	"github.com/ibilalkayy/backend/controller/authors"
	"github.com/ibilalkayy/backend/controller/billing"
	"github.com/ibilalkayy/backend/controller/books"
	"github.com/ibilalkayy/backend/controller/loops"
	"github.com/ibilalkayy/backend/controller/shopped"
	"github.com/ibilalkayy/backend/controller/user"
	"github.com/ibilalkayy/backend/middleware"
)

func Routes() {
	http.HandleFunc("/", middleware.ErrorHandling(controller.Index))

	// User pages
	http.HandleFunc("/signin", middleware.ErrorHandling(user.Signin))
	http.HandleFunc("/signup", middleware.ErrorHandling(user.Signup))
	http.HandleFunc("/signout", middleware.ErrorHandling(user.Signout))

	// Loops page
	http.HandleFunc("/loops", middleware.ErrorHandling(loops.Loops))

	// Authors page
	http.HandleFunc("/authors", middleware.ErrorHandling(authors.Authors))

	// Books page
	http.HandleFunc("/books", middleware.ErrorHandling(books.Books))

	// Shopped pages
	http.HandleFunc("/cart", middleware.ErrorHandling(shopped.Cart))
	http.HandleFunc("/checkout", middleware.ErrorHandling(shopped.Checkout))

	// Author Signup pages
	http.HandleFunc("/author-signup", middleware.ErrorHandling(author_signup.AuthorSignup))
	http.HandleFunc("/author-pricing", middleware.ErrorHandling(author_signup.AuthorPricing))

	// Billing Info page
	http.HandleFunc("/billing-info", middleware.ErrorHandling(billing.BillingInfo))

	// Account page
	http.HandleFunc("/account", middleware.ErrorHandling(account.AccountPage))
}
