package handlers

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/bekryasheva/url-shortener/internal/app/storage"
	"github.com/bekryasheva/url-shortener/pkg"
)

func RedirectHandler(s storage.Storage) echo.HandlerFunc {
	return func(c echo.Context) error {
		shortenedURL := c.Param("url")

		id, err := pkg.DecodeBase63(shortenedURL)
		if err != nil {
			return echo.ErrBadRequest
		}

		url, err := s.Get(id)
		if err != nil {
			if errors.Is(err, pkg.ErrNotFound) {
				return echo.ErrNotFound
			}
			return echo.ErrInternalServerError
		}

		return c.Redirect(http.StatusMovedPermanently, url)
	}
}
