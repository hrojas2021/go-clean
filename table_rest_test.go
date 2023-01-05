/* package handlers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"Accounts/internal/entity"
	"Accounts/internal/iface/mock"
	"Accounts/internal/rest"
	"Accounts/internal/rest/private/handlers"
	"Accounts/internal/rest/private/model"
	serviceModel "Accounts/internal/service/model"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestUpdateAccount(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type Response struct {
		rest.ErrResponse
		Account model.Account
	}

	ID := uuid.New()
	errOpz := errors.New("opz")

	tests := []struct {
		name               string
		pre                func(m *mock.MockService)
		accountID          string
		body               *bytes.Buffer
		responseStatusCode int
		pos                func(r *Response)
	}{
		{
			name: "HTTP Request OK - FROZEN STATUS",
			pre: func(m *mock.MockService) {
				m.EXPECT().UpdateAccount(gomock.All(), gomock.All()).
					DoAndReturn(func(_ context.Context, payload *serviceModel.AccountUpdateStatusPayload) (*entity.Account, error) {
						return &entity.Account{
							ID:     ID,
							Status: payload.Status,
						}, nil
					})
			},
			accountID:          ID.String(),
			body:               bytes.NewBufferString(`{"status":"FROZEN"}`),
			responseStatusCode: http.StatusOK,
			pos: func(r *Response) {
				assert.Equal(t, ID, r.Account.ID)
				assert.Equal(t, entity.AccountStatusFrozen, r.Account.Status)
			},
		},
		{
			name:               "HTTP Request FAIL - WRONG PARSE UUID",
			pre:                func(m *mock.MockService) {},
			accountID:          "INVALID",
			body:               bytes.NewBufferString(`{"status":"FROZEN"}`),
			responseStatusCode: http.StatusBadRequest,
			pos: func(r *Response) {
				assert.Len(t, r.Error.Codes, 2)
				assert.Equal(t, "invalid_account_id", r.Error.Codes[0].Code)
				assert.Equal(t, "bad_request", r.Error.Codes[1].Code)
			},
		},
		{
			name:               "HTTP Request FAIL - Invalid JSON payload",
			pre:                func(m *mock.MockService) {},
			body:               bytes.NewBufferString(`invalid_value`),
			responseStatusCode: http.StatusBadRequest,
			pos: func(r *Response) {
				assert.Len(t, r.Error.Codes, 2)
				assert.Equal(t, "invalid_account_id", r.Error.Codes[0].Code)
				assert.Equal(t, "bad_request", r.Error.Codes[1].Code)
				assert.Equal(t, "invalid account ID; bad request", r.Error.Msg)
			},
		},
		{
			name: "HTTP Request FAIL - Service Fail",
			pre: func(m *mock.MockService) {
				m.EXPECT().UpdateAccount(gomock.All(), gomock.All()).Return(nil, errOpz)
			},
			accountID:          ID.String(),
			body:               bytes.NewBufferString(`{"status":"PENDING_REGISTRATION"}`),
			responseStatusCode: http.StatusInternalServerError,
			pos: func(r *Response) {
				assert.Len(t, r.Error.Codes, 1)
				assert.Equal(t, "Internal Server Error", r.Error.Msg)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			// server
			m := mock.NewMockService(ctrl)
			test.pre(m)
			h := handlers.New(m, new(rest.DefaultResp))

			r := http.NewServeMux()
			r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				params := httprouter.Params{
					{
						Key:   "id",
						Value: test.accountID,
					},
				}
				ctx := context.WithValue(r.Context(), httprouter.ParamsKey, params)
				h.UpdateAccount(w, r.WithContext(ctx))
			})
			ts := httptest.NewServer(r)
			defer ts.Close()

			// client
			req, err := http.NewRequest(http.MethodPut, ts.URL, test.body)
			assert.Nil(t, err)

			res, err := http.DefaultClient.Do(req)
			assert.Nil(t, err)

			assert.Equal(t, test.responseStatusCode, res.StatusCode)

			var response Response
			assert.Nil(t, json.NewDecoder(res.Body).Decode(&response))
			assert.Nil(t, res.Body.Close())
			test.pos(&response)
		})
	}
}
*/