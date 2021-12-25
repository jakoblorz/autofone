package www

import "github.com/gofiber/fiber/v2"

type AdditionalPropsAppender interface {
	AppendAdditionalProps(fiber.Map) fiber.Map
}
