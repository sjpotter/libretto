package fixture

import (
	"fmt"
	"net/http"
	"testing"

	th "github.com/apcera/libretto/Godeps/_workspace/src/github.com/rackspace/gophercloud/testhelper"
	"github.com/apcera/libretto/Godeps/_workspace/src/github.com/rackspace/gophercloud/testhelper/client"
)

func SetupHandler(t *testing.T, url, method, requestBody, responseBody string, status int) {
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, method)
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		if requestBody != "" {
			th.TestJSONRequest(t, r, requestBody)
		}

		if responseBody != "" {
			w.Header().Add("Content-Type", "application/json")
		}

		w.WriteHeader(status)

		if responseBody != "" {
			fmt.Fprintf(w, responseBody)
		}
	})
}