package integration

import (
	"testing"

	"github.com/stretchr/testify/require"
)

/*
This integration test uses the fixtures in order to call directly to service and IO write in the DB
*/
func TestCreateRoom(t *testing.T) {
	t.Parallel()
	r, err := fixtures.createGenericRoom("Gamers-Online")
	require.NoError(t, err)
	require.NotEmpty(t, r.Name)
	require.Equal(t, r.Name, "Gamers-OnlineQ")
}
