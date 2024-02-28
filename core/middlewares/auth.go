package middlewares

import (
	"os"
	"strings"
	"tde/fiber-api/core/handlers"
	"tde/fiber-api/core/helpers"

	"github.com/gofiber/fiber/v2"
)

func JWTAuthenticate() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if os.Getenv("JWT_BYPASS") == "1" {
			return c.Next()
		}
		// fmt.Println(c.GetReqHeaders())
		bearer := c.Get("Authorization", "")
		// fmt.Println(clientToken)
		if bearer == "" {
			return handlers.HandleError(c, &handlers.UnauthorizedError{Message: "Not Authorized"})
		}
		clientToken := strings.Split(bearer, " ")
		if len(clientToken) < 2 {
			return handlers.HandleError(c, &handlers.UnauthorizedError{Message: "Not Authorized"})
		}

		claims, err := helpers.ValidateToken(&clientToken[1])
		if err != nil {
			return handlers.HandleError(c, &handlers.UnauthorizedError{Message: err.Error()})
		}
		// fmt.Println(claims)
		c.Set("email", claims.Email)
		c.Set("first_name", claims.FirstName)
		c.Set("last_name", claims.LastName)
		c.Set("username", claims.Username)
		err2 := c.Next()

		return err2
	}
}
