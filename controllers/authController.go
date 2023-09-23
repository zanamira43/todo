package controllers

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/zanamira43/todo/database"
	"github.com/zanamira43/todo/models"
	"github.com/zanamira43/todo/util"
)

// Register new User
func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		c.SendStatus(400)
		return c.JSON(fiber.Map{
			"message": "password does not match",
		})
	}

	user := models.User{
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
		RoleId:    1,
	}

	user.SetPassword(data["password"])
	database.DB.Create(&user)
	return c.JSON(user)
}

// login authenticate user
func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User
	// find if user exists or not
	database.DB.Where("email = ?", data["email"]).First(&user)
	if user.Id == 0 {
		c.SendStatus(404)
		return c.JSON(fiber.Map{
			"message": "User not found",
		})
	}
	// bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"]))
	if err := user.ComparePassword(data["password"]); err != nil {
		c.SendStatus(400)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}
	// generate jwt token ----
	token, err := util.GenerateJwt(strconv.Itoa(int(user.Id)))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	cookies := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookies)
	return c.JSON(fiber.Map{
		"message": "Success",
	})

}

// get authenticte user
func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	id, err := util.ParseJwt(cookie)
	if err != nil {
		c.SendStatus(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "User UnAuthenticated",
		})
	}

	var user models.User

	database.DB.Where("id = ?", id).First(&user)
	return c.JSON(user)

}

// logout user
func Logout(c *fiber.Ctx) error {

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "Success",
	})

}

// update user info
func UpdateInfo(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	cookie := c.Cookies("jwt")
	id, _ := util.ParseJwt(cookie)

	userId, _ := strconv.Atoi(id)

	user := models.User{
		Id:        uint(userId),
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
	}

	database.DB.Model(&user).Updates(user)
	return c.JSON(user)

}

// update password
func UpdatePassword(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		c.SendStatus(400)
		return c.JSON(fiber.Map{
			"message": "password does not match",
		})
	}

	cookie := c.Cookies("jwt")
	id, _ := util.ParseJwt(cookie)

	userId, _ := strconv.Atoi(id)

	user := models.User{
		Id: uint(userId),
	}

	user.SetPassword(data["password"])
	database.DB.Model(&user).Updates(user)
	return c.JSON(user)
}
