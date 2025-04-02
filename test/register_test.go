package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"RestuarantBackend/custom"
	dto "RestuarantBackend/models/dto"

	"github.com/stretchr/testify/assert"
)

func TestRegisterHandlers(t *testing.T) {
	router := SetUpRoutes()
	signupSamples := []struct {
		Name               string
		Body               dto.SignupRequest
		ExpectedStatusCode int
		ExpectedErrMessage string
	}{
		{
			Name: "Happy Ending",
			Body: dto.SignupRequest{
				PhoneNumber: "0702830141",
				Password:    "HJ10xugb123*",
				Email:       "chuongnguyen16112003@gmail.com",
				FullName:    "Chuong Nguyen",
			},
			ExpectedStatusCode: http.StatusOK,
			ExpectedErrMessage: `{"Message":"Success! Please wait..."}`,
		},
		{
			Name:               "Empty Fields",
			Body:               dto.SignupRequest{},
			ExpectedStatusCode: http.StatusBadRequest,
			ExpectedErrMessage: fmt.Sprintf(`{"error":"%s"}`, custom.ErrEmptySignupRequest),
		},
		{
			Name: "Invalid Phone Format",
			Body: dto.SignupRequest{
				PhoneNumber: "1702830141222200000",
				Password:    "HJ10xugb123*",
				Email:       "chuongnguyen16112004@gmail.com",
				FullName:    "Chuong Nguyen",
			},
			ExpectedStatusCode: http.StatusBadRequest,
			ExpectedErrMessage: fmt.Sprintf(`{"error":"%s"}`, custom.ErrInValidPhone),
		},
	}

	for _, tc := range signupSamples {
		t.Run(tc.Name, func(t *testing.T) {
			bodyBytes, _ := json.Marshal(tc.Body)
			req, err := http.NewRequest(http.MethodPost, "/api/v1/users/signup", bytes.NewBuffer(bodyBytes))
			assert.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")

			rec := httptest.NewRecorder()
			router.ServeHTTP(rec, req)

			assert.Equal(t, tc.ExpectedStatusCode, rec.Code)
			assert.Equal(t, tc.ExpectedErrMessage, rec.Body.String())
		})
	}
}
