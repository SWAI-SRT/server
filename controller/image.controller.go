package controller

import (
	"swai/service"

	"github.com/gofiber/fiber/v2"
)

type ImageController struct {
	imageService *service.ImageService
}

func NewImageController(imageService *service.ImageService) *ImageController {
	return &ImageController{imageService: imageService}
}

func (c *ImageController) UploadImage(ctx *fiber.Ctx) error {
	file, err := ctx.FormFile("imageUri")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "파일을 찾을 수 없습니다"})
	}

	fileContent, err := file.Open()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "파일을 열 수 없습니다"})
	}
	defer fileContent.Close()

	imageUrl, err := c.imageService.Upload(file.Filename, fileContent)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "이미지 업로드 실패"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"url": imageUrl})
}
