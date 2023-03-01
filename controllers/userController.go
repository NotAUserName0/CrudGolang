package controllers

import (
	"crud/models"
	"strconv"
	//"log"

	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	result,err := models.Show()
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": "false",
			"message": err,
		})
	}

	if len(result) == 0 {
		return c.Status(200).JSON(&fiber.Map{
			"no_data": "true",
		})
	}

	return c.Status(200).JSON(result)
	
}

func GetUser(c *fiber.Ctx) error {
	Sid := c.Params("id")
	id, _ := strconv.Atoi(Sid)
	result,err := models.ShowOne(id)

	if err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
	}

	return c.Status(200).JSON(result)
}

func AddUsers(c *fiber.Ctx) error {
	newUser := new(models.User)
	err := c.BodyParser(newUser)

	if err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"success": "false",
			"message": err,
		})
	}

	if (newUser.NAME != "") && (newUser.GRADE > 0) {
		result, err := models.Add(newUser.NAME,newUser.GRADE)

		if err != nil {
			return c.Status(400).JSON(&fiber.Map{
				"success": "false",
				"message": err,
			})
		}

		c.Status(200).JSON(&fiber.Map{
			"success": "true",
			"nombre":  result.NAME,
			"grado": result.GRADE,
		})
		return nil
	}

	var nameCheck = false
	var gradeCheck = 0
	if newUser.NAME != "" {
		nameCheck = true
	}
	if newUser.GRADE > 0 {
		gradeCheck = 1
	}

	return c.Status(406).JSON(&fiber.Map{
		"no_data": "true",
		"name": strconv.FormatBool(nameCheck),
		"grade": gradeCheck,
	})

}

func DelUsers(c *fiber.Ctx) error {
	Sid := c.Params("id")
	id, _ := strconv.Atoi(Sid)

	result, err := models.Delete(id)
	
	if err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
		})
	}

	 c.Status(200).JSON(&fiber.Map{
		"success": true,
		"nombre":  result.NAME,
		"grado": result.GRADE,
	})
	return nil
}

func PutUsers(c *fiber.Ctx) error {
	Sid := c.Params("id")
	id, _ := strconv.Atoi(Sid)

	newUser := new(models.User)
	err := c.BodyParser(newUser)

	if err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
			"data":    nil,
		})
	}

	result, err := models.Update(id, newUser.NAME, newUser.GRADE)
	
	if err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
		})
	}

	 c.Status(200).JSON(&fiber.Map{
		"success": true,
		"data":    result,
	})
	return nil
}