package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/taciomcosta/chesstournament/internal/data"
	"github.com/taciomcosta/chesstournament/internal/shared"
)

func TestMain(m *testing.M) {
	service = shared.NewService(&data.MockRepository{}, &data.MockChessClubRepository{}, data.MockPlayerRepository{})
	os.Exit(m.Run())
}

type handlerTest struct {
	description    string
	request        *http.Request
	handle         http.HandlerFunc
	expectedStatus int
}

var handlerTests []handlerTest = []handlerTest{
	{
		description:    "should handle details of existing Chessclub",
		request:        newRequestBuilder().withPathVar("id", "1").build(),
		handle:         GetChessclubDetailsHandler,
		expectedStatus: http.StatusOK,
	},
	{
		description:    "should handle details of unexistent Chessclub",
		request:        newRequestBuilder().withPathVar("id", "unexistent").build(),
		handle:         GetChessclubDetailsHandler,
		expectedStatus: http.StatusNotFound,
	},
	{
		description:    "should handle creation of Chessclub",
		request:        newRequestBuilder().withBody(toJSONString(data.MockValidChessClub)).build(),
		handle:         CreateChessclubHandler,
		expectedStatus: http.StatusCreated,
	},
	{
		description:    "should handle invalid input on Chessclub creation",
		request:        newRequestBuilder().withBody(`{"invalid": "body"}`).build(),
		handle:         CreateChessclubHandler,
		expectedStatus: http.StatusBadRequest,
	},
	{
		description:    "should handle Chessclubt editing",
		request:        newRequestBuilder().withBody(toJSONString(data.MockValidChessClub)).withPathVar("id", "1").build(),
		handle:         EditChessclubHandler,
		expectedStatus: http.StatusOK,
	},
	{
		description:    "should handle invalid input on Chessclub editing",
		request:        newRequestBuilder().withBody(toJSONString(`{"id":"1"}`)).withPathVar("id", "1").build(),
		handle:         EditChessclubHandler,
		expectedStatus: http.StatusBadRequest,
	},
	{
		description:    "should list Chessclubs",
		request:        newRequestBuilder().build(),
		handle:         ListChessclubsHandler,
		expectedStatus: http.StatusOK,
	},
	{
		description:    "should handle invalid filter options on Chessclub listing",
		request:        newRequestBuilder().withQueryParam("$orderBy", "invalid").build(),
		handle:         ListChessclubsHandler,
		expectedStatus: http.StatusBadRequest,
	},
	{
		description:    "should handle Chessclub deletion",
		request:        newRequestBuilder().withPathVar("id", "1").build(),
		handle:         DeleteChessclubHandler,
		expectedStatus: http.StatusOK,
	},
	{
		description:    "should handle deletion of unexistent Chessclub",
		request:        newRequestBuilder().withPathVar("id", "-1").build(),
		handle:         DeleteChessclubHandler,
		expectedStatus: http.StatusNotFound,
	},
	{
		description:    "should handle details of existing Player",
		request:        newRequestBuilder().withPathVar("id", "1").build(),
		handle:         GetPlayerDetailsHandler,
		expectedStatus: http.StatusOK,
	},
	{
		description:    "should handle details of unexistent Player",
		request:        newRequestBuilder().withPathVar("id", "unexistent").build(),
		handle:         GetPlayerDetailsHandler,
		expectedStatus: http.StatusNotFound,
	},
	{
		description:    "should handle creation of Player",
		request:        newRequestBuilder().withBody(toJSONString(data.MockValidPlayer)).build(),
		handle:         CreatePlayerHandler,
		expectedStatus: http.StatusCreated,
	},
	{
		description:    "should handle invalid input on Player creation",
		request:        newRequestBuilder().withBody(`{"invalid": "body"}`).build(),
		handle:         CreatePlayerHandler,
		expectedStatus: http.StatusBadRequest,
	},
}

func TestHandlers(t *testing.T) {
	for _, test := range handlerTests {
		testHandler(t, test)
	}
}

func testHandler(t *testing.T, test handlerTest) {
	recorder := httptest.NewRecorder()
	test.handle(recorder, test.request)
	if recorder.Code != test.expectedStatus {
		t.Error(test.description)
	}
}
