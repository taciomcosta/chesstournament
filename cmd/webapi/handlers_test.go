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
	service = shared.NewService(&data.MockRepository{}, &data.MockClubRepository{}, data.MockPlayerRepository{})
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
		description:    "should handle details of existing Club",
		request:        newRequestBuilder().withPathVar("id", "1").build(),
		handle:         GetClubDetailsHandler,
		expectedStatus: http.StatusOK,
	},
	{
		description:    "should handle details of unexistent Club",
		request:        newRequestBuilder().withPathVar("id", "unexistent").build(),
		handle:         GetClubDetailsHandler,
		expectedStatus: http.StatusNotFound,
	},
	{
		description:    "should handle creation of Club",
		request:        newRequestBuilder().withBody(toJSONString(data.MockValidClub)).build(),
		handle:         CreateClubHandler,
		expectedStatus: http.StatusCreated,
	},
	{
		description:    "should handle invalid input on Club creation",
		request:        newRequestBuilder().withBody(`{"invalid": "body"}`).build(),
		handle:         CreateClubHandler,
		expectedStatus: http.StatusBadRequest,
	},
	{
		description:    "should handle Club editing",
		request:        newRequestBuilder().withBody(toJSONString(data.MockValidClub)).withPathVar("id", "1").build(),
		handle:         EditClubHandler,
		expectedStatus: http.StatusOK,
	},
	{
		description:    "should handle invalid input on Club editing",
		request:        newRequestBuilder().withBody(toJSONString(`{"id":"1"}`)).withPathVar("id", "1").build(),
		handle:         EditClubHandler,
		expectedStatus: http.StatusBadRequest,
	},
	{
		description:    "should list Clubs",
		request:        newRequestBuilder().build(),
		handle:         ListClubsHandler,
		expectedStatus: http.StatusOK,
	},
	{
		description:    "should handle invalid filter options on Club listing",
		request:        newRequestBuilder().withQueryParam("$orderBy", "invalid").build(),
		handle:         ListClubsHandler,
		expectedStatus: http.StatusBadRequest,
	},
	{
		description:    "should handle Club deletion",
		request:        newRequestBuilder().withPathVar("id", "1").build(),
		handle:         DeleteClubHandler,
		expectedStatus: http.StatusOK,
	},
	{
		description:    "should handle deletion of unexistent Club",
		request:        newRequestBuilder().withPathVar("id", "-1").build(),
		handle:         DeleteClubHandler,
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
		request:        newRequestBuilder().withBody(toJSONString(shared.MockCreatePlayerDTO)).build(),
		handle:         CreatePlayerHandler,
		expectedStatus: http.StatusCreated,
	},
	{
		description:    "should handle invalid input on Player creation",
		request:        newRequestBuilder().withBody(`{"invalid": "body"}`).build(),
		handle:         CreatePlayerHandler,
		expectedStatus: http.StatusBadRequest,
	},
	{
		description:    "should handle Player deletion",
		request:        newRequestBuilder().withPathVar("id", "1").build(),
		handle:         DeletePlayerHandler,
		expectedStatus: http.StatusOK,
	},
	{
		description:    "should handle deletion of unexistent Player",
		request:        newRequestBuilder().withPathVar("id", "-1").build(),
		handle:         DeletePlayerHandler,
		expectedStatus: http.StatusNotFound,
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
