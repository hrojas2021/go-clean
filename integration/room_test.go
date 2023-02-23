package integration

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetRoom(t *testing.T) {
	t.Parallel()
	r, err := createGenericRoom("American Room")
	require.NoError(t, err)
	require.NotEmpty(t, r.Name)
	// import httpclient package https://github.com/federicoleon/go-httpclient
	//implement the client in the fixture location
	// make localhost call with getRoomByID
}
