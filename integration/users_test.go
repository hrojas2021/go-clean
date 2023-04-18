package integration

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/hugo.rojas/custom-api/internal/domain/entities"
	"github.com/stretchr/testify/require"
)

/*
This integration test uses the httpClient to simulate an HTTP request to an existing endpoint
*/
type UsersResponse struct {
	Users []entities.User `json:"users"`
}

func TestGetUsers(t *testing.T) {
	t.Parallel()
	var users UsersResponse
	resp, err := httpClient.Get(fmt.Sprintf("%s/api/users", localURL))
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)
	err = resp.UnmarshalJson(&users)
	require.NoError(t, err)
	// require.GreaterOrEqual(t, len(users.Users), 1)
}
