package middlewares

import (
	"him/fiber-api/core/handlers"
	"him/fiber-api/core/helpers"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func JWTAuthenticate() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if os.Getenv("JWT_BYPASS") == "1" {
			return c.Next()
		}
		// fmt.Println(c.GetReqHeaders())
		bearer := c.Get("Authorization", "")
		// fmt.Println(bearer)
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
		c.Locals("email", claims.Email)
		c.Locals("first_name", claims.FirstName)
		c.Locals("last_name", claims.LastName)
		c.Locals("username", claims.Username)
		// fmt.Println("JOIN: ", strings.Join(claims.Roles, "|"))
		c.Locals("roles", claims.Roles)
		// fmt.Println("EMAIL: ", c.Locals("email"))
		// fmt.Println("ROLES: ", c.Locals("roles"))
		err2 := c.Next()

		return err2
	}
}
