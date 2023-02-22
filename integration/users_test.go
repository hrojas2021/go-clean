package integration

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetUsers(t *testing.T) {
	t.Parallel()
	users, err := fixture.ListUser(ctx)
	require.NoError(t, err)
	require.Greater(t, len(users), 0)
}
