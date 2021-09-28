// *build unit
package main

import (
	"testing"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

func TestHTTPRequest(t *testing.T) {
	handler := function(w, http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "{\"status\": \"good\"}")
	}

	req := httptest.NewRequest("GET", "http://localhost:8080/Up", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))

	if 200 != res.StatusCode {
		t.Fatal("Status Code Not OK/200")
	}
}