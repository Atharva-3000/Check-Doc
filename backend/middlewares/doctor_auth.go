package middlewares

import (
    "hi-doctor-be/utils"
    "strings"
    "github.com/gofiber/fiber/v2"
)

func DoctorAuth() fiber.Handler {
    return func(c *fiber.Ctx) error {
        authHeader := c.Get("Authorization")
        if authHeader == "" {
            return c.Status(401).JSON(fiber.Map{
                "error": "authorization header required",
            })
        }

        tokenParts := strings.Split(authHeader, " ")
        if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
            return c.Status(401).JSON(fiber.Map{
                "error": "invalid authorization header format",
            })
        }

        claims, err := utils.ValidateToken(tokenParts[1])
        if err != nil {
            return c.Status(401).JSON(fiber.Map{
                "error": "invalid token",
            })
        }

        // Add doctor info to context
        c.Locals("doctorID", claims.DoctorID)
        c.Locals("doctorPhone", claims.Phone)

        return c.Next()
    }
}