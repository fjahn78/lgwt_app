package poker

import (
	"testing"
	"reflect"
	"net/http/httptest"
)

func assertNoError(t testing.TB, err error){
	t.Helper()
	if err != nil {
		t.Fatalf("didn't expect an error but got one, %v", err)
	}
}

// assertContentType checks that the response has the expected content type.
func assertContentType(t testing.TB, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	if response.Result().Header.Get("content-type") != want {
		t.Errorf("response did not have content-type of %s, got %v", want, response.Result().Header)
	}
}

func assertLeague(t testing.TB, got, want League) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func assertStatus(t *testing.T, got int, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got status %d, want %d", got, want)
	}
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q, want %q", got, want)
	}
}

func assertScoreEquals(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
			t.Errorf("got %d want %d", got, want)
		}
}
