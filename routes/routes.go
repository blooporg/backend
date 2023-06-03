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
)

func Routes() {
	http.HandleFunc("/", controller.Index)

	// User pages
	http.HandleFunc("/signin", user.Signin)
	http.HandleFunc("/signup", user.Signup)
	http.HandleFunc("/signout", user.Signout)

	// Loops page
	http.HandleFunc("/loops", loops.Loops)

	// Authors page
	http.HandleFunc("/authors", authors.Authors)

	// Books page
	http.HandleFunc("/books", books.Books)

	// Shopped pages
	http.HandleFunc("/cart", shopped.Cart)
	http.HandleFunc("/checkout", shopped.Checkout)

	// Author Signup pages
	http.HandleFunc("/author-signup", author_signup.AuthorSignup)
	http.HandleFunc("/author-pricing", author_signup.AuthorPricing)

	// Billing Info page
	http.HandleFunc("/billing-info", billing.BillingInfo)

	// Account page
	http.HandleFunc("/account", account.AccountPage)
}
