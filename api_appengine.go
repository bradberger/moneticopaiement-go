// +build !appengine

package moneticopaiement

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"google.golang.org/appengine/urlfetch"
)

var (
	// TimeoutDuration specifies the timeout length for external requests
	TimeoutDuration = 1 * time.Minute
)

// getClient (the GAE version) should return a urlfetch.Client
func (a *API) getClient(ctx context.Context) *http.Client {
	if a.Client != nil {
		return a.Client
	}
	ctx, _ = context.WithTimeout(ctx, TimeoutDuration)
	return urlfetch.Client(ctx)
}
