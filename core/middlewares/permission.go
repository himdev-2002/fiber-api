package middlewares

import (
	"him/fiber-api/core/handlers"
	"him/fiber-api/core/helpers"

	"github.com/gofiber/fiber/v2"
)

func RoutePermission(roles *helpers.Set[string], rel helpers.RelOp) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// fmt.Println(roles)
		// fmt.Println(rel)
		if len(roles.Values()) <= 0 {
			return c.Next()
		}

		if rel != helpers.IN && rel != helpers.NOTIN {
			return handlers.HandleError(c, &handlers.InternalServerError{Message: "Invalid Route Config"})
		}

		// auth_roles := c.Locals("roles")
		// fmt.Println(c.Locals("roles"))
		// auth_roles := strings.Split(c.Get("roles"), "|")
		if len(c.Locals("roles").([]string)) <= 0 {
			return handlers.HandleError(c, &handlers.UnauthorizedError{Message: "Not Authorized"})
		}
		// fmt.Println(auth_roles)

		var found = rel != helpers.IN
		for _, v := range c.Locals("roles").([]string) {
			// fmt.Printf("Index %d, Value %d\n", i, v)
			if roles.Contains(v) {
				if rel == helpers.IN {
					found = true
					break
				} else {
					found = false
					break
				}
			}
		}

		if !found {
			return handlers.HandleError(c, &handlers.UnauthorizedError{Message: "Not Authorized"})
		}

		return c.Next()
	}
}
