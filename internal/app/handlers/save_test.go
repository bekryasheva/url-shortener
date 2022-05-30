package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/bekryasheva/url-shortener/internal/app/storage/mock"
)

var (
	saveRequestBodyJSON        = `{"url":"https://www.youtube.com/"}`
	saveInvalidRequestBodyJSON = `{"url":"://www.youtube.com/"}`
	saveResponseBodyJSON       = `{"url":"http://localhost:8080/0000000001"}`
	savedURL                   = "https://www.youtube.com/"
	saveUrlPrefix              = "http://localhost:8080/"
)

func TestSaveHandlerOK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mock.NewMockStorage(ctrl)
	handler := SaveHandler(storage, saveUrlPrefix)

	req := httptest.NewRequest(http.MethodPost, "/url", strings.NewReader(saveRequestBodyJSON))
	rec := httptest.NewRecorder()

	e := echo.New()
	ctx := e.NewContext(req, rec)

	storage.EXPECT().Save(savedURL).Return(int64(1), nil)

	err := handler(ctx)

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, rec.Code)
	require.Equal(t, saveResponseBodyJSON, strings.TrimSpace(rec.Body.String()))
}

func TestSaveHandlerInvalidInput(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mock.NewMockStorage(ctrl)
	handler := SaveHandler(storage, saveUrlPrefix)

	req := httptest.NewRequest(http.MethodPost, "/url", strings.NewReader(saveInvalidRequestBodyJSON))
	rec := httptest.NewRecorder()

	e := echo.New()
	ctx := e.NewContext(req, rec)

	err := handler(ctx)

	require.Equal(t, echo.ErrBadRequest, err)
}
