package controllers

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func UploadImage(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	var filename string
	files := form.File["image"]
	path := "uploads"
	os.Mkdir(path, 0755)

	for _, file := range files {
		filename = path + "/" + file.Filename

		if err := c.SaveFile(file, filename); err != nil {
			return err
		}
	}

	return c.JSON(fiber.Map{
		"url": os.Getenv("HOST") + "/api/" + filename,
	})
}
