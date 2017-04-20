package moneticopaiement

import "net/http"

// API is the Monetico Paiement API client
type API struct {
	Client *http.Client
}
