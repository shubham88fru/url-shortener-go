package routes

import (
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/shubham88fru/url-shortener-go/database"
)

func ResolveURL(ctx *fiber.Ctx) error {
    url := ctx.Params("url")

    r := database.CreateClient(0)
    defer r.Close()

    value, err := r.Get(database.Ctx, url).Result()

    if err == redis.Nil {
        return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "Short not found",
        })
    } 

    if err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Internal server error", 
        })
    }

    rInr := database.CreateClient(1)
    defer rInr.Close()

    _ = rInr.Incr(database.Ctx, "counter")

    return ctx.Redirect(value, fiber.StatusMovedPermanently)

}