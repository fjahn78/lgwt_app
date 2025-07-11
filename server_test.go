package poker

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		nil,
		nil,
	}
	server := NewPlayerServer(&store)

	tests := []struct {
		name               string
		player             string
		expectedHttpStatus int
		expectedScore      string
	}{
		{
			name:               "Returns Pepper's score",
			player:             "Pepper",
			expectedHttpStatus: http.StatusOK,
			expectedScore:      "20",
		},
		{
			name:               "Returns Floyd's score",
			player:             "Floyd",
			expectedHttpStatus: http.StatusOK,
			expectedScore:      "10",
		},
		{
			name:               "Returns 404 on missing players",
			player:             "Apollo",
			expectedHttpStatus: http.StatusNotFound,
			expectedScore:      "0",
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			request := NewGetScoreRequest(tt.player)
			response := httptest.NewRecorder()

			server.ServeHTTP(response, request)

			AssertStatus(t, response.Code, tt.expectedHttpStatus)
			AssertResponseBody(t, response.Body.String(), tt.expectedScore)
		})
	}

}

func TestStoreWins(t *testing.T) {
	store := &StubPlayerStore{
		map[string]int{},
		nil,
		nil,
	}
	server := NewPlayerServer(store)

	t.Run("it records wins on POST", func(t *testing.T) {
		player := "Pepper"

		request := NewPostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		AssertStatus(t, response.Code, http.StatusAccepted)

		AssertPlayerWin(t, store, player)
	})
}

func TestLeague(t *testing.T) {
	t.Run("it returns the league table as JSON", func(t *testing.T) {
		wantedLeague := League{
			{"Chleo", 32},
			{"Chris", 20},
			{"Tiest", 14},
		}

		store := &StubPlayerStore{nil, nil, wantedLeague}
		server := NewPlayerServer(store)

		request := NewLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		var got = GetLeagueFromResponse(t, response.Body)

		AssertContentType(t, response, jsonContentType)

		AssertStatus(t, response.Code, http.StatusOK)

		AssertLeague(t, got, wantedLeague)

	})
}
