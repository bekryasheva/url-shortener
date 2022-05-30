package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"

	"github.com/bekryasheva/url-shortener/internal/app/storage"
	"github.com/bekryasheva/url-shortener/pkg"
)

type URLResponseBody struct {
	Url string `json:"url"`
}

type SaveRequestBody struct {
	Url string `json:"url"`
}

func SaveHandler(s storage.Storage, urlPrefix string) echo.HandlerFunc {
	return func(c echo.Context) error {
		d, err := GetBodyData(c)
		if err != nil {
			return echo.ErrBadRequest
		}

		id, err := s.Save(d)
		if err != nil {
			return echo.ErrInternalServerError
		}

		encodedID, err := pkg.EncodeBase63(id)
		if err != nil {
			return echo.ErrInternalServerError
		}

		shortenedURL := fmt.Sprintf("%s%s", urlPrefix, encodedID)

		response := &URLResponseBody{
			Url: shortenedURL,
		}

		return c.JSON(http.StatusOK, response)
	}
}

func GetBodyData(c echo.Context) (string, error) {
	var o SaveRequestBody

	decoder := json.NewDecoder(c.Request().Body)

	decoder.DisallowUnknownFields()

	err := decoder.Decode(&o)
	if err != nil {
		return "", err
	}

	_, err = url.ParseRequestURI(o.Url)
	if err != nil {
		return "", err
	}

	return o.Url, nil
}
