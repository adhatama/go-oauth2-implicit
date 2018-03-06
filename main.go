package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	username := "tama"
	password := "tama"
	clientID := "my-client-id"
	responseType := "token"
	redirectURI := "http://localhost:8000/callback.html"
	accessToken := "great-token"

	e.POST("/authorize", func(c echo.Context) error {
		reqUsername := c.FormValue("username")
		reqPassword := c.FormValue("password")
		reqClientID := c.FormValue("client_id")
		reqResponseType := c.FormValue("response_type")
		reqRedirectUri := c.FormValue("redirect_uri")
		reqState := c.FormValue("state")

		if reqUsername != username && reqPassword != password {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error_message": "Invalid username or password",
			})
		}

		if reqRedirectUri != redirectURI {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error_message": "Invalid redirect URI",
			})
		}

		if reqResponseType != responseType {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error_message": "Invalid response type",
			})
		}

		if reqClientID != clientID {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error_message": "Invalid client id",
			})
		}

		c.Response().Header().Set(echo.HeaderAuthorization, "Bearer "+accessToken)
		redirectURI += "#" + "access_token=" + accessToken + "&state=" + reqState + "&expires_in=3600"

		return c.Redirect(302, redirectURI)
	})

	e.GET("/callback", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "GREAT")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
