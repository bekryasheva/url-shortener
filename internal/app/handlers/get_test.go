package handlers

import (
	"github.com/bekryasheva/url-shortener/pkg"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"

	"github.com/bekryasheva/url-shortener/internal/app/storage/mock"
)

var (
	getRequestParamValue       = "0000000001"
	getInvalidRequestBodyJSON  = "0####00001"
	getNotFoundRequestBodyJSON = "938494hfu3"
	getNotFoundId              = 141480565690862100
	getResponseBodyJSON        = `{"url":"https://www.youtube.com/"}`
)

func TestGetHandlerOK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mock.NewMockStorage(ctrl)
	handler := GetHandler(storage)

	req := httptest.NewRequest(http.MethodGet, "/url/", nil)
	rec := httptest.NewRecorder()

	e := echo.New()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("/:url")
	ctx.SetParamNames("url")
	ctx.SetParamValues(getRequestParamValue)

	storage.EXPECT().Get(int64(1)).Return(savedURL, nil)

	err := handler(ctx)

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, rec.Code)
	require.Equal(t, getResponseBodyJSON, strings.TrimSpace(rec.Body.String()))
}

func TestGetHandlerInvalidInput(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mock.NewMockStorage(ctrl)
	handler := GetHandler(storage)

	req := httptest.NewRequest(http.MethodPost, "/url", strings.NewReader(saveInvalidRequestBodyJSON))
	rec := httptest.NewRecorder()

	e := echo.New()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("/:url")
	ctx.SetParamNames("url")
	ctx.SetParamValues(getInvalidRequestBodyJSON)

	err := handler(ctx)

	require.Equal(t, echo.ErrBadRequest, err)
}

func TestGetHandlerNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mock.NewMockStorage(ctrl)
	handler := GetHandler(storage)

	req := httptest.NewRequest(http.MethodPost, "/url", strings.NewReader(saveInvalidRequestBodyJSON))
	rec := httptest.NewRecorder()

	e := echo.New()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("/:url")
	ctx.SetParamNames("url")
	ctx.SetParamValues(getNotFoundRequestBodyJSON)

	storage.EXPECT().Get(int64(getNotFoundId)).Return("", pkg.ErrNotFound)

	err := handler(ctx)

	require.Equal(t, echo.ErrNotFound, err)
}
